# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o app

# Final image
FROM alpine:latest
RUN apk add --no-cache ca-certificates
COPY --from=builder /app/app /app
EXPOSE 10000
CMD ["/app"]
