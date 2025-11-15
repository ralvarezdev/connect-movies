package service

import (
	gotmdbapi "github.com/ralvarezdev/go-tmdb-api"

	internalloader "github.com/ralvarezdev/connect-movies/internal/loader"
)

const (
	// EnvTMDBAPIKey is the TMDB API key environment variable
	EnvTMDBAPIKey = "TMDB_API_KEY"
)

var (
	// TMDBAPIKey is the TMDB API key
	TMDBAPIKey string

	// TMDBClient is the TMDB API client
	TMDBClient *gotmdbapi.Client
)

// Load loads the TMDB API key
func Load() {
	// Get the TMDB API key from the environment variable
	if err := internalloader.Loader.LoadVariable(
		EnvTMDBAPIKey,
		&TMDBAPIKey,
	); err != nil {
		panic(err)
	}

	// Initialize the TMDB API client
	tmdbClient, err := gotmdbapi.NewClient(TMDBAPIKey)
	if err != nil {
		panic(err)
	}
	TMDBClient = tmdbClient
}
