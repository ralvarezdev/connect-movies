package postgres

import (
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	goflagsmode "github.com/ralvarezdev/go-flags/mode"

	internalloader "github.com/ralvarezdev/connect-movies/internal/loader"
)

const (
	// EnvDSN is the key of the DSN for the Postgres database
	EnvDSN = "POSTGRES_DSN"

	// EnvMaxOpenConnections is the key of the maximum number of open connections for the Postgres database
	EnvMaxOpenConnections = "POSTGRES_MAX_OPEN_CONNECTIONS"

	// EnvMaxIdleConnections is the key of the maximum number of idle connections for the Postgres database
	EnvMaxIdleConnections = "POSTGRES_MAX_IDLE_CONNECTIONS"
)

var (
	// DSN is the DSN for the Postgres database
	DSN string

	// MaxOpenConnections is the maximum number of open connections for the Postgres database
	MaxOpenConnections int

	// MaxIdleConnections is the maximum number of idle connections for the Postgres database
	MaxIdleConnections int

	// PoolConfig is the Postgres pool configuration
	PoolConfig *pgxpool.Config
)

// Load loads the Postgres constants
//
// Parameters:
//
//   - mode: the mode flag to determine the logging level
func Load(mode *goflagsmode.Flag) {
	// Load the DSN for the Postgres database
	if err := internalloader.Loader.LoadVariable(
		EnvDSN,
		&DSN,
	); err != nil {
		panic(err)
	}

	// Load the maximum number of open and idle connections for the Postgres database
	for key, dest := range map[string]*int{
		EnvMaxOpenConnections: &MaxOpenConnections,
		EnvMaxIdleConnections: &MaxIdleConnections,
	} {
		if err := internalloader.Loader.LoadIntVariable(
			key,
			dest,
		); err != nil {
			panic(err)
		}
	}

	// Create the Postgres pool configuration
	poolConfig, err := pgxpool.ParseConfig(DSN)
	if err != nil {
		panic(err)
	}

	// Set the maximum number of open and idle connections
	poolConfig.MaxConns = int32(MaxOpenConnections)
	poolConfig.MinConns = int32(MaxIdleConnections)
	poolConfig.MaxConnLifetime = time.Hour
	poolConfig.MaxConnIdleTime = time.Hour
	poolConfig.HealthCheckPeriod = 5 * time.Minute
	poolConfig.MaxConnLifetimeJitter = 5 * time.Minute

	PoolConfig = poolConfig
}
