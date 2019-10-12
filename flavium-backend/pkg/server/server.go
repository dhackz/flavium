package server

import (
	pb "../torrents"
	"context"
)
type TorrentServer struct {

}

func (t *TorrentServer) AddTorrent(context.Context, *pb.AddTorrentRequest) (*pb.AddTorrentResponse, error) {
	return &pb.AddTorrentResponse{Ok:true}, nil
}

func (t *TorrentServer) GetStatus(context.Context, *pb.GetStatusRequest) (*pb.GetStatusResponse, error) {
	return &pb.GetStatusResponse{
		Torrents: []*pb.TorrentStatus{
			&pb.TorrentStatus{
				MagnetLink:           "magnetLOLabc123",
				Size:                 100000,
				DoneSize:             50000,
				Status:               "downloading",
				Started:              true,
			},
			&pb.TorrentStatus{
				MagnetLink:           "magnetLOLabc1234",
				Size:                 100000,
				DoneSize:             75000,
				Status:               "downloading",
				Started:              true,
			},
			&pb.TorrentStatus{
				MagnetLink:           "magnetLOLabc12345",
				Size:                 100000,
				DoneSize:             25000,
				Status:               "downloading",
				Started:              true,
			},
		},
	}, nil
}
