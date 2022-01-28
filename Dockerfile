FROM golang:1.18beta1

LABEL "env"="local"
LABEL "author"="Alexander Lowey-Weber <alexander@lowey.family>"

COPY pkg /go/src/pkg
COPY go.mod /go/src/.

WORKDIR /go/src

RUN go get github.com/google/uuid

WORKDIR /go/src/pkg/f3

ENTRYPOINT ["go", "test", "-v", "-tags=int"]
