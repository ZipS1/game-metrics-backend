FROM golang:alpine3.20 AS builder
ENV CGO_ENABLED=0
ENV GOOS=linux
RUN apk update --no-cache && \
    apk add --no-cache tzdata

WORKDIR /build
COPY ./api-gateway/go.mod ./api-gateway/go.sum /build/
RUN go mod download
COPY ./api-gateway .
RUN go build -ldflags="-s -w" -o /app/api-gateway ./cmd/api-gateway/main.go

FROM alpine:3.20 AS final
ENV GIN_MODE=release
ENV TZ=Europe/Moscow

RUN apk update --no-cache && apk add --no-cache ca-certificates
COPY --from=builder /usr/share/zoneinfo/Europe/Moscow /usr/share/zoneinfo/Europe/Moscow

WORKDIR /app
COPY --from=builder /app/api-gateway .

ENTRYPOINT ["./api-gateway"]
CMD ["--config", "config.yml"]
EXPOSE 8080
