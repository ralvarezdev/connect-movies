package service

import (
	"context"

	gotmdbapi "github.com/ralvarezdev/go-tmdb-api"
	v1 "github.com/ralvarezdev/proto-movies/gen/go/ralvarezdev/v1"

	internaltmdb "github.com/ralvarezdev/connect-movies/internal/tmdb"
)

type (
	// Service is the service for the gRPC server
	Service struct {
		tmdbClient *gotmdbapi.Client
	}
)

// NewService creates a new service
//
// Parameters:
//
// - tmdbClient: the TMDB API client
//
// Returns:
//
// - *Service: the service
// - error: if there was an error creating the service
func NewService(
	tmdbClient *gotmdbapi.Client,
) (*Service, error) {
	// Check if the TMDB API client is nil
	if tmdbClient == nil {
		return nil, gotmdbapi.ErrNilClient
	}

	return &Service{
		tmdbClient: tmdbClient,
	}, nil
}

// GetMovieCredits gets the movie credits
//
// Parameters:
//
// - ctx: the context
// - request: the get movie credits request
//
// Returns:
//
// - *v1.GetMovieCreditsResponse: the get movie credits response
// - error: if there was an error getting the movie credits
func (s *Service) GetMovieCredits(
	ctx context.Context,
	request *v1.GetMovieCreditsRequest,
) (*v1.GetMovieCreditsResponse, error) {
	if s == nil {
		panic(ErrNilService)
	}

	// Call TMDB API to get movie credits
	apiResponse, err := s.tmdbClient.GetMovieCredits(ctx, request.GetId(), request.GetLanguage())
	if err != nil {
		panic(err)
	}

	// Map TMDB API response to gRPC response
	return internaltmdb.MapToGetMovieCreditsResponse(apiResponse), nil
}

// GetTopRatedMovies gets the top rated movies
//
// Parameters:
//
// - ctx: the context
// - request: the get top rated movies request
//
// Returns:
//
// - *v1.GetTopRatedMoviesResponse: the get top rated movies response
// - error: if there was an error getting the top rated movies
func (s *Service) GetTopRatedMovies(
	ctx context.Context,
	request *v1.GetTopRatedMoviesRequest,
) (*v1.GetTopRatedMoviesResponse, error) {
	if s == nil {
		panic(ErrNilService)
	}

	// Call TMDB API to get top rated movies
	apiResponse, err := s.tmdbClient.GetMoviesTopRated(
		ctx,
		request.GetLanguage(),
		request.GetPage(),
		request.GetRegion(),
	)
	if err != nil {
		panic(err)
	}

	// Map TMDB API response to gRPC response
	return internaltmdb.MapToGetTopRatedMoviesResponse(apiResponse), nil
}

// GetPopularMovies gets the popular movies
//
// Parameters:
//
// - ctx: the context
// - request: the get popular movies request
//
// Returns:
//
// - *v1.GetPopularMoviesResponse: the get popular movies response
// - error: if there was an error getting the popular movies
func (s *Service) GetPopularMovies(
	ctx context.Context,
	request *v1.GetPopularMoviesRequest,
) (*v1.GetPopularMoviesResponse, error) {
	if s == nil {
		panic(ErrNilService)
	}

	// Call TMDB API to get popular movies
	apiResponse, err := s.tmdbClient.GetMoviesPopular(
		ctx,
		request.GetLanguage(),
		request.GetPage(),
		request.GetRegion(),
	)
	if err != nil {
		panic(err)
	}

	return internaltmdb.MapToGetPopularMoviesResponse(apiResponse), nil
}

// GetNowPlayingMovies gets the now playing movies
//
// Parameters:
//
// - ctx: the context
// - request: the get now playing movies request
//
// Returns:
//
// - *v1.GetNowPlayingMoviesResponse: the get now playing movies response
// - error: if there was an error getting the now playing movies
func (s *Service) GetNowPlayingMovies(
	ctx context.Context,
	request *v1.GetNowPlayingMoviesRequest,
) (*v1.GetNowPlayingMoviesResponse, error) {
	if s == nil {
		panic(ErrNilService)
	}

	// Call TMDB API to get now playing movies
	apiResponse, err := s.tmdbClient.GetMoviesNowPlaying(
		ctx,
		request.GetLanguage(),
		request.GetPage(),
		request.GetRegion(),
	)
	if err != nil {
		panic(err)
	}

	return internaltmdb.MapToGetNowPlayingMoviesResponse(apiResponse), nil
}

