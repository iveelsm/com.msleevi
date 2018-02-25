.PHONY: clean build deploy

GOPATH=$(shell pwd)
GOBIN=$(shell pwd)/bin

init:
	sudo gem install fpm

clean:
	$(MAKE) -C src clean GOPATH=$(GOPATH)	GOBIN=$(GOBIN)
	$(MAKE) -C package clean

build:
	$(MAKE) -C src build GOPATH=$(GOPATH)	GOBIN=$(GOBIN)
	$(MAKE) -C package build

run:
	$(MAKE) src run

deploy:
	echo "Deploying API"
