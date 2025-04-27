FROM golang:alpine3.20 AS builder
ENV CGO_ENABLED=0
ENV GOOS=linux
RUN apk update --no-cache && \
    apk add --no-cache tzdata

WORKDIR /build
COPY ./players-service/go.mod ./players-service/go.sum /build/
COPY ./libs /libs
RUN go mod download
COPY ./players-service ./
RUN go build -ldflags="-s -w" -o /app/players-service ./cmd/players-service/main.go

FROM alpine:3.20 AS final
ENV GIN_MODE=release
ENV TZ=Europe/Moscow

RUN apk update --no-cache && apk add --no-cache ca-certificates
COPY --from=builder /usr/share/zoneinfo/Europe/Moscow /usr/share/zoneinfo/Europe/Moscow

WORKDIR /app
COPY --from=builder /app/players-service .
COPY ./players-service/configs/config.yml .

ENTRYPOINT ["./players-service"]
CMD ["--config", "config.yml"]
EXPOSE 8080
