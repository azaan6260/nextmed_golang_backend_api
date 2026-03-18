# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy the source code
COPY . .

# Run go mod tidy to generate go.sum and download dependencies
RUN go mod tidy
RUN go mod download

# Build the application
RUN go build -o nextmed-backend main.go

# Final stage
FROM alpine:latest

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/nextmed-backend .

# Expose the port the app runs on
EXPOSE 3000

# Run the application
CMD ["./nextmed-backend"]
