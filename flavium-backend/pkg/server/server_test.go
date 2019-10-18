package server

import (
	pb "../torrents"
	"reflect"
	"testing"
)

const (
	head = "ID     Done       Have  ETA           Up    Down  Ratio  Status       Name \n"
	tail = "Sum:           3.60 GB             180.0     0.0\n"
 	body =
	"   1   100%    1.59 GB  Done         0.0     0.0    0.6  Idle         My day at the zoo\n" +
	"   2   100%    2.01 GB  Done       180.0     0.0    0.0  Seeding      Wedding photos\n" +
	"   3    93%    2.11 GB  23 sec       0.0  6819.0    0.0  Downloading  Long name of video with S[p]e.ci#al Ch^a.rs}{ \n" +
	"   4    n/a       None  Unknown      0.0     0.0   None  Downloading  Long+name+with+pluses+for+some+reason\n" +
	"   5    88%    1.48 GB  31 sec       0.0  6008.0    0.0  Up & Down    Up and down torrent\n"
)

var (
	expectedResults = []*pb.TorrentStatus {
		{
			Id:     "1",
			Done:   "100%",
			Have:   "1.59 GB",
			Eta:    "Done",
			Up:     "0.0",
			Down:   "0.0",
			Result: "0.6",
			Status: "Idle",
			Name:   "My day at the zoo",
		}, {
			Id:     "2",
			Done:   "100%",
			Have:   "2.01 G",
			Eta:    "Done",
			Up:     "180.0",
			Down:   "0.0",
			Result: "Seeding",
			Status: "Seeding",
			Name:   "Wedding photos",
		}, {
			Id:     "3",
			Done:   "93%",
			Have:   "2.11 GB",
			Eta:    "23 sec",
			Up:     "0.0",
			Down:   "6819.0",
			Result: "0.0",
			Status: "Downloading",
			Name:   "Long name of video with S[p]e.ci#al Ch^a.rs}{ ",
		}, {
			Id:     "4",
			Done:   "n/a",
			Have:   "None",
			Eta:    "Unknown",
			Up:     "0.0",
			Down:   "0.0",
			Result: "None",
			Status: "Downloading",
			Name:   "Long+name+with+pluses+for+some+reason",
		}, {
			Id:     "5",
			Done:   "88%",
			Have:   "1.48 GB",
			Eta:    "31 sec",
			Up:     "0.0",
			Down:   "6008.0",
			Result: "0.0",
			Status: "Up & Down",
			Name:   "Up and down torrent",
		},
	}
)

func TestRegex(t *testing.T) {
	testOutput := head + body + tail
	result := getTorrentStatus(testOutput)
	for i := range expectedResults {
		if reflect.DeepEqual(result, expectedResults) {
			t.Errorf("Test case #%d failed", i)
		}
	}
}
