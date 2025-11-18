package connect

import (
	"context"
	"log/slog"

	"github.com/ralvarezdev/proto-movies/gen/go/ralvarezdev/v1"
	"github.com/ralvarezdev/proto-movies/gen/go/ralvarezdev/v1/v1connect"

	internalservice "github.com/ralvarezdev/connect-movies/internal/service"
)

type (
	// Server is the gRPC server
	Server struct {
		logger  *slog.Logger
		service *internalservice.Service
		v1connect.UnimplementedMoviesServiceHandler
	}
)

// NewServer creates a new gRPC server
//
// Parameters:
//
//   - service: the service for the server
//   - logger: the logger
//
// Returns:
//
//   - *Server: the gRPC server
//   - error: if there was an error creating the server
func NewServer(service *internalservice.Service, logger *slog.Logger) (*Server, error) {
	// Check if the service is nil
	if service == nil {
		return nil, internalservice.ErrNilService
	}

	// Create the logger for the gRPC server
	if logger != nil {
		logger = logger.With(
			slog.String("component", "grpc_server"),
		)
	}

	return &Server{
		service: service,
		logger:  logger,
	}, nil
}

func (s Server) GetMovieCredits(
	ctx context.Context,
	request *v1.GetMovieCreditsRequest,
) (*v1.GetMovieCreditsResponse, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	// Call the service to get movie credits
	response, err := s.service.GetMovieCredits(ctx, request)
	if err != nil {
		if s.logger != nil {
			s.logger.Error("Error getting movie credits", slog.String("error", err.Error()))
		}
		return nil, err
	}
	return response, nil
}

func (s Server) GetNowPlayingMovies(
	ctx context.Context,
	request *v1.GetNowPlayingMoviesRequest,
) (*v1.GetNowPlayingMoviesResponse, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	// Call the service to get now playing movies
	response, err := s.service.GetNowPlayingMovies(ctx, request)
	if err != nil {
		if s.logger != nil {
			s.logger.Error("Error getting now playing movies", slog.String("error", err.Error()))
		}
		return nil, err
	}
	return response, nil
}

func (s Server) GetPopularMovies(
	ctx context.Context,
	request *v1.GetPopularMoviesRequest,
) (*v1.GetPopularMoviesResponse, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	// Call the service to get popular movies
	response, err := s.service.GetPopularMovies(ctx, request)
	if err != nil {
		if s.logger != nil {
			s.logger.Error("Error getting popular movies", slog.String("error", err.Error()))
		}
		return nil, err
	}
	return response, nil
}

func (s Server) GetTopRatedMovies(
	ctx context.Context,
	request *v1.GetTopRatedMoviesRequest,
) (*v1.GetTopRatedMoviesResponse, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	// Call the service to get top rated movies
	response, err := s.service.GetTopRatedMovies(ctx, request)
	if err != nil {
		if s.logger != nil {
			s.logger.Error("Error getting top rated movies", slog.String("error", err.Error()))
		}
		return nil, err
	}
	return response, nil
}

func (s Server) GetUpcomingMovies(
	ctx context.Context,
	request *v1.GetUpcomingMoviesRequest,
) (*v1.GetUpcomingMoviesResponse, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	// Call the service to get upcoming movies
	response, err := s.service.GetUpcomingMovies(ctx, request)
	if err != nil {
		if s.logger != nil {
			s.logger.Error("Error getting upcoming movies", slog.String("error", err.Error()))
		}
		return nil, err
	}
	return response, nil
}

func (s Server) SimilarMovies(
	ctx context.Context,
	request *v1.SimilarMoviesRequest,
) (*v1.SimilarMoviesResponse, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	// Call the service to get similar movies
	response, err := s.service.SimilarMovies(ctx, request)
	if err != nil {
		if s.logger != nil {
			s.logger.Error("Error getting similar movies", slog.String("error", err.Error()))
		}
		return nil, err
	}
	return response, nil
}

