package service

import (
	"errors"

	"connectrpc.com/connect"
)

var (
	ErrMovieNotFound  = errors.New("movie not found for the given ID and this request")
	ConnErrMovieNotFound = connect.NewError(connect.CodeNotFound, ErrMovieNotFound)
)

var (
	ErrNilService    = errors.New("service is nil")
	ErrNilModelToMap = errors.New("model to map is nil")
)
