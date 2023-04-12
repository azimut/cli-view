GO_FILES = $(shell find . -type f -name '*.go')

.PHONY: all install clean test testrun

all: twitterview hackerview redditview

hackerview:
redditview:
twitterview: $(GO_FILES)
	go build -v -ldflags="-s -w" ./cmd/$@
	ls -lh $@

install: hackerview twitterview redditview
	mv hackerview  $(HOME)/go/bin/
	mv twitterview $(HOME)/go/bin/
	mv redditview  $(HOME)/go/bin/

clean: ; go clean -x ./...
test:  ; go test -vet=all -v -race ./...
