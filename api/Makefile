.PHONY: clean build deploy

clean:
	$(MAKE) -c src clean
	$(MAKE) -c package clean

build:
	$(MAKE) -C src build
	$(MAKE) -C package build

run:
	$(MAKE) src run

deploy:
	echo "Deploying API"
