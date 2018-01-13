.PHONY: clean test build dev deploy

clean:

test:
	yarn test

build:
	yarn prod:build
	$(MAKE) package -C build

dev:
	yarn dev:start

deploy:
