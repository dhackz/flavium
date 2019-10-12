#!/bin/sh
export PATH=$PATH:$GOPATH/bin
sh -c 'protoc -I/usr/local/include -I. \
  -I/tmp/go/src \
  -I/tmp/go/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --go_out=plugins=grpc:. \
  --grpc-gateway_out=logtostderr=true:. pkg/torrents/torrent.proto'

go build -v pkg/cmd/main.go
