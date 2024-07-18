FROM golang:1.22.5-alpine AS builder

WORKDIR /app

RUN go install github.com/air-verse/air@latest

# dependencies
COPY go.* ./
RUN go mod download

COPY . .

# build
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

EXPOSE 8000

CMD ["air", "-c", ".air.toml"]