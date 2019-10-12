build: build-container
	docker run -v $(PWD):/tmp/app-dir/ -w /tmp/app-dir flavium/dashboard:build-1.0 run build

build-container:
	docker build -f build-Dockerfile -t flavium/dashboard:build-1.0 .
