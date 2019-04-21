GOPATH := $(shell pwd)/vendor:$(shell pwd):$(shell echo $$GOPATH)
PATH := $(PATH):$(shell echo $$GOPATH/bin)
GOBIN := $(shell pwd)/bin
GONAME := scif
GOFILES?=$$(find . -name '*.go' | grep -v vendor | grep -v _extra )
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
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go get -u github.com/mitchellh/gox

# dev creates binaries for testing scif locally. These are put
# into ./bin/ as well as $GOPATH/bin
dev: fmtcheck
	go install -mod=vendor .

docs: fmt
	godoc -notes="TODO|BUG" -http=:6060

fmt:
	gofmt -w $(GOFILES)

vet:
	go vet ./cmd/scif

fmtcheck:
	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"

get:
  @GOPATH=$(GOPATH) GOBIN=$(GOBIN) go get .

install:
	@echo "Installing to: $(GOBIN)"
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go install $(LDFLAGS) ./cmd/scif

uninstall: clean
	@rm -f $$(which ${TARGET})

clean:
	@rm -f $(TARGET)

release:
	@mkdir -p dist
	# for windows, need to install mousetrap and add windows back here
	gox -os="linux darwin" -arch="amd64" -output="dist/scif_{{.OS}}_{{.Arch}}" ./cmd/scif
	@cd dist/ && gzip *

simplify:
	@gofmt -s -l -w $(SRC)

test:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go test -v github.com/sci-f/scif-go/internal/pkg/logger

.PHONY: all build clean deps docs dev fmt fmtcheck get install uninstall vet simplify quickdev release test
