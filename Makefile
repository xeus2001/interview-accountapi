# PATH := $(PATH):$(shell echo $$GOPATH/bin)
GOBIN := $(shell pwd)/bin
LIBNAME := libf3
EXECNAME := f3
SRC := $(shell find . -type f -name '*.go' -path "./pkg/f3/*")
VERSION := $(shell cat pkg/f3/version.go |grep "const Version ="|cut -d"\"" -f2)
GIT_COMMIT := $(shell git rev-parse HEAD)
TARGET := $(shell echo $${PWD})
LDFLAGS=-ldflags "-X=main.Version=$(VERSION) -X=main.Build=$(GIT_COMMIT)"

$(TARGET): $(SRC)
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go build $(LDFLAGS) -o $(TARGET)

build:
	@echo "GOPATH: $(GOPATH)"
	@echo "LDFLAGS: $(LDFLAGS)"
	@echo "FILES: $(SRC)"
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go build $(LDFLAGS) -o bin/$(LIBNAME) cmd/main.go
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o bin/$(EXECNAME) cmd/main.go
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o bin/$(EXECNAME).exe cmd/main.go

docs: fmt
	gomarkdoc --output doc/f3.md pkg/f3/*.go

open-swagger-ui:
	chromium-browser \
      --disable-web-security \
      --user-data-dir="/tmp/chromium-debug/" \
      'http://localhost:7080/#/Health/get_health' 'http://localhost:7080/#/Health/get_health'

fmt:
	@echo $(SRC)
	gofmt -w $(SRC)

check:
	@sh -c "'$(CURDIR)/scripts/fmtcheck.sh'"

get:
	go get -u github.com/google/uuid

clean:
	@rm -f bin/$(LIBNAME)
	@rm -f bin/$(EXECNAME)
	@rm -f bin/$(EXECNAME).exe

simplify:
	@gofmt -s -l -w $(SRC)

test:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go test -v github.com/xeus2001/interview-accountapi/pkg/f3

test-int:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go test -tags=int -v github.com/xeus2001/interview-accountapi/pkg/f3 -f3.endpoint=http://localhost:8080/v1

.PHONY: all build docs fmt check get simplify test test-int open-swagger-ui
