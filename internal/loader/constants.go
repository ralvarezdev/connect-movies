package loader

import (
	"log/slog"

	"github.com/joho/godotenv"
	goflagsmode "github.com/ralvarezdev/go-flags/mode"
	goloaderenv "github.com/ralvarezdev/go-loader/env"

	internallogger "github.com/ralvarezdev/connect-movies/internal/logger"
)

var (
	// Loader is the environment variables loader
	Loader goloaderenv.Loader
)

// Load loads the loader
//
// Parameters:
//
//	mode: The application mode
//	logger: The logger to use
func Load(mode *goflagsmode.Flag, logger *slog.Logger) {
	// Load the environment variables loader
	loader, err := goloaderenv.NewDefaultLoader(
		func() error {
			// Load .env file
			if mode != nil && !mode.IsProd() {
				if err := godotenv.Load(); err != nil {
					internallogger.Logger.Warn("Could not load .env file", slog.String("error", err.Error()))
				}
			}
			return nil
		},
		logger,
	)
	if err != nil {
		panic(err)
	}
	Loader = loader
}
