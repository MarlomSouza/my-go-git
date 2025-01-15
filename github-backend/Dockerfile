# Build stage
FROM golang:1.23.4-alpine AS builder

# Set destination for COPY
WORKDIR /app

# Install git (required for fetching dependencies)
RUN apk add --no-cache git

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Print Go environment and dependencies
RUN go env
RUN go list -m all

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping ./cmd/main.go

# Final stage
FROM alpine:latest

# Set destination for the binary
WORKDIR /root/


# Copy the binary from the builder stage
COPY --from=builder /docker-gs-ping .

# Expose the application port
EXPOSE 5000

# Run the application
CMD ["./docker-gs-ping"]