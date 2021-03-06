FROM golang:1.14

RUN mkdir -p /go/src/github.com/yaoshipu/shared-library/

WORKDIR /go/src/github.com/yaoshipu/shared-library/

COPY ./cmd/library ./cmd/library
COPY go.mod .
COPY go.sum .

RUN GOOS=linux OARCH=amd64 go build ./cmd/library

FROM debian

RUN mkdir -p /library

WORKDIR /library/

COPY --from=0 /go/src/github.com/yaoshipu/shared-library/library .

CMD ["/library/library"]
