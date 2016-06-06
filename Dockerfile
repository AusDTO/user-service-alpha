FROM golang:latest

ADD . /go/src/github.com/ausdto/user-service
WORKDIR /go/src/github.com/ausdto/user-service

RUN go get
RUN go build -o app
CMD sleep 5 && ./app

