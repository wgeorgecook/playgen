# Build phase
FROM golang:1.23-alpine AS builder

# Install protoc and plugin dependencies
RUN apk add --no-cache git make protobuf protobuf-dev \
    && go install google.golang.org/protobuf/cmd/protoc-gen-go@latest \
    && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Set the cwd inside the container
WORKDIR /playgen

# Copy go.mod and go.sum files to the workspace
COPY go.mod go.sum ./

# Install Go dependencies
RUN go mod download

# Copy the source code to the workspace
COPY . .

# Generate Proto code
RUN go generate ./...

# Build the Go app
RUN go build -o playgen .

# Run phase
FROM alpine:3.18

# Set the Current Working Directory inside the container
WORKDIR /playgen

# Copy the binary from the builder phase
COPY --from=builder /playgen/playgen .

# Command to run the executable
CMD ["./playgen"]
