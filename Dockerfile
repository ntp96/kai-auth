# Dockerfile References: https://docs.docker.com/engine/reference/builder/
FROM golang:1.14
WORKDIR /gitlab.com/tego-partner/kardiachain/kai-auth

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

COPY . .
RUN go install ./cmd/auth-service/...
