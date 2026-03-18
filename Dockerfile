# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum (if it exists)
COPY go.mod ./
# Since we don't have go.sum, we run go mod tidy here
RUN go mod tidy

# Copy the rest of the source code
COPY . .

# Build the application
RUN go build -o nextmed-backend main.go

# Final stage
FROM alpine:latest

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/nextmed-backend .
# Copy the .env file if it exists
COPY --from=builder /app/.env .

# Expose the port the app runs on
EXPOSE 3000

# Run the application
CMD ["./nextmed-backend"]
