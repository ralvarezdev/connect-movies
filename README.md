# connect-movies

Movies microservice for the Connect platform, written in Go with ConnectRPC. It wraps [The Movie Database (TMDB)](https://www.themoviedb.org/) through the `go-tmdb-api` library to browse, search and discover movies, and lets authenticated users store their own reviews in PostgreSQL.

---

## Architecture

- ConnectRPC server (`cmd/server/main.go`) serving `MoviesService` from [`proto-movies`](https://github.com/ralvarezdev/proto-movies) over HTTP on `PORT`; `GET /health` returns `200 OK`.
- Request validation with `connectrpc.com/validate`, error-handling interceptor from `go-connect`, gRPC reflection (v1 and v1alpha).
- PostgreSQL via `pgx/v5`, schema in [`sql-movies`](https://github.com/ralvarezdev/sql-movies); Redis via `go-redis/v9`.
- JWT validation with the public key given by `-public-key`; token types from `connect-auth-types`.
- Helper libraries: `go-connect`, `go-jwt`, `go-databases`, `go-flags`, `go-loader`.

## RPC methods

- **TMDB catalogue** — `DiscoverMovies`, `SearchMovies`, `GetMovieDetails`, `GetMovieCredits`, `GetMovieGenres`, `GetMovieReviews`, `SimilarMovies`, `GetPopularMovies`, `GetTopRatedMovies`, `GetNowPlayingMovies`, `GetUpcomingMovies`.
- **User reviews (stored locally)** — `AddUserMovieReview`, `GetUserMovieReview`, `UpdateUserMovieReview`, `DeleteUserMovieReview`.

## Prerequisites

- Go 1.25.4 or newer, PostgreSQL and Redis.
- A running `connect-auth` service (`AUTH_SERVICE_ADDRESS`) and the JWT public key PEM matching its signing key.
- A TMDB API key of your own, kept out of version control.
- Access to the private `github.com/ralvarezdev/*` modules (set `GOPRIVATE`; the Dockerfile uses a `GITHUB_TOKEN` BuildKit secret).

## Configuration

Environment variables (template in `.env.example`):

- **General** — `MODE`, `PORT`, `AUTH_SERVICE_ADDRESS`
- **PostgreSQL** — `POSTGRES_DSN`, `POSTGRES_MAX_IDLE_CONNECTIONS`, `POSTGRES_MAX_OPEN_CONNECTIONS`
- **Redis** — `REDIS_ADDRESS`, `REDIS_USERNAME`, `REDIS_PASSWORD`, `REDIS_DB`
- **TMDB** — `TMDB_API_KEY` and `TMDB_*_IMAGE_WIDTH_SIZE` (cast and crew profile, movie poster, production company logo, details poster, avatar)

## Build and run

```bash
sh compile.sh      # go mod tidy, then go build -o bin/server/server ./cmd/server
sh serve.sh dev    # sources ./.env if present; first argument is the mode
```

Flags: `-mode` (default `dev`) and `-public-key` (default `/app/config/public_key.pem`, or `JWT_PUBLIC_KEY_PATH` via `serve.sh`). `format_lines.sh` normalizes CRLF with `dos2unix` and runs in the Dockerfile.

The two-stage `Dockerfile` (golang alpine, alpine runtime) exposes `8080` and runs `serve.sh`. Mount the JWT public key at the expected path:

```bash
docker build --secret id=GITHUB_TOKEN,env=GITHUB_TOKEN -t connect-movies .
```

## Project structure

```
cmd/server/          entry point
internal/connect/    RPC handlers
internal/databases/  postgres and redis setup
internal/jwt/        JWT configuration
internal/loader/     environment loading
internal/logger/     logger setup
```

There are no automated tests; lint with `golangci-lint run` (`.golangci.yml`).

## Related repositories

- [`proto-movies sql-movies`](https://github.com/ralvarezdev/proto-movies sql-movies) — protobuf definitions for the same microservice.
## License

GNU General Public License v3.0. See [LICENSE](LICENSE).
