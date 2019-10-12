build: build-container
	docker run flavium/dashboard:build-1.0 -v .:/tmp/app-dir/ -w /tmp/app-dir run build

build-container:
	docker build -f build-Dockerfile -t flavium/dashboard:build-1.0 .
