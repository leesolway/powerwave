FROM golang:1.22 as builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o app ./src

# Run app
FROM alpine:latest as production

WORKDIR /app

COPY --from=builder /app/app .

# Add an alias for the app binary for easy access
RUN echo 'alias powerwave="./app"' >> ~/.profile

EXPOSE 8080

CMD ["./app"]
