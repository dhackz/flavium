build: build-container
	docker run flavium/dashboard:build-1.0 -v .:/tmp/app_dir/ -w /tmp/app_dir run build

build-container:
	docker build -f build-Dockerfile -t flavium/dashboard:build-1.0 .
