.PHONY: clean build

clean:
	rm -rf bin
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go clean ./main/Main.go

build:
	(cd scripts && ./install_requirements.sh)
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go build ./main/Main.go
	mkdir bin
	mv Main bin/
