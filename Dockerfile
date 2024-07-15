FROM golang:1.21-alpine AS builder

WORKDIR /app

RUN go version
ENV GOPATH=/

# dependencies
COPY ./ ./
RUN go mod download

# build
RUN go build -o main ./cmd/main.go
CMD ["./main"]