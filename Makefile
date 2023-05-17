GO_FILES := $(shell find . -type f -name '*.go')
BINARIES := twitterview hackerview redditview fourchanview vichanview discourseview lobstersview
LDFLAGS  := -ldflags="-s -w"

ifdef DEBUG
undefine LDFLAGS
endif

.PHONY: all install clean test

all: $(BINARIES)

$(BINARIES): $(GO_FILES)
	go build -v $(LDFLAGS) ./cmd/$@
	ls -lh $@

install: test $(BINARIES)
	 mv $(BINARIES) $(HOME)/go/bin/

clean: ; go clean -x ./...
test:  ; go test -vet=all -race ./...
