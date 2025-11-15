# Stage 1: Build the Go app
FROM golang:alpine AS builder

# Install git to fetch dependencies
RUN apk add --no-cache git dos2unix

# Add GitHub token secret for private repo access
RUN --mount=type=secret,id=GITHUB_TOKEN \
    GITHUB_TOKEN=$(cat /run/secrets/GITHUB_TOKEN) && \
    git config --global url."https://${GITHUB_TOKEN}:x-oauth-basic@github.com/".insteadOf "https://github.com/"

# Set the Current Working Directory inside the container and add source files
WORKDIR /app
COPY . .

# Run script to convert line endings
RUN sh convert_line_endings.sh

# Install Go dependencies
RUN go mod download

# Compile the Go app
RUN sh compile.sh

# Stage 2: Create a lightweight image for the Go app
FROM alpine AS server

# Set the Current Working Directory inside the container and copy the compiled Go app from the builder stage
WORKDIR /app
COPY --from=builder /app .

# Expose the HTTP and gRPC ports
EXPOSE 8080
EXPOSE 50051

# Run the Go app when the container launches
ENTRYPOINT ["sh", "serve.sh"]