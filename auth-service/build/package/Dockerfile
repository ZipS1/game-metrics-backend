FROM golang:alpine3.20 AS builder
ENV CGO_ENABLED=0
ENV GOOS=linux
RUN apk update --no-cache && \
    apk add --no-cache tzdata

WORKDIR /build
COPY go.mod go.sum /build/
RUN go mod download
COPY . .
RUN go build -ldflags="-s -w" -o /app/auth-service ./cmd/auth-service/main.go

FROM alpine:3.20 AS final
ENV GIN_MODE=release
ENV TZ=Europe/Moscow

RUN apk update --no-cache && apk add --no-cache ca-certificates
COPY --from=builder /usr/share/zoneinfo/Europe/Moscow /usr/share/zoneinfo/Europe/Moscow

WORKDIR /app
COPY --from=builder /app/auth-service .
COPY configs/config.yml .

ENTRYPOINT ["./auth-service"]
CMD ["--config", "config.yml"]
EXPOSE 8080
