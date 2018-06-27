from golang:alpine

RUN apk update && apk upgrade && apk add --no-cache git

ENV GOBIN=/go/bin
ENV GOOS=darwin
ENV GOARCH=amd64
#ENV GOMAXPROCS=1

ADD ./src /go

WORKDIR /go

CMD go build -v -o $GOPATH/bin/go-test
