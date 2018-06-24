from golang:alpine

RUN apk update && apk upgrade && apk add --no-cache git

ENV GOMAXPROCS=1

ADD ./src /go

WORKDIR /go

#RUN go get golang.org/x/crypto/ssh

CMD go run main.go
