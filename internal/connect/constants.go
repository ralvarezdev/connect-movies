package connect

import (
	goconnectrequest "github.com/ralvarezdev/go-connect/server/request"
	goconnectresponse "github.com/ralvarezdev/go-connect/server/response"

	internalloader "github.com/ralvarezdev/connect-movies/internal/loader"
)

const (
	// EnvAuthServiceAddress is the environment variable for the auth service address
	EnvAuthServiceAddress = "AUTH_SERVICE_ADDRESS"

	// EnvPort is the environment variable for the service Movies server port
	EnvPort = "PORT"
)

var (
	// AuthServiceAddress is the auth service address
	AuthServiceAddress string

	// Port is the Movies server port
	Port int

	// RequestInjector is the request injector
	RequestInjector goconnectrequest.Injector

	// ResponseInjector is the response injector
	ResponseInjector goconnectresponse.Injector
)

// Load loads the request and response injectors
func Load() {
	// Get the auth service address from the environment variable
	if err := internalloader.Loader.LoadVariable(
		EnvAuthServiceAddress,
		&AuthServiceAddress,
	); err != nil {
		panic(err)
	}

	// Get the service port from the environment variable
	if err := internalloader.Loader.LoadIntVariable(
		EnvPort,
		&Port,
		); err != nil {
			panic(err)
		}

	// Create the request injector
	RequestInjector = goconnectrequest.NewDefaultInterceptor()

	// Create the response injector
	ResponseInjector = goconnectresponse.NewDefaultInterceptor(nil)
}
