FROM golang:1.14.0-alpine3.11 as build

# ENV GOPROXY=https://goproxy.io

ADD . /address

WORKDIR /address

RUN go build -o cmd/main cmd/address/main.go

FROM alpine:3.7

ENV GIN_MODE="release"

WORKDIR /www

COPY --from=build /address/cmd/main /usr/bin/main
COPY ./data /data
COPY ./data /usr/bin/data

RUN chmod +x /usr/bin/main
RUN chmod +x /usr/bin/data

ENTRYPOINT ["/usr/bin/main"]