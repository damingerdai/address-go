FROM golang:1.13.5-alpine3.10 as build

# ENV GOPROXY=https://goproxy.io

ADD . /address

WORKDIR /address

RUN go build -o cmd/main cmd/address/main.go

FROM alpine:3.7

ENV GIN_MODE="release"

WORKDIR /www

COPY --from=build /address/cmd/main /usr/bin/main
COPY ./data /usr/bin
COPY ./.env /usr/bin

RUN chmod +x /usr/bin/main
RUN chmod +x /usr/bin/.env

ENTRYPOINT ["/usr/bin/main"]