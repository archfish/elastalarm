FROM golang:alpine as builder

WORKDIR $GOPATH/src/github.com/turnon/elastalarm
COPY . ./

RUN apk add --no-cache git \
    && export GO111MODULE=on \
    && go get ./... \
    && go build -o /elastalarm \
    && apk del git

FROM alpine:latest

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
WORKDIR /usr/bin
COPY --from=builder /elastalarm .

ENTRYPOINT ["/usr/bin/elastalarm", "-configs", "/configs"]