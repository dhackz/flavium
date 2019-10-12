package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	pb "../torrents"
	server "../server"
)

const (
	port = ":8081"
)

func main(){
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTorrentServer(s, &server.TorrentServer{})
	print("Serving...\n")
	if err := s.Serve(lis); err != nil {
		log.Fatal("failed to serve: %v", err)
	}
}
