FROM golang:1.9.2

ENV GOPATH /go:$GOPATH
ENV PATH /go/bin:$PATH

ADD . /go/src
WORKDIR /go/src

RUN go get github.com/dispatchlabs/disgo
ENTRYPOINT cd github.com/dispatchlabs/disgo && go run main.go
