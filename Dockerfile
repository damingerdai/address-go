FROM golang:1.20-alpine3.17 as build

ENV GOPROXY=https://goproxy.cn

ADD . /address

WORKDIR /address

RUN go build -o cmd/main cmd/address/main.go

FROM alpine:3.17

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
