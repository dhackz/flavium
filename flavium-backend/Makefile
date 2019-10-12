build: build-container
	docker run -v $(PWD):/tmp/app-dir -w /tmp/app-dir flavium/backend:build-1.0 build -v

build-container:
	docker build -f build-Dockerfile . -t flavium/backend:build-1.0
