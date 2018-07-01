.PHONY: clean build run

clean:
	rm -rf bin
	mkdir bin
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go clean ./api/*.go

build: clean
	(cd scripts && ./install_requirements.sh)
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go build ./api/*.go
	mv api bin/api

run: build
	@GOPATH=$(GOPATH) @GOBIN=$(GOBIN) go run ./api/*.go
