FROM golang:alpine

RUN apk add --no-cache git mercurial

LABEL ddsaas.version="1.1.6" maintainer="Jeorch"

RUN go get github.com/alfredyang1986/blackmirror
RUN go get github.com/alfredyang1986/ddsaas

ADD deploy-config/ /go/bin/

RUN go install -v github.com/alfredyang1986/ddsaas

WORKDIR /go/bin

ENTRYPOINT ["ddsaas"]
