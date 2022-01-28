GOBIN := $(shell pwd)/bin
SRC := $(shell find . -type f -name '*.go' -path "./pkg/f3/*")
VERSION := $(shell cat pkg/f3/version.go |grep "const Version ="|cut -d"\"" -f2)
GIT_COMMIT := $(shell git rev-parse HEAD)
LIBNAME := libf3
EXENAME := f3

build: FLAGS := -ldflags "-X github.com/xeus2001/interview-accountapi/pkg/f3.DefaultEndPoint=http://localhost:8080/v1"
release: FLAGS := -ldflags "-X github.com/xeus2001/interview-accountapi/pkg/f3.DefaultEndPoint=https://api.f3.tech/v1"

build: do-build
release: do-build

do-build:
	@echo "GOPATH: $(GOPATH)"
	@echo "LDFLAGS: $(FLAGS)"
	@echo "FILES: $(SRC)"
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go build $(FLAGS) -o bin/$(LIBNAME) cmd/main.go
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) GOOS=linux GOARCH=amd64 go build $(FLAGS) -o bin/$(EXENAME) cmd/main.go
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) GOOS=windows GOARCH=amd64 go build $(FLAGS) -o bin/$(EXENAME).exe cmd/main.go

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
	@rm -f bin/$(EXENAME)
	@rm -f bin/$(EXENAME).exe
	@docker image rm f3.int.test:latest 2>/dev/null || true

simplify:
	@gofmt -s -l -w $(SRC)

test:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go test -v github.com/xeus2001/interview-accountapi/pkg/f3

test-int:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go test -tags=int -v github.com/xeus2001/interview-accountapi/pkg/f3 -f3.endpoint=http://localhost:8080/v1

.PHONY: all build release do-build docs open-swagger-ui fmt check get clean simplify test test-int
