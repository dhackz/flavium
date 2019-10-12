build: build-container
	docker run flavium/backend:build-1.0 -v .:/tmp/app-dir -w /tmp/app-dir build -v

build-container:
	docker build -f build-Dockerfile . -t flavium/backend:build-1.0
