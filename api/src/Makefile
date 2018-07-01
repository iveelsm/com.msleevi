.PHONY: clean build run

clean:
	rm -rf bin
	mkdir bin
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go clean ./api/main.go

build: clean
	(cd scripts && ./install_requirements.sh)
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go build ./api/main.go
	mv main bin/api

run:
	@GOPATH=$(GOPATH) go run ./api/main.go
