FROM golang:1.18beta1

LABEL "env"="local"
LABEL "author"="Alexander Lowey-Weber <alexander@lowey.family>"

COPY pkg /go/src/pkg
COPY go.mod /go/src/.
COPY Makefile /go/src/.

WORKDIR /go/src
RUN make get

ENTRYPOINT ["make", "test-docker"]