func (s Server) SearchMovies(ctx context.Context, request *v1.SearchMoviesRequest) (*v1.SearchMoviesResponse, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	// Call the service to search movies
	response, err := s.service.SearchMovies(ctx, request)
	if err != nil {
		if s.logger != nil {
			s.logger.Error("Error searching movies", slog.String("error", err.Error()))
		}
		return nil, err
	}
	return response, nil
}

func (s Server) GetMovieDetails(
	ctx context.Context,
	request *v1.GetMovieDetailsRequest,
) (*v1.GetMovieDetailsResponse, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	// Call the service to get movie details
	response, err := s.service.GetMovieDetails(ctx, request)
	if err != nil {
		if s.logger != nil {
			s.logger.Error("Error getting movie details", slog.String("error", err.Error()))
		}
		return nil, err
	}
	return response, nil
}

func (s Server) GetMovieReviews(
	ctx context.Context,
	request *v1.GetMovieReviewsRequest,
) (*v1.GetMovieReviewsResponse, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	// Call the service to get movie reviews
	response, err := s.service.GetMovieReviews(ctx, request)
	if err != nil {
		if s.logger != nil {
			s.logger.Error("Error getting movie reviews", slog.String("error", err.Error()))
		}
		return nil, err
	}
	return response, nil
}

func (s Server) GetMovieGenres(
	ctx context.Context,
	request *v1.GetMovieGenresRequest,
) (*v1.GetMovieGenresResponse, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	// Call the service to get movie genres
	response, err := s.service.GetMovieGenres(ctx, request)
	if err != nil {
		if s.logger != nil {
			s.logger.Error("Error getting movie genres", slog.String("error", err.Error()))
		}
		return nil, err
	}
	return response, nil
}	

func (s Server) DiscoverMovies(
	ctx context.Context,
	request *v1.DiscoverMoviesRequest,
) (*v1.DiscoverMoviesResponse, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	// Call the service to discover movies
	response, err := s.service.DiscoverMovies(ctx, request)
	if err != nil {
		if s.logger != nil {
			s.logger.Error("Error discovering movies", slog.String("error", err.Error()))
		}
		return nil, err
	}
	return response, nil
}

func (s Server) DeleteUserMovieReview(
	ctx context.Context,
	request *v1.DeleteUserMovieReviewRequest,
) (*v1.DeleteUserMovieReviewResponse, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	
	// Call the service to delete user movie review
	response, err := s.service.DeleteUserMovieReview(ctx, request)
	if err != nil {
		if s.logger != nil {
			s.logger.Error("Error deleting movie review", slog.String("error", err.Error()))
		}
		return nil, err
	}
	return response, nil
}

func (s Server) UpdateUserMovieReview(
	ctx context.Context,
	request *v1.UpdateUserMovieReviewRequest,
) (*v1.UpdateUserMovieReviewResponse, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	
	// Call the service to update user movie review
	response, err := s.service.UpdateUserMovieReview(ctx, request)
	if err != nil {
		if s.logger != nil {
			s.logger.Error("Error updating movie review", slog.String("error", err.Error()))
		}
		return nil, err
	}
	return response, nil
}

func (s Server) AddUserMovieReview(
	ctx context.Context,
	request *v1.AddUserMovieReviewRequest,
) (*v1.AddUserMovieReviewResponse, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	
	// Call the service to add user movie review
	response, err := s.service.AddUserMovieReview(ctx, request)
	if err != nil {
		if s.logger != nil {
			s.logger.Error("Error adding movie review", slog.String("error", err.Error()))
		}
		return nil, err
	}
	return response, nil
}

func (s Server) GetUserMovieReview(
	ctx context.Context,
	request *v1.GetUserMovieReviewRequest,
) (*v1.GetUserMovieReviewResponse, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	
	// Call the service to get user movie review
	response, err := s.service.GetUserMovieReview(ctx, request)
	if err != nil {
		if s.logger != nil {
			s.logger.Error("Error getting user movie review", slog.String("error", err.Error()))
		}
		return nil, err
	}
	return response, nil
}	