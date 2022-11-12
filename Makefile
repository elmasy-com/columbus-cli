LDFLAGS = -s
LDFLAGS += -w
LDFLAGS += -X 'main.Version=$(shell git describe --tags --abbrev=0)'
LDFLAGS += -X 'main.Commit=$(shell git rev-list -1 HEAD)'

build:
	go build -o columbus -ldflags="$(LDFLAGS)" .

release: build
	# Use goreleaser!
	sha512sum columbus > columbus.sha