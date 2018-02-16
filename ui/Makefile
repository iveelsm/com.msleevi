.PHONY: clean build dev deploy

clean:
	rm -rf build/

test:
	yarn test

build:
	yarn test
	yarn prod:build
	$(MAKE) package -C build

dev:
	yarn dev:start

deploy:
