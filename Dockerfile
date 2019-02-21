FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git && mkdir sql
COPY . $GOPATH/src/gitlab.com/zenport.io/go-assignment
WORKDIR $GOPATH/src/gitlab.com/zenport.io/go-assignment
RUN go get -u github.com/golang/dep/cmd/dep && dep ensure
ENTRYPOINT ["go", "run", "main.go"]