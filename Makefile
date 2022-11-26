LDFLAGS = -s
LDFLAGS += -w
LDFLAGS += -X 'main.Version=$(shell git describe --tags --abbrev=0)'
LDFLAGS += -X 'main.Commit=$(shell git rev-list -1 HEAD)'

dev:
	@if [ -f "./columbus" ]; then rm columbus; fi  
	go build -o columbus --race -ldflags="$(LDFLAGS)" .

clean:
	@if [ -f "./columbus" ]; then rm columbus; fi  
	@if [ -f "./columbus-linux-amd64" ]; then rm columbus-linux-amd64; fi
	@if [ -f "./columbus-linux-arm64 " ]; then rm columbus-linux-arm64 ; fi
	@if [ -f "./columbus-darwin-arm64" ]; then rm columbus-darwin-arm64; fi  
	@if [ -f "./columbus-windows-amd64" ]; then rm columbus-windows-amd64; fi  

build-linux-amd64:
	GOOS=linux   GOARCH=amd64 go build -o columbus-linux-amd64   -ldflags="$(LDFLAGS)" .

build-linux-arm64:
	GOOS=linux   GOARCH=arm64 go build -o columbus-linux-arm64   -ldflags="$(LDFLAGS)" .

build-linux: build-linux-amd64 build-linux-arm64

build-darwin-arm64:
	GOOS=darwin  GOARCH=arm64 go build -o columbus-darwin-arm64  -ldflags="$(LDFLAGS)" .

build-darwin-amd64:
	GOOS=darwin  GOARCH=amd64 go build -o columbus-darwin-amd64  -ldflags="$(LDFLAGS)" .

build-darwin: build-darwin-arm64 build-darwin-amd64

build-windows-amd64:
	GOOS=windows GOARCH=amd64 go build -o columbus-windows-amd64 -ldflags="$(LDFLAGS)" .

build-windows: build-windows-amd64

build-all: build-linux build-darwin build-windows

release: clean build-all
	sha512sum columbus* | gpg --clearsign -u daniel@elmasy.com > checksums