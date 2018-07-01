.PHONY: clean build deploy

GOPATH=$(shell pwd)
GOBIN=$(shell pwd)/bin

init:
	sudo gem install fpm

clean:
	rm -rf pkg
	$(MAKE) -C src clean GOPATH=$(GOPATH)	GOBIN=$(GOBIN)
	$(MAKE) -C package clean

build: clean
	$(MAKE) -C src build GOPATH=$(GOPATH)	GOBIN=$(GOBIN)
	$(MAKE) -C package build

run:
	$(MAKE) -C src run GOPATH=$(GOPATH)

deploy:
	sudo yum install --assumeyes package/build/dist/com-msleevi-api-*.rpm
