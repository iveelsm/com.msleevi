.PHONY: clean build

clean:
	rm -rf bin
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go clean ./api/main.go

build:
	(cd scripts && ./install_requirements.sh)
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go build ./api/main.go
	mkdir bin
	mv main bin/api

run:
	@GOPATH=$(GOPATH) @GOBIN=$(GOBIN) go run ./api/main.gp