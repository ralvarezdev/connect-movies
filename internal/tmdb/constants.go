package service

import (
	gotmdbapi "github.com/ralvarezdev/go-tmdb-api"

	internalloader "github.com/ralvarezdev/connect-movies/internal/loader"
)

const (
	// EnvTMDBAPIKey is the TMDB API key environment variable
	EnvTMDBAPIKey = "TMDB_API_KEY"

	// EnvCastMemberProfileImageWidthSize is the TMDB image width size for cast member profile images environment
	// variable
	EnvCastMemberProfileImageWidthSize = "TMDB_CAST_MEMBER_PROFILE_IMAGE_WIDTH_SIZE"

	// EnvCrewMemberProfileImageWidthSize is the TMDB image width size for crew member profile images environment
	// variable
	EnvCrewMemberProfileImageWidthSize = "TMDB_CREW_MEMBER_PROFILE_IMAGE_WIDTH_SIZE"

	// EnvSimpleMoviePosterImageWidthSize is the TMDB image width size for movie poster images on listings environment
	// variable
	EnvSimpleMoviePosterImageWidthSize = "TMDB_SIMPLE_MOVIE_POSTER_IMAGE_WIDTH_SIZE"

	// EnvProductionCompanyLogoImageWidthSize is the TMDB image width size for production company logo images
	// environment variable
	EnvProductionCompanyLogoImageWidthSize = "TMDB_PRODUCTION_COMPANY_LOGO_IMAGE_WIDTH_SIZE"

	// EnvMovieDetailsPosterImageWidthSize is the TMDB image width size for movie poster images on movie details
	// environment variable
	EnvMovieDetailsPosterImageWidthSize = "TMDB_MOVIE_DETAILS_POSTER_IMAGE_WIDTH_SIZE"

	// EnvAvatarImageWidthSize is the TMDB image width size for user avatar images environment variable
	EnvAvatarImageWidthSize = "TMDB_AVATAR_IMAGE_WIDTH_SIZE"
)

var (
	// TMDBAPIKey is the TMDB API key
	TMDBAPIKey string

	// TMDBClient is the TMDB API client
	TMDBClient *gotmdbapi.Client

	// CastMemberProfileImageWidthSize is the TMDB image width size for cast member profile images
	CastMemberProfileImageWidthSize int

	// CrewMemberProfileImageWidthSize is the TMDB image width size for crew member profile images
	CrewMemberProfileImageWidthSize int

	// SimpleMoviePosterImageWidthSize is the TMDB image width size for movie poster images on listings
	SimpleMoviePosterImageWidthSize int

	// ProductionCompanyLogoImageWidthSize is the TMDB image width size for production company logo images
	ProductionCompanyLogoImageWidthSize int

	// MovieDetailsPosterImageWidthSize is the TMDB image width size for movie poster images on movie details
	MovieDetailsPosterImageWidthSize int

	// AvatarImageWidthSize is the TMDB image width size for user avatar images
	AvatarImageWidthSize int
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

	// Load the image width sizes from the environment variables
	for env, dest := range map[string]*int{
		EnvCastMemberProfileImageWidthSize:     &CastMemberProfileImageWidthSize,
		EnvCrewMemberProfileImageWidthSize:     &CrewMemberProfileImageWidthSize,
		EnvSimpleMoviePosterImageWidthSize:     &SimpleMoviePosterImageWidthSize,
		EnvProductionCompanyLogoImageWidthSize: &ProductionCompanyLogoImageWidthSize,
		EnvMovieDetailsPosterImageWidthSize:    &MovieDetailsPosterImageWidthSize,
		EnvAvatarImageWidthSize:                &AvatarImageWidthSize,
	} {
		if err := internalloader.Loader.LoadIntVariable(
			env,
			dest,
		); err != nil {
			panic(err)
		}
	}

	// Initialize the TMDB API client
	tmdbClient, err := gotmdbapi.NewClient(TMDBAPIKey)
	if err != nil {
		panic(err)
	}
	TMDBClient = tmdbClient
}
