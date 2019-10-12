mkfile_path=$(abspath $(lastword $(MAKEFILE_LIST)))
mkfile_dir=$(abspath $(dir mkfile_path))

build: build-container
	docker run \
		--user $(shell id -u):$(shell id -g) \
		-v $(mkfile_dir):/tmp/app-dir \
		-w /tmp/app-dir \
		flavium/backend:build-1.0 \
		/bin/sh build.sh

build-container:
	docker build \
		--build-arg UID=$(shell id -u) \
		--build-arg GID=$(shell id -g) \
		--build-arg UNAME=$(shell whoami) \
		-f build-Dockerfile \
		-t flavium/backend:build-1.0 \
		.
