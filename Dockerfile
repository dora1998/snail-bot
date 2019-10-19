FROM golang:1.13 AS build-env

ENV GO111MODULE on

WORKDIR /go/src/github.com/dora1998/snail-bot

ADD . /go/src/github.com/dora1998/snail-bot

RUN CGO_ENABLED=0 GOOS=linux go install -v \
    -ldflags="-w -s" \
    github.com/dora1998/snail-bot

EXPOSE 8080
ENTRYPOINT ["/go/bin/snail-bot"]
