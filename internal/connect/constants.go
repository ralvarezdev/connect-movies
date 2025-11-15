package connect

import (
	goconnectrequest "github.com/ralvarezdev/go-connect/server/request"
	goconnectresponse "github.com/ralvarezdev/go-connect/server/response"

	internalloader "github.com/ralvarezdev/connect-movies/internal/loader"
)

const (
	// EnvAuthServiceAddress is the environment variable for the auth service address
	EnvAuthServiceAddress = "AUTH_SERVICE_ADDRESS"

	// EnvHTTPPort is the environment variable for the service HTTP port
	EnvHTTPPort = "HTTP_PORT"

	// EnvGRPCPort is the environment variable for the service gRPC port
	EnvGRPCPort = "GRPC_PORT"
)

var (
	// AuthServiceAddress is the auth service address
	AuthServiceAddress string

	// HTTPPort is the service HTPP port
	HTTPPort int

	// GRPCPort is the service gRPC port
	GRPCPort int

	// RequestInjector is the request injector
	RequestInjector goconnectrequest.Injector

	// ResponseInjector is the response injector
	ResponseInjector goconnectresponse.Injector
)

// Load loads the request and response injectors
func Load() {
	// Get the auth service address from the environment variable
	if err := internalloader.Loader.LoadVariable(
		EnvAuthServiceAddress,
		&AuthServiceAddress,
	); err != nil {
		panic(err)
	}

	// Get the service ports from the environment variable
	for env, dest := range map[string]*int{
		EnvHTTPPort: &HTTPPort,
		EnvGRPCPort: &GRPCPort,
	} {
		if err := internalloader.Loader.LoadIntVariable(
			env,
			dest,
		); err != nil {
			panic(err)
		}
	}

	// Create the request injector
	RequestInjector = goconnectrequest.NewDefaultInterceptor()

	// Create the response injector
	ResponseInjector = goconnectresponse.NewDefaultInterceptor(nil)
}
