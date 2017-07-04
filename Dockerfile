FROM golang:alpine

RUN apt-get update && apt-get install -y ca-certificates git-core ssh

ADD . /go/src/github.com/twogg-git/go-online

RUN go get github.com/twogg-git/go-online

RUN go install github.com/twogg-git/go-online
