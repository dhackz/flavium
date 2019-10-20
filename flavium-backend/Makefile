mkfile_path=$(abspath $(lastword $(MAKEFILE_LIST)))
mkfile_dir=$(abspath $(dir mkfile_path))

proto_files=$(wildcard pkg/**/*.proto)
proto_pb_files=$(proto_files:.proto=.pb.go)
proto_pb_gw_files=$(proto_files:.proto=.pb.gw.go)
proto_targets=$(proto_pb_files) $(proto_pb_gw_files)
$(info $(proto_targets))

main: $(proto_targets) build-container
	docker run \
		--user $(shell id -u):$(shell id -g) \
		-v $(mkfile_dir):/tmp/app-dir \
		-w /tmp/app-dir \
		flavium/backend:build-1.0 \
		go build -v pkg/cmd/main.go

gen: $(proto_targets) build-container

%.pb.go %.pb.gw.go: %.proto build-container 
	docker run \
		--user $(shell id -u):$(shell id -g) \
		-v $(mkfile_dir):/tmp/app-dir \
		-w /tmp/app-dir \
		flavium/backend:build-1.0 \
		sh -c "PATH=$$PATH:/tmp/go/bin protoc -I/usr/local/include -I. \
			-I/tmp/go/src \
			-I/tmp/go/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
			--go_out=plugins=grpc:. \
			--grpc-gateway_out=logtostderr=true:. $<"

build-container:
	docker build \
		--build-arg UID=$(shell id -u) \
		--build-arg GID=$(shell id -g) \
		--build-arg UNAME=$(shell whoami) \
		-f build-Dockerfile \
		-t flavium/backend:build-1.0 \
		.

.PHONY: gen
