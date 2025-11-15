#!/bin/sh
set -e

# Load environment variables from .env file
set -a
if [ -f ./.env ]; then
  . ./.env
fi
set +a

# Check if the mode variable is set, default to the first argument or dev if not provided
if [ -z "$MODE" ]; then
  MODE=${1:-dev}
fi
export MODE

# Check if the HTTP port variable is set, default to the second argument or 8080 if not provided
if [ -z "$HTTP_PORT" ]; then
  HTTP_PORT=${2:-8080}
fi
export HTTP_PORT

# Check if the gRPC port variable is set, default to the second argument or 50051 if not provided
if [ -z "$GRPC_PORT" ]; then
  GRPC_PORT=${3:-50051}
fi
export GRPC_PORT

# Execute the Go binary on port specified
exec bin/server/server -mode=$MODE