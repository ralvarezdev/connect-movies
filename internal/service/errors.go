package service

import (
	"errors"
)

var (
	ErrNilService    = errors.New("service is nil")
	ErrNilModelToMap = errors.New("model to map is nil")
)
