FROM golang:1.22.5-alpine AS builder

WORKDIR /app

RUN go install github.com/air-verse/air@latest

# dependencies
COPY go.* ./
RUN go mod download

COPY . .

# build
RUN CGO_ENABLED=0 GOOS=linux go build -o ../tmp/main ./cmd/main.go

CMD ["air", "-c", ".air.toml"]