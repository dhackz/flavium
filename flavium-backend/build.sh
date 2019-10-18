#!/bin/sh
export PATH=$PATH:$(go env GOPATH)/bin
go mod download
protoc -I/usr/local/include -I. \
    -I$(go env GOPATH)/src \
    -I$(go env GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --go_out=plugins=grpc:. \
  --grpc-gateway_out=logtostderr=true:. pkg/torrents/torrent.proto

go build -v pkg/cmd/main.go
