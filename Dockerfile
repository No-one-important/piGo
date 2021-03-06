## Build stage
FROM golang:1.18-alpine AS build-env
ADD . /app
WORKDIR /app
RUN env CGO_ENABLED=0 go build -ldflags="-s -w" -o piGo main.go

## Create image
FROM scratch
COPY --from=build-env /app/piGo /
ENTRYPOINT ["/piGo"]