.PHONY: clean build run

clean:
	rm -rf bin
	mkdir bin
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go clean ./github.com/msleevi/com.msleevi/cmd/main.go

build: clean
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go build ./github.com/msleevi/com.msleevi/cmd/main.go

run:
	@GOPATH=$(GOPATH) go run ./api/main.go
