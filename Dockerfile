# Start from the official Golang image for building
FROM golang:1.24.5-alpine AS builder

RUN apk add --no-cache build-base

WORKDIR /app

RUN apk add --no-cache build-base
COPY db/ db/
COPY routes/ routes/
COPY static/ static/
COPY main.go .

# Copy go mod and sum first, then source code
COPY go.mod go.sum ./
RUN go mod tidy

# Build the Go app (static binary)
RUN CGO_ENABLED=1  go build -ldflags="-s -w" -o main .

# Start a minimal image for runtime
FROM alpine:latest

WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/main .

# Expose app port
EXPOSE 7666

# Run the Go app
CMD ["./main"]
