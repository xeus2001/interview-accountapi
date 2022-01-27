FROM golang:1.18beta1

LABEL "env"="local"
LABEL "author"="Alexander Lowey-Weber <alexander@lowey.family>"

COPY *.go /go/src/
COPY go.mod /go/src/.

WORKDIR /go/src/

RUN go get github.com/google/uuid

ENTRYPOINT ["go", "test", "-v", "-tags=int"]