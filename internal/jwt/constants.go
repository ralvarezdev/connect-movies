package jwt

import (
	"log/slog"
	"os"

	goflagsmode "github.com/ralvarezdev/go-flags/mode"
	gojwtflags "github.com/ralvarezdev/go-jwt/flags"
	gojwttokenclaims "github.com/ralvarezdev/go-jwt/token/claims"
	gojwttokenclaimsredis "github.com/ralvarezdev/go-jwt/token/claims/redis"
	gojwtvalidator "github.com/ralvarezdev/go-jwt/token/validator"
	"github.com/redis/go-redis/v9"
)

const (
	// EnvPublicKey is the key of the JWT public key
	EnvPublicKey = "JWT_PUBLIC_KEY"

	// EnvAccessTokenDuration is the key of the access token duration
	EnvAccessTokenDuration = "ACCESS_TOKEN_DURATION"

	// EnvRefreshTokenDuration is the key of the refresh token duration
	EnvRefreshTokenDuration = "REFRESH_TOKEN_DURATION"
)

var (
	// PublicKey is the JWT public key
	PublicKey []byte

	// TokenValidator is the cache token validator
	TokenValidator gojwttokenclaims.TokenValidator

	// ClaimsValidator is the cache claims validator
	ClaimsValidator gojwttokenclaims.ClaimsValidator

	// Validator is the JWT validator
	Validator gojwtvalidator.Validator
)

// Load loads the JWT constants
//
// Parameters:
//
//   - mode: the mode flag to determine the logging level
//   - publicKeyFlag: the public key flag
//   - redisClient: the Redis client to use
//   - logger: the logger to use
func Load(
	mode *goflagsmode.Flag,
	publicKeyFlag *gojwtflags.PublicKeyFlag,
	redisClient *redis.Client,
	logger *slog.Logger,
) {
	// Load the JWT public key from file
	publicKeyContent, err := os.ReadFile(publicKeyFlag.Path())
	if err != nil {
		panic(err)
	}
	PublicKey = publicKeyContent

	// Initialize the token validator
	tokenValidator, err := gojwttokenclaimsredis.NewTokenValidator(
		redisClient,
		logger,
	)
	if err != nil {
		panic(err)
	}
	TokenValidator = tokenValidator

	// Initialize the claims validator based on the mode
	claimsValidator, err := gojwttokenclaims.NewDefaultClaimsValidator(
		TokenValidator,
	)
	if err != nil {
		panic(err)
	}
	ClaimsValidator = claimsValidator

	// Create the JWT validator with ED25519 public key
	validator, err := gojwtvalidator.NewEd25519Validator(
		PublicKey,
		claimsValidator,
		mode,
	)
	if err != nil {
		panic(err)
	}
	Validator = validator
}
