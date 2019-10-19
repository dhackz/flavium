package server

import (
	pb "flavium-backend/pkg/torrents"
	"context"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"
)

const TRANSMISSION_BODY_EXPRESSION = "^\\s+" +
"(?P<id>\\d+\\*?)" +
"\\s+" +
"(?P<done>(?:\\d+%|n/a))" +
"\\s+" +
"(?P<have>\\d+.\\d+ [MGKmkg]B|None)" +
"\\s+" +
"(?P<eta>(?:\\d+ (?:min|hrs|sec))|Done|Unknown)" +
"\\s+" +
"(?P<up>\\d+.\\d+)" +
"\\s+" +
"(?P<down>\\d+.\\d+)" +
"\\s+" +
"(?P<ratio>\\d+.\\d+|None)" +
"\\s+" +
"(?P<status>(Up & Down)|\\w+)" +
"\\s+" +
"(?P<name>.+)$"


var TRANSMISSION_BODY_PARSER = regexp.MustCompile(TRANSMISSION_BODY_EXPRESSION)

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

func getTorrentStatus(transmissionOutput string) []*pb.TorrentStatus {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from", r)
        }
    }()
    outputLines := strings.Split(transmissionOutput, "\n")
    fmt.Printf("%+v\n", outputLines)
    fmt.Println(len(outputLines))
    if len(outputLines) > 2 {
        outputLines = outputLines[1:len(outputLines)-2]
    } else {
        return nil
    }
    torrents := make([]*pb.TorrentStatus, len(outputLines))
    fmt.Println(len(outputLines))
    for i, line := range outputLines {
        fmt.Printf("Parsing line: \"%s\"\n", line)
        match := TRANSMISSION_BODY_PARSER.FindStringSubmatch(line)
        result := make(map[string]string)
        for j, name := range TRANSMISSION_BODY_PARSER.SubexpNames() {
            if j != 0 && name != "" {
                result[name] = match[j]
                fmt.Printf("%s - %s\n", name, result[name])
            }
        }
        fmt.Println("Adding Id")
        fmt.Println(len(torrents))
        torrents[i] = &pb.TorrentStatus {
            Id : result["id"],
            Done : result["done"],
            Have : result["have"],
            Eta : result["eta"],
            Up : result["up"],
            Down : result["down"],
            Result : result["result"],
            Status : result["status"],
            Name : result["name"],
        }
    }
    fmt.Printf("Parsed %d lines\n", len(outputLines))
    return torrents
}

func (t *TorrentServer) GetStatus(context.Context, *pb.GetStatusRequest) (*pb.GetStatusResponse, error) {
	cmd := fmt.Sprintf("transmission-remote %s -l",  os.Getenv("TRANSMISSION_HOST"))
	if t.IsDryRun {
		fmt.Println("DRYRUN: " + cmd)
		return &pb.GetStatusResponse{}, nil
	} else {
		fmt.Println("RUNNING: " + cmd)

		output, err := exec.Command("transmission-remote",os.Getenv("TRANSMISSION_HOST"),"-l").Output()
        fmt.Printf("CMD output: %s\n", output)
		if err != nil{
			fmt.Println(err.Error())
		}
		torrents := getTorrentStatus(string(output))
		return &pb.GetStatusResponse{
			Torrents: torrents,
		}, nil


	}

	return &pb.GetStatusResponse{}, nil
}


func ScheduleTorrentListener(delay time.Duration) {
	go func() {
		for {
			output, err := exec.Command("transmission-remote",os.Getenv("TRANSMISSION_HOST"),"-l").Output()
			if err != nil{
				fmt.Println(err.Error())
			}
			torrents := getTorrentStatus(string(output))
			for i := range torrents {
				if torrentIsFinished(*torrents[i]) {
					err := exec.Command("rsync", "-r", "/var/lib/flavium/downloads/complete/"+torrents[i].Name, "/tmp").Run()
					if err != nil{
						fmt.Printf("Copy failed: %v\n", err)
					}else{
					// run filebot
					output, err := exec.Command("filebot", "-rename", "/tmp/"+torrents[i].Name).Output()
					fmt.Printf("Filebot ouput '%s'", output)
					if err != nil{
						fmt.Printf("Filebot failed: %v\n", err)
					}else{
						//check if movie or series
						// move to plex
						err = exec.Command("mv", "/tmp/"+torrents[i].Name, "/var/lib/plex/data").Run()
						if err != nil{
							fmt.Printf("Move failed: %v\n", err)
						} else {
							// remove torrent and delete data
							err := exec.Command("transmission-remote", os.Getenv("TRANSMISSION_HOST"), "--torrent", torrents[i].Id, "--remove-and-delete").Run()
							if err != nil{
								fmt.Println(err.Error())
							}
						}
					}
					}
				}
			}
			time.Sleep(delay)
		}
	}()
}

func torrentIsFinished(torrent pb.TorrentStatus) bool {
	return torrent.Done == "100%" && (torrent.Status == "Finished" || torrent.Status == "Idle" || torrent.Status == "Seeding" || torrent.Status == "Stopped")
}