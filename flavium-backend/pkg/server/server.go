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

var execCommand = exec.Command

type TorrentServer struct {
	IsDryRun bool
}

func (t *TorrentServer) AddTorrent(_ context.Context, req *pb.AddTorrentRequest) (*pb.AddTorrentResponse, error) {
	cmd := fmt.Sprintf("transmission-remote %s -a %s", os.Getenv("TRANSMISSION_HOST"), req.MagnetLink)
	if t.IsDryRun {
		fmt.Println("DRYRUN: " + cmd)
	} else {
		fmt.Println("RUNNING: " + cmd)

		exe := execCommand("transmission-remote", os.Getenv("TRANSMISSION_HOST"),"-a", req.MagnetLink)

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

func parseTorrentStatusOutput(transmissionOutput string) []*pb.TorrentStatus {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from", r)
        }
    }()
    outputLines := strings.Split(transmissionOutput, "\n")
    if len(outputLines) > 2 {
        outputLines = outputLines[1:len(outputLines)-2]
    } else {
        return nil
    }
    fmt.Printf(outputLines[0])
    torrents := make([]*pb.TorrentStatus, len(outputLines))
    for i, line := range outputLines {
        match := TRANSMISSION_BODY_PARSER.FindStringSubmatch(line)
        result := make(map[string]string)
        for j, name := range TRANSMISSION_BODY_PARSER.SubexpNames() {
            if j != 0 && name != "" {
                result[name] = match[j]
            }
        }
        torrents[i] = &pb.TorrentStatus {
            Id : result["id"],
            Done : result["done"],
            Have : result["have"],
            Eta : result["eta"],
            Up : result["up"],
            Down : result["down"],
            Ratio : result["ratio"],
            Status : result["status"],
            Name : result["name"],
        }
    }
    fmt.Printf("Parsed %d lines\n", len(outputLines))
    return torrents
}

func (t *TorrentServer) GetTorrentStatus() []*pb.TorrentStatus {
	cmd := fmt.Sprintf("transmission-remote %s -l",  os.Getenv("TRANSMISSION_HOST"))
    var torrents []*pb.TorrentStatus
	if t.IsDryRun {
		fmt.Println("DRYRUN: " + cmd)
	} else {
		fmt.Println("RUNNING: " + cmd)

		output, err := execCommand("transmission-remote",os.Getenv("TRANSMISSION_HOST"),"-l").Output()
        fmt.Printf("CMD output: %s\n", output)
		if err != nil {
			fmt.Println(err.Error())
		}
		torrents = parseTorrentStatusOutput(string(output))
	}
    return torrents
}

func (t *TorrentServer) GetStatus(context.Context, *pb.GetStatusRequest) (*pb.GetStatusResponse, error) {
    torrents := t.GetTorrentStatus()
    return &pb.GetStatusResponse{
        Torrents: torrents,
    }, nil
}


func ScheduleTorrentListener(delay time.Duration) {
	go func() {
		for {
			output, err := execCommand("transmission-remote",os.Getenv("TRANSMISSION_HOST"),"-l").Output()
			if err != nil{
				fmt.Println(err.Error())
			}
			torrents := parseTorrentStatusOutput(string(output))
			for i := range torrents {
				if torrentIsFinished(*torrents[i]) {
					err := execCommand("rsync", "-r", "/var/lib/flavium/downloads/complete/"+torrents[i].Name, "/tmp").Run()
					if err != nil{
						fmt.Printf("Copy failed: %v\n", err)
					}else{
					// run filebot
					output, err := execCommand("filebot", "-rename", "/tmp/"+torrents[i].Name).Output()
					fmt.Printf("Filebot ouput '%s'", output)
					if err != nil{
						fmt.Printf("Filebot failed: %v\n", err)
					}else{
						//check if movie or series
						// move to plex
						err = execCommand("mv", "/tmp/"+torrents[i].Name, "/var/lib/plex/data").Run()
						if err != nil{
							fmt.Printf("Move failed: %v\n", err)
						} else {
							// remove torrent and delete data
							err := execCommand("transmission-remote", os.Getenv("TRANSMISSION_HOST"), "--torrent", torrents[i].Id, "--remove-and-delete").Run()
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
