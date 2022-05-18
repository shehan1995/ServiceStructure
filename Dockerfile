FROM golang:1.17-alpine AS build

WORKDIR /src
RUN apk update && apk add --no-cache git

RUN go mod download
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o main

FROM alpine:3.10
RUN apk add --no-cache tzdata ca-certificates
COPY --from=build /src/main /main
CMD ["/main"]
