# --- Build stage ---
FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

# --- Run stage ---
FROM debian:bookworm-slim

WORKDIR /app
COPY --from=builder /app/main .

# Load environment variables
ENV GIN_MODE=release
CMD ["./main"]
