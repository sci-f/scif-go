GOPATH := $(shell pwd)/vendor:$(shell pwd):$(shell echo $$GOPATH)
PATH := $(PATH):$(shell echo $$GOPATH/bin)
GOBIN := $(shell pwd)/bin
GONAME := scif
GOFILES?=$$(find . -name '*.go' | grep -v vendor)
VERSION := $(shell cat pkg/version/version.go |grep "const Version ="|cut -d"\"" -f2)
GIT_COMMIT := $(shell git rev-parse HEAD)
TARGET := $(shell echo $${PWD\#\#*/})
PID=/tmp/go-$(GONAME).pid

# Use linker flags to provide version/build settings to the target
LDFLAGS=-ldflags "-X=main.Version=$(VERSION) -X=main.Build=$(GIT_COMMIT)"

# go source files, ignore vendor directory
SRC = $(shell find . -type f -name '*.go' -not -path "./vendor/*")

$(TARGET): $(SRC)
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go build $(LDFLAGS) -o $(TARGET)

build:
	@echo "LDFLAGS: $(LDFLAGS)"
	@echo "GOFILES: $(GOFILES)"
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go build $(LDFLAGS) -o bin/$(GONAME) ./cmd/scif

deps:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go get -u golang.org/x/sys/unix
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go get -u golang.org/x/lint/golint
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go get -u github.com/spf13/cobra
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go get -u github.com/google/shlex

# dev creates binaries for testing scif locally. These are put
# into ./bin/ as well as $GOPATH/bin
dev: fmtcheck
	go install -mod=vendor .

fmt:
	gofmt -w $(GOFILES)

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

.PHONY: all build clean deps dev fmt fmtcheck get install uninstall simplify quickdev run
