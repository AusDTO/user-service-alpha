FROM golang:latest

ADD . /go/src/github.com/ausdto/user-seed
WORKDIR /go/src/github.com/ausdto/user-seed

RUN go get
RUN go build -o seed
CMD sleep 5 && ./seed