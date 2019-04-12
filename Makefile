GOPATH=$(shell pwd)/vendor:$(shell pwd)
GOBIN=$(shell pwd)/bin
GOFILES=$(wildcard *.go)
GONAME=$(shell basename "$(PWD)")
GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)
VERSION=$(shell cat pkg/version.go |grep "const Version ="|cut -d"\"" -f2)
GIT_COMMIT=$(git rev-parse HEAD)
TARGET := $(shell echo $${PWD\#\#*/})
PID=/tmp/go-$(GONAME).pid

# Use linker flags to provide version/build settings to the target
LDFLAGS=-ldflags "-X=main.Version=$(VERSION) -X=main.Build=$(GIT_COMMIT)"

# go source files, ignore vendor directory
SRC = $(shell find . -type f -name '*.go' -not -path "./vendor/*")

$(TARGET): $(SRC)
	@go build $(LDFLAGS) -o $(TARGET)

build:
  @GOPATH=$(GOPATH) GOBIN=$(GOBIN) go build $(LDFLAGS) -o bin/$(GONAME) $(GOFILES)

check:
	@test -z $(shell gofmt -l cmd/scif/cli.go | tee /dev/stderr) || echo "[WARN] Fix formatting issues with 'make fmt'"
	@for d in $$(@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go list ./... | grep -v /vendor/); do golint $${d}; done
	@go tool vet ${SRC}

deps:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go get -u golang.org/x/lint/golint

# dev creates binaries for testing scif locally. These are put
# into ./bin/ as well as $GOPATH/bin
dev: fmtcheck
	go install -mod=vendor .

fmt:
	gofmt -w $(GOFMT_FILES)

fmtcheck:
	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"

get:
  @GOPATH=$(GOPATH) GOBIN=$(GOBIN) go get .

install:
  #@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go install $(GOFILES)
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) @go install $(LDFLAGS)

uninstall: clean
	@rm -f $$(which ${TARGET})

clean:
	@rm -f $(TARGET)

simplify:
	@gofmt -s -l -w $(SRC)

.PHONY: all build check clean deps dev fmt fmtcheck get install uninstall simplify quickdev run
