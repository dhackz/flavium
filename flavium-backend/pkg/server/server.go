package server

import (
	pb "../torrents"
	"context"
	"fmt"
	"os"
	"os/exec"
    "regexp"
    "strings"
)

const TRANSMISSION_BODY_EXPRESSION = "^\\s+" +
"(?P<id>\\d+\\*?)" +
"\\s+" +
"(?P<done>(?:\\d+%|n/a))" +
"\\s+" +
"(?P<have>\\d+.\\d+ [MGKmkg]B)" +
"\\s+" +
"(?P<eta>(?:\\d+ (?:min|hours))|Done|Unknown)" +
"\\s+" +
"(?P<up>\\d+.\\d+)" +
"\\s+" +
"(?P<down>\\d+.\\d+)" +
"\\s+" +
"(?P<ratio>\\d+.\\d+)" +
"\\s+" +
"(?P<status>\\w+|(\\w+ & \\w+))" +
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
    fmt.Println(outputLines)
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
