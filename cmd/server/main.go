package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"connectrpc.com/connect"
	"connectrpc.com/grpcreflect"
	"connectrpc.com/validate"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/joho/godotenv"
	goconnectralvarezdevv1auth "github.com/ralvarezdev/go-connect-ralvarezdev/v1/auth"
	goconnectauth "github.com/ralvarezdev/go-connect/server/interceptor/auth"
	goconnecterrorhandler "github.com/ralvarezdev/go-connect/server/interceptor/errorhandler"
	goflagsmode "github.com/ralvarezdev/go-flags/mode"
	gojwtflags "github.com/ralvarezdev/go-jwt/flags"
	redisauthtypes "github.com/ralvarezdev/redis-auth-types-go"

	authv1connect "github.com/ralvarezdev/proto-auth/gen/go/ralvarezdev/v1/v1connect"
	protomovies "github.com/ralvarezdev/proto-movies/gen/go"
	"github.com/ralvarezdev/proto-movies/gen/go/ralvarezdev/v1/v1connect"

	internalconnect "github.com/ralvarezdev/connect-movies/internal/connect"
	internalpostgres "github.com/ralvarezdev/connect-movies/internal/databases/postgres"
	internalredis "github.com/ralvarezdev/connect-movies/internal/databases/redis"
	internaljwt "github.com/ralvarezdev/connect-movies/internal/jwt"
	internalloader "github.com/ralvarezdev/connect-movies/internal/loader"
	internallogger "github.com/ralvarezdev/connect-movies/internal/logger"
	internalservice "github.com/ralvarezdev/connect-movies/internal/service"
	internaltmdb "github.com/ralvarezdev/connect-movies/internal/tmdb"
)

var (
	// DefaultPublicKeyPath is the default path to the JWT public key
	DefaultPublicKeyPath = "/app/config/public_key.pem"

	// ModeFlag is the mode flag
	ModeFlag = goflagsmode.NewFlag(
		goflagsmode.Dev,
		goflagsmode.AllowedModes,
	)

	// PublicKeyPathFlag is the public key flag
	PublicKeyPathFlag = gojwtflags.NewPublicKeyFlag(
		&DefaultPublicKeyPath,
	)

	// ListenConfig is the net.ListenConfig to use
	ListenConfig = net.ListenConfig{}
)

// init initializes the flags and calls the load functions
func init() {
	// Define the mode flag and the JWT public key file path flag
	goflagsmode.SetFlag(ModeFlag)
	gojwtflags.SetPublicKeyFlag(PublicKeyPathFlag)

	// Parse the flags
	flag.Parse()

	// Call the load functions
	internallogger.Load(ModeFlag)
	internalloader.Load(ModeFlag, internallogger.Logger)
	internalpostgres.Load(ModeFlag)
	internalredis.Load()
	internaljwt.Load(ModeFlag, PublicKeyPathFlag, internalredis.Client, internallogger.Logger)
	internaltmdb.Load()
	internalconnect.Load()

	// Log that the load functions were called
	internallogger.Logger.Info(
		"Called load functions",
		slog.String("mode", ModeFlag.Value()),
	)
}

func main() {
	// Ensure the Redis client is closed on exit
	defer internalredis.Client.Close()

	// Create a context that is canceled on SIGINT or SIGTERM
	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	// Create the auth gRPC service client
	authClient := authv1connect.NewAuthServiceClient(
		http.DefaultClient,
		internalconnect.AuthServiceAddress,
		connect.WithGRPC(),
	)

	// Create the Postgres database service
	postgresPool, err := pgxpool.NewWithConfig(
		ctx,
		internalpostgres.PoolConfig,
	)
	if err != nil {
		panic(err)
	}
	defer postgresPool.Close()

	// Create the Redis username handler
	redisUsernameHandler, err := redisauthtypes.NewUsernameHandler(
		internalredis.Client,
	)
	if err != nil {
		panic(err)
	}

	// Create the service
	service, err := internalservice.NewService(
		internaltmdb.TMDBClient,
		postgresPool,
		redisUsernameHandler,
		// internalconnect.RequestInjector,
		// internalconnect.ResponseInjector,
	)
	if err != nil {
		panic(err)
	}

	// Create the gRPC auth Server
	connectServer, err := internalconnect.NewServer(
		service,
		internallogger.Logger,
	)
	if err != nil {
		panic(err)
	}

	// Create the refresh token function
	refreshTokenFn, err := goconnectralvarezdevv1auth.CreateRefreshTokenFn(
		ctx,
		authClient,
		internalconnect.RequestInjector,
		internalconnect.ResponseInjector,
	)
	if err != nil {
		panic(err)
	}

	// Initialize the auth interceptor
	authInterceptor, err := goconnectauth.NewInterceptor(
		ModeFlag,
		internaljwt.Validator,
		protomovies.Interceptions,
		&goconnectauth.Options{
			RefreshTokenFn: refreshTokenFn,
		},
		internallogger.Logger,
	)
	if err != nil {
		panic(err)
	}

	// Initialize the error handler interceptor
	errorHandler, err := goconnecterrorhandler.NewInterceptor(ModeFlag, internallogger.Logger)
	if err != nil {
		panic(err)
	}

	// Create the Connect mux and register the auth service handler
	mux := http.NewServeMux()
	path, handler := v1connect.NewMoviesServiceHandler(
		connectServer,
		connect.WithInterceptors(
			validate.NewInterceptor(),
			errorHandler.HandleError(),
			authInterceptor.Authenticate(),
		),
	)
	mux.Handle(path, handler)

	// Add a health check endpoint
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// Rgister reflection service on gRPC server.
	reflector := grpcreflect.NewStaticReflector(
		v1connect.MoviesServiceName,
	)
	mux.Handle(grpcreflect.NewHandlerV1(reflector))

	// Many tools still expect the older version of the server reflection API, so
	// most servers should mount both handlers.
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))

	// Create the protocols for HTTP/1.1 and HTTP/2 Cleartext
	protocols := new(http.Protocols)
	protocols.SetHTTP1(true)
	protocols.SetHTTP2(true)
	protocols.SetUnencryptedHTTP2(true)

	// Create server for Movies Service
	server := http.Server{
		Addr:      fmt.Sprintf("0.0.0.0:%d", internalconnect.Port),
		Handler:   mux,
		Protocols: protocols,
	}

	// Start the Movies server
		internallogger.Logger.Info(
			"Starting Movies server...",
			slog.Int("port", internalconnect.Port),
		)
		if listenErr := server.ListenAndServe(); listenErr != nil && !errors.Is(listenErr, http.ErrServerClosed) {
			internallogger.Logger.Error(
				"Could not start Movies server",
				slog.String("error", listenErr.Error()),
			)
			panic(listenErr)
		}

	// Wait for signal
	<-ctx.Done()
	internallogger.Logger.Info("Shutting down gracefully...")
}
