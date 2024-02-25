FROM golang:1.22 as builder

WORKDIR /app

COPY . .

# Build HTTP server
RUN CGO_ENABLED=0 GOOS=linux go build -o http ./cmd/http

# Build CLI tool
RUN CGO_ENABLED=0 GOOS=linux go build -o cli ./cmd/cli

# Serve HTTP server
FROM alpine:latest as http

WORKDIR /http

COPY --from=builder /app/http .

EXPOSE 8080

CMD ["./http"]


# CLI tool
FROM alpine:latest as cli

WORKDIR /cli

COPY --from=builder /app/cli .

ENTRYPOINT ["./cli"]