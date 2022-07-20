FROM golang:1.18.4-alpine3.15 as build

ENV GOPROXY=https://goproxy.io

ADD . /address

WORKDIR /address

RUN go build -o cmd/main cmd/address/main.go

FROM alpine:3.15

ENV GIN_MODE="release"

WORKDIR /www

COPY --from=build /address/cmd/main /usr/bin/main
COPY ./data /data
COPY ./data /usr/bin/data

RUN chmod +x /usr/bin/main
RUN chmod +x /usr/bin/data

ENV TZ=Aisa/Shanghai
RUN ln -snf /usr/shar/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

ENTRYPOINT ["/usr/bin/main"]