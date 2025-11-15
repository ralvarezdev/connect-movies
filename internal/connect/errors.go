package connect

import (
	"errors"
)

var (
	ErrInDevelopment = errors.New("rpc method in development")
)