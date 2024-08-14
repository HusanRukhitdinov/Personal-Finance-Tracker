FROM golang:1.22.6 AS builder

WORKDIR /app

COPY . .

RUN go mod download
RUN go mod tidy

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o myapp ./cmd

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/myapp .
COPY --from=builder /app/.env .

EXPOSE 8082

CMD ["./myapp"]
