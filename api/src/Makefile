.PHONY: clean build

clean:
	rm -rf bin
	go clean ./main/Main.go

build:
	go build ./main/Main.go
	mkdir bin
	mv Main bin/
