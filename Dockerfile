# Build stage
FROM golang:1.22.2-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o app

# Final image
FROM alpine
RUN apk add --no-cache ca-certificates

WORKDIR /app
COPY --from=builder /app/app ./app
COPY .env.dev .env

EXPOSE 10000
CMD ["./app"]