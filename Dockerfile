FROM golang:1.9.2

RUN go get -u github.com/golang/dep/cmd/dep

RUN mkdir -p /go/src/kube-hpamemory/

COPY . /go/src/kube-hpamemor/
WORKDIR /go/src/kube-hpamemor/

RUN dep ensure -vendor-only -v
