HN_URL := 'https://news.ycombinator.com/item?id=3078128'
TW_URL := 'https://twitter.com/TwitterDev/status/1443269993676763138'
GO_FILES = $(shell find . -type f -name '*.go')

.PHONY: all install clean test testrun

all: twitterview hackerview

hackerview: $(GO_FILES)
	go build -race -v -ldflags="-s -w" ./cmd/$@
	ls -lh $@

twitterview: $(GO_FILES)
	go build -v -ldflags="-s -w" ./cmd/$@
	ls -lh $@

install: hackerview twitterview
	mv hackerview  $(HOME)/go/bin/
	mv twitterview $(HOME)/go/bin/

clean: ; go clean -x ./...
test:  ; go test -vet=all -v ./...

testrun:
	go run cmd/hackerview/main.go -l 10 -t 20s $(HN_URL)
	go run cmd/twitterview/main.go -t 20s $(TW_URL)
