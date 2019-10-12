mkfile_path=$(abspath $(lastword $(MAKEFILE_LIST)))
mkfile_dir=$(abspath $(dir mkfile_path))

$(info $(mkfile_path))
$(info $(mkfile_dir))

build: build-container
	docker run \
		--user $(shell id -u):$(shell id -g) \
		-v $(mkfile_dir):/tmp/app-dir/ \
		-w /tmp/app-dir \
		flavium/dashboard:build-1.0 \
		ash -c "npm install && npm run build"

build-container:
	docker build \
		--build-arg UID=$(shell id -u) \
		--build-arg GID=$(shell id -g) \
		--build-arg UNAME=$(shell whoami) \
		-f build-Dockerfile \
		-t flavium/dashboard:build-1.0 \
		.
