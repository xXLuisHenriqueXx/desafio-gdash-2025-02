FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build static binary (required for Alpine)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -a -installsuffix cgo -o queue-worker .

FROM alpine:latest

WORKDIR /app

# Add CA certificates (needed for HTTPS calls)
RUN apk add --no-cache ca-certificates

COPY --from=builder /app/queue-worker .

RUN chmod +x /app/queue-worker

CMD ["./queue-worker"]
