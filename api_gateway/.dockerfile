# Builder bosqichi
FROM golang:1.22.5 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o auth ./cmd

# Asosiy bosqich
FROM alpine:latest


WORKDIR /app

COPY --from=builder /app/auth .
COPY .env .env

RUN chmod +x auth



CMD ["./auth"]