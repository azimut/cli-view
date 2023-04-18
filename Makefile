GO_FILES := $(shell find . -type f -name '*.go')
BINARIES := twitterview hackerview redditview fourchanview

.PHONY: all install clean test

all: $(BINARIES)

$(BINARIES): $(GO_FILES)
	go build -v -ldflags="-s -w" ./cmd/$@
	ls -lh $@

install: $(BINARIES);
	 mv $(BINARIES) $(HOME)/go/bin/

clean: ; go clean -x ./...
test:  ; go test -vet=all -v -race ./...
