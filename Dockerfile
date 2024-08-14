# Dockerfile

FROM golang:1.20-alpine AS builder

# Install Air
RUN go install github.com/cosmtrek/air@latest

# Create and set working directory
WORKDIR /app

# Copy the source code
COPY . .

# Build the Go application
RUN go build -o tsa .

# Create the final image
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/tsa .
CMD ["./tsa"]
