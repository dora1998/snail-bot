FROM golang:1.13 AS build-env

ENV GO111MODULE on

WORKDIR /go/src/github.com/dora1998/snail-bot

ADD . /go/src/github.com/dora1998/snail-bot

RUN CGO_ENABLED=0 GOOS=linux go install -v \
    -ldflags="-w -s" \
    github.com/dora1998/snail-bot/cmd/server

FROM alpine:latest

RUN apk --update add tzdata && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime

RUN apk add bash ca-certificates curl mariadb-client
COPY --from=build-env /go/src/github.com/dora1998/snail-bot/migrations /migrations
COPY --from=build-env /go/bin/server /snail-bot
RUN chmod a+x /snail-bot

COPY ./docker-entrypoint.sh /docker-entrypoint.sh
RUN chmod a+x /docker-entrypoint.sh

EXPOSE 8080
ENTRYPOINT ["/docker-entrypoint.sh"]
CMD ["/snail-bot"]
