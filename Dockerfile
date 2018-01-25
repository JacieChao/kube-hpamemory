FROM golang:1.9.2

RUN go get -u github.com/golang/dep/cmd/dep

RUN dep ensure -vendor-only -v
