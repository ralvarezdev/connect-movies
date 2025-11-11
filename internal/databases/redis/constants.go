package redis

import (
	"github.com/go-redis/redis/v8"

	internalloader "github.com/ralvarezdev/connect-movies-go/internal/loader"
)

const (
	// EnvRedisAddress is the environment variable for the Redis address
	EnvRedisAddress = "REDIS_ADDRESS"

	// EnvRedisUsername is the environment variable for the Redis username
	EnvRedisUsername = "REDIS_USERNAME"

	// EnvRedisPassword is the environment variable for the Redis password
	EnvRedisPassword = "REDIS_PASSWORD"

	// EnvRedisDB is the environment variable for the Redis database number
	EnvRedisDB = "REDIS_DB"
)

var (
	// RedisAddress is the Redis address
	RedisAddress string

	// RedisUsername is the Redis username
	RedisUsername string

	// RedisPassword is the Redis password
	RedisPassword string

	// RedisDB is the Redis database number
	RedisDB int

	// Client is the Redis client
	Client *redis.Client
)

// Load initializes the Redis client
func Load() {
	// Get the Redis address, username and password from the environment variables
	for env, dest := range map[string]*string{
		EnvRedisAddress:  &RedisAddress,
		EnvRedisUsername: &RedisUsername,
		EnvRedisPassword: &RedisPassword,
	} {
		if err := internalloader.Loader.LoadVariable(
			env,
			dest,
		); err != nil {
			panic(err)
		}
	}

	// Get the Redis database number from the environment variable
	if err := internalloader.Loader.LoadIntVariable(
		EnvRedisDB,
		&RedisDB,
	); err != nil {
		panic(err)
	}

	// Create the redis client
	Client = redis.NewClient(
		&redis.Options{
			Addr:     RedisAddress,
			Username: RedisUsername,
			Password: RedisPassword,
			DB:       RedisDB,
		},
	)
}
