package server

import (
	pb "../torrents"
	"context"
	"fmt"
	"os"
	"os/exec"
)

type TorrentServer struct {
	IsDryRun bool
}

func (t *TorrentServer) AddTorrent(_ context.Context, req *pb.AddTorrentRequest) (*pb.AddTorrentResponse, error) {
	cmd := fmt.Sprintf("transmission-remote %s -a %s", os.Getenv("TRANSMISSION_HOST"), req.MagnetLink)
	if t.IsDryRun {
		fmt.Println("DRYRUN: " + cmd)
	} else {
		fmt.Println("RUNNING: " + cmd)

		exe := exec.Command("transmission-remote", os.Getenv("TRANSMISSION_HOST"),"-a", req.MagnetLink)

		err := exe.Start()
		if err != nil{
			fmt.Printf(err.Error())
		}

		err = exe.Wait()
		if err != nil{
			fmt.Printf(err.Error())
		}
	}
	return &pb.AddTorrentResponse{Ok:true}, nil
}

func (t *TorrentServer) GetStatus(context.Context, *pb.GetStatusRequest) (*pb.GetStatusResponse, error) {
	cmd := fmt.Sprintf("transmission-remote %s -l",  os.Getenv("TRANSMISSION_HOST"))
	if t.IsDryRun {
		fmt.Println("DRYRUN: " + cmd)
	} else {
		fmt.Println("RUNNING: " + cmd)

		output, err := exec.Command("transmission-remote",os.Getenv("TRANSMISSION_HOST"),"-l").Output()
		if err != nil{
			fmt.Printf(err.Error())
		}

		fmt.Printf("CMD output: %s", output)

	}
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
