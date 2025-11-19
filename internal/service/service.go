package service

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	godatabases "github.com/ralvarezdev/go-databases"
	godatabasespgx "github.com/ralvarezdev/go-databases/sql/pgx"
	gotmdbapi "github.com/ralvarezdev/go-tmdb-api"
	v1 "github.com/ralvarezdev/proto-movies/gen/go/ralvarezdev/v1"
	redisauthtypes "github.com/ralvarezdev/redis-auth-types-go"
	sqlmovies "github.com/ralvarezdev/sql-movies/go"

	goauthjwtclaims "github.com/ralvarezdev/connect-auth-types-go/jwt/claims"

	internaltmdb "github.com/ralvarezdev/connect-movies/internal/tmdb"
)

type (
	// Service is the service for the gRPC server
	Service struct {
		tmdbClient           *gotmdbapi.Client
		pool                 *pgxpool.Pool
		redisUsernameHandler *redisauthtypes.UsernameHandler
	}
)

// NewService creates a new service
//
// Parameters:
//
// - tmdbClient: the TMDB API client
// - pool: the Postgres connection pool
// - redisUsernameHandler: the Redis username handler
//
// Returns:
//
// - *Service: the service
// - error: if there was an error creating the service
func NewService(
	tmdbClient *gotmdbapi.Client,
	pool *pgxpool.Pool,
	redisUsernameHandler *redisauthtypes.UsernameHandler,
) (*Service, error) {
	// Check if the Postgres pool is nil
	if pool == nil {
		return nil, godatabases.ErrNilPool
	}

	// Check if the Redis username handler is nil
	if redisUsernameHandler == nil {
		return nil, redisauthtypes.ErrNilUsernameHandler
	}

	// Check if the TMDB API client is nil
	if tmdbClient == nil {
		return nil, gotmdbapi.ErrNilClient
	}

	return &Service{
		tmdbClient:           tmdbClient,
		pool:                 pool,
		redisUsernameHandler: redisUsernameHandler,
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
	apiResponse, statusCode, err := s.tmdbClient.GetMovieCredits(ctx, request.GetId(), request.GetLanguage())
	if err != nil {
		if statusCode == http.StatusNotFound {
			return nil, ConnErrMovieNotFound
		}
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
	apiResponse, _, err := s.tmdbClient.GetMoviesTopRated(
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
	apiResponse, _, err := s.tmdbClient.GetMoviesPopular(
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
	apiResponse, _, err := s.tmdbClient.GetMoviesNowPlaying(
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
	apiResponse, _, err := s.tmdbClient.GetMoviesUpcoming(
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
	apiResponse, statusCode, err := s.tmdbClient.SimilarMovies(
		ctx,
		request.GetId(),
		request.GetLanguage(),
		request.GetPage(),
	)
	if err != nil {
		if statusCode == http.StatusNotFound {
			return nil, ConnErrMovieNotFound
		}
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
	apiResponse, _, err := s.tmdbClient.SearchMovies(
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
	apiResponse, statusCode, err := s.tmdbClient.GetMovieDetails(ctx, request.GetId(), request.GetLanguage())
	if err != nil {
		if statusCode == http.StatusNotFound {
			return nil, ConnErrMovieNotFound
		}
		panic(err)
	}

	return internaltmdb.MapToGetMovieDetailsResponse(apiResponse), nil
}

// GetMovieGenres gets the movie genres
//
// Parameters:
//
// - ctx: the context
// - request: the get movie genres request
//
// Returns:
//
// - *v1.GetMovieGenresResponse: the get movie genres response
// - error: if there was an error getting the movie genres
func (s *Service) GetMovieGenres(
	ctx context.Context,
	request *v1.GetMovieGenresRequest,
) (*v1.GetMovieGenresResponse, error) {
	if s == nil {
		panic(ErrNilService)
	}

	// Call TMDB API to get movie genres
	apiResponse, _, err := s.tmdbClient.GetGenresMovieList(
		ctx,
		request.GetLanguage(),
	)
	if err != nil {
		panic(err)
	}

	return internaltmdb.MapToGetMovieGenresResponse(apiResponse), nil
}

// DiscoverMovies discovers movies
//
// Parameters:
//
// - ctx: the context
// - request: the discover movies request
//
// Returns:
//
// - *v1.DiscoverMoviesResponse: the discover movies response
// - error: if there was an error discovering movies
func (s *Service) DiscoverMovies(
	ctx context.Context,
	request *v1.DiscoverMoviesRequest,
) (*v1.DiscoverMoviesResponse, error) {
	if s == nil {
		panic(ErrNilService)
	}

	// Map sort by enum to SortByEnum
	mappedSortBy := internaltmdb.MapToSortBy(*request.GetSortBy().Enum())

	// Map watch monetization types enum to WatchMonetizationTypesEnum slice
	mappedWatchMonetizationTypes := internaltmdb.MapToWatchMonetizationTypes(request.GetWithWatchMonetizationTypes())

	// Call TMDB API to discover movies
	apiResponse, _, err := s.tmdbClient.DiscoverMovies(
		ctx,
		&gotmdbapi.DiscoverMoviesQueryParameters{
			Certification:              request.GetCertification(),
			CertificationCountry:       request.GetCertificationCountry(),
			CertificationGTE:           request.GetCertificationGte(),
			CertificationLTE:           request.GetCertificationLte(),
			IncludeAdult:               request.GetIncludeAdult(),
			IncludeVideo:               request.GetIncludeVideo(),
			Language:                   request.GetLanguage(),
			PrimaryReleaseYear:         request.GetPrimaryReleaseYear(),
			PrimaryReleaseYearGTE:      request.GetPrimaryReleaseYearGte(),
			PrimaryReleaseYearLTE:      request.GetPrimaryReleaseYearLte(),
			Page:                       request.GetPage(),
			Region:                     request.GetRegion(),
			ReleaseDateGTE:             request.GetReleaseDateGte(),
			ReleaseDateLTE:             request.GetReleaseDateLte(),
			SortBy:                     mappedSortBy,
			VoteAverageGTE:             request.GetVoteAverageGte(),
			VoteAverageLTE:             request.GetVoteAverageLte(),
			VoteCountGTE:               request.GetVoteCountGte(),
			VoteCountLTE:               request.GetVoteCountLte(),
			WithGenres:                 request.GetWithGenres(),
			WithCompanies:              request.GetWithCompanies(),
			WithKeywords:               request.GetWithKeywords(),
			WithCast:                   request.GetWithCast(),
			WithCrew:                   request.GetWithCrew(),
			WithPeople:                 request.GetWithPeople(),
			WithOriginCountry:          request.GetWithOriginCountry(),
			WithOriginalLanguage:       request.GetWithOriginalLanguage(),
			WatchRegion:                request.GetWatchRegion(),
			WithRuntimeGTE:             request.GetWithRuntimeGte(),
			WithRuntimeLTE:             request.GetWithRuntimeLte(),
			WithWatchMonetizationTypes: mappedWatchMonetizationTypes,
			WithWatchProviders:         request.GetWithWatchProviders(),
			WithoutCompanies:           request.GetWithoutCompanies(),
			WithoutGenres:              request.GetWithoutGenres(),
			WithoutKeywords:            request.GetWithoutKeywords(),
			Year:                       request.GetYear(),
		},
	)
	if err != nil {
		panic(err)
	}

	return internaltmdb.MapToDiscoverMoviesResponse(apiResponse), nil
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
	apiResponse, statusCode, err := s.tmdbClient.GetMovieReviews(
		ctx,
		request.GetId(),
		request.GetLanguage(),
		request.GetPage(),
	)
	if err != nil {
		if statusCode == http.StatusNotFound {
			return nil, ConnErrMovieNotFound
		}
		panic(err)
	}

	// TODO: Add user reviews too

	return internaltmdb.MapToGetMovieReviewsResponse(apiResponse), nil
}

// AddUserMovieReview adds a user movie review
//
// Parameters:
//
// - ctx: the context
// - request: the add user movie review request
//
// Returns:
//
// - *v1.AddUserMovieReviewResponse: the add user movie review response
// - error: if there was an error adding the user movie review
func (s *Service) AddUserMovieReview(
	ctx context.Context,
	request *v1.AddUserMovieReviewRequest,
) (*v1.AddUserMovieReviewResponse, error) {
	if s == nil {
		panic(ErrNilService)
	}

	// Get the user ID from the auth response
	userID, err := goauthjwtclaims.GetSubject(ctx)
	if err != nil {
		panic(err)
	}

	// Call the stored procedure to create the movie review in Postgres
	if _, queryErr := s.pool.Exec(
		ctx,
		sqlmovies.CreateUserReviewProc,
		userID,
		request.GetId(),
		request.GetRating(),
		request.GetReview(),
	); queryErr != nil {
		isUniqueViolation, constraintName := godatabasespgx.IsUniqueViolationError(queryErr)
		if !isUniqueViolation {
			panic(queryErr)
		}

		// Check which unique constraint was violated
		if constraintName != sqlmovies.UserReviewsUniqueUserMovieReview {
			panic(queryErr)
		}
		return nil, ConnErrUserMovieReviewAlreadyExists
	}
	return &v1.AddUserMovieReviewResponse{}, nil
}

// UpdateUserMovieReview updates a user movie review
//
// Parameters:
//
// - ctx: the context
// - request: the update user movie review request
//
// Returns:
//
// - *v1.UpdateUserMovieReviewResponse: the update user movie review response
// - error: if there was an error updating the user movie review
func (s *Service) UpdateUserMovieReview(
	ctx context.Context,
	request *v1.UpdateUserMovieReviewRequest,
) (*v1.UpdateUserMovieReviewResponse, error) {
	if s == nil {
		panic(ErrNilService)
	}

	// Get the user ID from the auth response
	userID, err := goauthjwtclaims.GetSubject(ctx)
	if err != nil {
		panic(err)
	}

	// Call the stored procedure to update the movie review in Postgres
	var userReviewFound sql.NullBool
	if queryErr := s.pool.QueryRow(
		ctx,
		sqlmovies.UpdateUserReviewProc,
		userID,
		request.GetId(),
		request.GetRating(),
		request.GetReview(),
		nil,
	).Scan(
		&userReviewFound,
	); queryErr != nil {
		panic(queryErr)
	}

	// Check if the user review was found
	if !userReviewFound.Valid || !userReviewFound.Bool {
		return nil, ConnErrUserMovieReviewNotFound
	}

	return &v1.UpdateUserMovieReviewResponse{}, nil
}

// DeleteUserMovieReview deletes a user movie review
//
// Parameters:
//
// - ctx: the context
// - request: the delete user movie review request
//
// Returns:
//
// - *v1.DeleteUserMovieReviewResponse: the delete user movie review response
// - error: if there was an error deleting the user movie review
func (s *Service) DeleteUserMovieReview(
	ctx context.Context,
	request *v1.DeleteUserMovieReviewRequest,
) (*v1.DeleteUserMovieReviewResponse, error) {
	if s == nil {
		panic(ErrNilService)
	}

	// Get the user ID from the auth response
	userID, err := goauthjwtclaims.GetSubject(ctx)
	if err != nil {
		panic(err)
	}

	// Call the stored procedure to delete the movie review in Postgres
	var userReviewFound sql.NullBool
	if queryErr := s.pool.QueryRow(
		ctx,
		sqlmovies.DeleteUserReviewProc,
		userID,
		request.GetId(),
		nil,
	).Scan(
		&userReviewFound,
	); queryErr != nil {
		panic(queryErr)
	}

	// Check if the user review was found
	if !userReviewFound.Valid || !userReviewFound.Bool {
		return nil, ConnErrUserMovieReviewNotFound
	}

	return &v1.DeleteUserMovieReviewResponse{}, nil
}

func (s *Service) GetUserMovieReview(
	ctx context.Context,
	request *v1.GetUserMovieReviewRequest,
) (*v1.GetUserMovieReviewResponse, error) {
	if s == nil {
		panic(ErrNilService)
	}

	// Get the user ID from the auth response
	userID, err := goauthjwtclaims.GetSubject(ctx)
	if err != nil {
		panic(err)
	}

	// Call the stored procedure to get the movie review in Postgres
	var (
		outRating          sql.NullInt32
		outReviewText      sql.NullString
		outCreatedAt       sql.NullTime
		outUpdatedAt       sql.NullTime
		outUserReviewFound sql.NullBool
	)
	if queryErr := s.pool.QueryRow(
		ctx,
		sqlmovies.GetUserReviewProc,
		userID,
		request.GetId(),
		&outRating,
		&outReviewText,
		&outCreatedAt,
		&outUpdatedAt,
		&outUserReviewFound,
	).Scan(
		&outRating,
		&outReviewText,
		&outCreatedAt,
		&outUpdatedAt,
		&outUserReviewFound,
	); queryErr != nil {
		panic(queryErr)
	}

	// Check if the user review was found
	if !outUserReviewFound.Valid || !outUserReviewFound.Bool {
		return nil, ConnErrUserMovieReviewNotFound
	}

	return &v1.GetUserMovieReviewResponse{
		UserReview: &v1.UserMovieReview{
			Rating: int32(outRating.Int32),
			Review: outReviewText.String,
		},
	}, nil
}
