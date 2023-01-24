# Build stage
FROM golang:1.19-alpine3.16 AS builder
WORKDIR /app
COPY .. /app
RUN go build -o main weather_service/cmd/main.go

# Run stage
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/main .
COPY weather_service/config weather_service/config

EXPOSE 8000
CMD [ "/app/main" ]