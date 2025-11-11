package jwt

import (
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
	goflagsmode "github.com/ralvarezdev/go-flags/mode"
	gojwttoken "github.com/ralvarezdev/go-jwt/token"
	gojwtvalidator "github.com/ralvarezdev/go-jwt/token/validator"
	gojwttokenclaims "github.com/ralvarezdev/go-jwt/token/claims"
	gojwttokenclaimsredis "github.com/ralvarezdev/go-jwt/token/claims/redis"

	internalloader "github.com/ralvarezdev/connect-movies-go/internal/loader"
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
	PublicKey string
	
	// Durations are the JWT tokens duration
	Durations = make(map[gojwttoken.Token]time.Duration)
	
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
//   - redisClient: the Redis client to use
func Load(
	mode *goflagsmode.Flag,
	redisClient *redis.Client,
) {
	// Get the JWT public key
	if err := internalloader.Loader.LoadVariable(EnvPublicKey, &PublicKey); err != nil {
			panic(err)
		}
		PublicKey = strings.ReplaceAll(PublicKey, `\n`, "\n")
	

	// Get the JWT tokens duration
	for key, env := range map[gojwttoken.Token]string{
		gojwttoken.AccessToken:  EnvAccessTokenDuration,
		gojwttoken.RefreshToken: EnvRefreshTokenDuration,
	} {
		var tokenDuration time.Duration
		if err := internalloader.Loader.LoadDurationVariable(
			env,
			&tokenDuration,
		); err != nil {
			panic(err)
		}
		Durations[key] = tokenDuration
	}
	
	// Initialize the token validator
	tokenValidator, err := gojwttokenclaimsredis.NewTokenValidator(
		redisClient,
		nil,
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
		[]byte(PublicKey),
		claimsValidator,
		mode,
	)
	if err != nil {
		panic(err)
	}
	Validator = validator
}