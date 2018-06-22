FROM golang:1.10-alpine
WORKDIR /go/src/github.com/arachnys/athenapdf/weaver

RUN apk add --update git
RUN go get -u github.com/golang/dep/cmd/dep

COPY Gopkg.lock Gopkg.toml ./
RUN dep ensure --vendor-only -v

COPY . ./

RUN \
  CGO_ENABLED=0 go build -v -o weaver .

CMD ["/bin/sh"]
