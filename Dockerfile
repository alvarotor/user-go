FROM golang:1.23.2-alpine3.20 AS base
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

FROM base AS test-stage
WORKDIR /app
COPY . .
RUN go test -v ./...

FROM test-stage AS build-stage
WORKDIR /app/server
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /server main.go config.go

FROM alpine:3.20.3 AS build-release-stage
RUN apk update && apk upgrade && apk add --no-cache curl && rm -rf /var/cache/apk/*
COPY --from=build-stage /server /server

ENTRYPOINT ["/server"]