// GetUpcomingMovies gets the upcoming movies
//
// Parameters:
//
// - ctx: the context
// - request: the get upcoming movies request
//
// Returns:
//
// - *v1.GetUpcomingMoviesResponse: the get upcoming movies response
// - error: if there was an error getting the upcoming movies
func (s *Service) GetUpcomingMovies(
	ctx context.Context,
	request *v1.GetUpcomingMoviesRequest,
) (*v1.GetUpcomingMoviesResponse, error) {
	if s == nil {
		panic(ErrNilService)
	}

	// Call TMDB API to get upcoming movies
	apiResponse, err := s.tmdbClient.GetMoviesUpcoming(
		ctx,
		request.GetLanguage(),
		request.GetPage(),
		request.GetRegion(),
	)
	if err != nil {
		panic(err)
	}

	return internaltmdb.MapToGetUpcomingMoviesResponse(apiResponse), nil
}

// SimilarMovies maps similar movies
//
// Parameters:
//
// - ctx: the context
// - request: the get similar movies request
//
// Returns:
//
// - *v1.SimilarMoviesResponse: the get similar movies response
// - error: if there was an error getting the similar movies
func (s *Service) SimilarMovies(
	ctx context.Context,
	request *v1.SimilarMoviesRequest,
) (*v1.SimilarMoviesResponse, error) {
	if s == nil {
		panic(ErrNilService)
	}

	// Call TMDB API to get similar movies
	apiResponse, err := s.tmdbClient.SimilarMovies(ctx, request.GetId(), request.GetLanguage(), request.GetPage())
	if err != nil {
		panic(err)
	}

	return internaltmdb.MapToSimilarMoviesResponse(apiResponse), nil
}

// SearchMovies searches for movies
//
// Parameters:
//
// - ctx: the context
// - request: the search movies request
//
// Returns:
//
// - *v1.SearchMoviesResponse: the search movies response
// - error: if there was an error searching for movies
func (s *Service) SearchMovies(ctx context.Context, request *v1.SearchMoviesRequest) (*v1.SearchMoviesResponse, error) {
	if s == nil {
		panic(ErrNilService)
	}

	// Call TMDB API to search for movies
	apiResponse, err := s.tmdbClient.SearchMovies(
		ctx,
		request.GetQuery(),
		request.GetIncludeAdult(),
		request.GetLanguage(),
		request.GetPage(),
		request.GetYear(),
		request.GetRegion(),
		request.GetPrimaryReleaseYear(),
	)
	if err != nil {
		panic(err)
	}

	return internaltmdb.MapToSearchMoviesResponse(apiResponse), nil
}

// GetMovieDetails gets the movie details
//
// Parameters:
//
// - ctx: the context
// - request: the get movie details request
//
// Returns:
//
// - *v1.GetMovieDetailsResponse: the get movie details response
// - error: if there was an error getting the movie details
func (s *Service) GetMovieDetails(
	ctx context.Context,
	request *v1.GetMovieDetailsRequest,
) (*v1.GetMovieDetailsResponse, error) {
	if s == nil {
		panic(ErrNilService)
	}

	// Call TMDB API to get movie details
	apiResponse, err := s.tmdbClient.GetMovieDetails(ctx, request.GetId(), request.GetLanguage())
	if err != nil {
		panic(err)
	}

	return internaltmdb.MapToGetMovieDetailsResponse(apiResponse), nil
}

// GetMovieReviews gets the movie reviews
//
// Parameters:
//
// - ctx: the context
// - request: the get movie reviews request
//
// Returns:
//
// - *v1.GetMovieReviewsResponse: the get movie reviews response
// - error: if there was an error getting the movie reviews
func (s *Service) GetMovieReviews(
	ctx context.Context,
	request *v1.GetMovieReviewsRequest,
) (*v1.GetMovieReviewsResponse, error) {
	if s == nil {
		panic(ErrNilService)
	}

	// Call TMDB API to get movie reviews
	apiResponse, err := s.tmdbClient.GetMovieReviews(ctx, request.GetId(), request.GetLanguage(), request.GetPage())
	if err != nil {
		panic(err)
	}

	// TODO: Add user reviews too

	return internaltmdb.MapToGetMovieReviewsResponse(apiResponse), nil
}
