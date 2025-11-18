package service

import (
	"errors"

	"connectrpc.com/connect"
)

var (
	ErrMovieNotFound  = errors.New("movie not found for the given ID and this request")
	ConnErrMovieNotFound = connect.NewError(connect.CodeNotFound, ErrMovieNotFound)
	ErrUserMovieReviewAlreadyExists = errors.New("user movie review already exists for the given user and movie")
	ConnErrUserMovieReviewAlreadyExists = connect.NewError(connect.CodeAlreadyExists, ErrUserMovieReviewAlreadyExists)
	ErrUserMovieReviewNotFound = errors.New("user movie review not found for the given user and movie")
	ConnErrUserMovieReviewNotFound = connect.NewError(connect.CodeNotFound, ErrUserMovieReviewNotFound)
)

var (
	ErrNilService    = errors.New("service is nil")
	ErrNilModelToMap = errors.New("model to map is nil")
)
