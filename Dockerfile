# Start from the official Go image
FROM golang:1.22

# Set working directory
WORKDIR /app

# Copy go.mod and source files
COPY go.mod ./
COPY main.go ./

# Build the binary
RUN go build -o app .

# Command to run the app
CMD ["./app"]
