GOBIN := $(shell pwd)/bin
SRC := $(shell find . -type f -name '*.go' -path "./pkg/f3/*")
VERSION := $(shell cat pkg/f3/version.go |grep "const Version ="|cut -d"\"" -f2)
LIBNAME := libf3
EXENAME := f3

build: FLAGS := -ldflags "-X github.com/xeus2001/interview-accountapi/pkg/f3.DefaultEndPoint=http://localhost:8080/v1"
docker: FLAGS := -ldflags "-X github.com/xeus2001/interview-accountapi/pkg/f3.DefaultEndPoint=http://accountapi:8080/v1"
release: FLAGS := -ldflags "-X github.com/xeus2001/interview-accountapi/pkg/f3.DefaultEndPoint=https://api.f3.tech/v1"

build: do-build
docker: do-build
release: check do-build

do-build: get
	@echo "GOPATH: $(GOPATH)"
	@echo "LDFLAGS: $(FLAGS)"
	@echo "FILES: $(SRC)"
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go build $(FLAGS) -o bin/$(LIBNAME) cmd/main.go
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) GOOS=linux GOARCH=amd64 go build $(FLAGS) -o bin/$(EXENAME) cmd/main.go
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) GOOS=windows GOARCH=amd64 go build $(FLAGS) -o bin/$(EXENAME).exe cmd/main.go

doc: fmt
	gomarkdoc --output doc/f3.md pkg/f3/*.go

swagger-ui:
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
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go get github.com/google/uuid
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go get golang.org/x/tools/cmd/cover

clean:
	@rm -f bin/$(LIBNAME)
	@rm -f bin/$(EXENAME)
	@rm -f bin/$(EXENAME).exe
	@rm -f coverage.out
	@docker image rm f3.int.test:latest 2>/dev/null || true

simplify:
	@gofmt -s -l -w $(SRC)

test:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go test -cover -v github.com/xeus2001/interview-accountapi/pkg/f3

test-int:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go test -cover -coverprofile=coverage.out -v -tags=int github.com/xeus2001/interview-accountapi/pkg/f3 -f3.endpoint=http://localhost:8080/v1

test-int-result:
	@go tool cover -html=coverage.out

test-docker:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go test -cover -v -tags=int github.com/xeus2001/interview-accountapi/pkg/f3 -f3.endpoint=http://accountapi:8080/v1

.PHONY: all build release do-build doc swagger-ui fmt check get clean simplify test test-int
