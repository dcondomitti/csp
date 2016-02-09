FROM golang:1.5.3

ENV GO15VENDOREXPERIMENT=1

RUN mkdir -p /go/src/github.com/dcondomitti/csp
WORKDIR /go/src/github.com/dcondomitti/csp

ADD . /go/src/github.com/dcondomitti/csp

RUN go build -ldflags "-X main.revision=$(git rev-parse HEAD)" -o /usr/bin/csp *.go

EXPOSE 8080

ENTRYPOINT ["/usr/bin/csp"]
