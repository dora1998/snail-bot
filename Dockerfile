FROM golang:1.13 AS build-env

ENV GO111MODULE on

WORKDIR /go/src/github.com/dora1998/snail-bot

ADD . /go/src/github.com/dora1998/snail-bot

RUN CGO_ENABLED=0 GOOS=linux go install -v \
    -ldflags="-w -s" \
    github.com/dora1998/snail-bot/cmd/server

FROM alpine:latest

COPY --from=build-env /go/src/github.com/dora1998/snail-bot/migrations /migrations
COPY --from=build-env /go/bin/server /snail-bot
RUN chmod a+x /snail-bot

EXPOSE 8080
ENTRYPOINT ["/snail-bot"]
