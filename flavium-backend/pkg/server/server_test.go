package server

import (
	pb "flavium-backend/pkg/torrents"
    "fmt"
	"reflect"
	"testing"
	"os"
	"os/exec"
)

const (
	transmissionListHead = "ID     Done       Have  ETA           Up    Down  Ratio  Status       Name \n"
	transmissionListTail = "Sum:           3.60 GB             180.0     0.0\n"
 	transmissionListBody =
	"   1   100%    1.59 GB  Done         0.0     0.0    0.6  Idle         My day at the zoo\n" +
	"   2   100%    2.01 GB  Done       180.0     0.0    0.0  Seeding      Wedding photos\n" +
	"   3    93%    2.11 GB  23 sec       0.0  6819.0    0.0  Downloading  Long name of video with S[p]e.ci#al Ch^a.rs}{ \n" +
	"   4    n/a       None  Unknown      0.0     0.0   None  Downloading  Long+name+with+pluses+for+some+reason\n" +
	"   5    88%    1.48 GB  31 sec       0.0  6008.0    0.0  Up & Down    Up and down torrent\n"
)

var (
	expectedTorrentResults = []*pb.TorrentStatus {
		{
			Id:     "1",
			Done:   "100%",
			Have:   "1.59 GB",
			Eta:    "Done",
			Up:     "0.0",
			Down:   "0.0",
			Ratio: "0.6",
			Status: "Idle",
			Name:   "My day at the zoo",
		}, {
			Id:     "2",
			Done:   "100%",
			Have:   "2.01 GB",
			Eta:    "Done",
			Up:     "180.0",
			Down:   "0.0",
			Ratio:  "0.0",
			Status: "Seeding",
			Name:   "Wedding photos",
		}, {
			Id:     "3",
			Done:   "93%",
			Have:   "2.11 GB",
			Eta:    "23 sec",
			Up:     "0.0",
			Down:   "6819.0",
			Ratio: "0.0",
			Status: "Downloading",
			Name:   "Long name of video with S[p]e.ci#al Ch^a.rs}{ ",
		}, {
			Id:     "4",
			Done:   "n/a",
			Have:   "None",
			Eta:    "Unknown",
			Up:     "0.0",
			Down:   "0.0",
			Ratio: "None",
			Status: "Downloading",
			Name:   "Long+name+with+pluses+for+some+reason",
		}, {
			Id:     "5",
			Done:   "88%",
			Have:   "1.48 GB",
			Eta:    "31 sec",
			Up:     "0.0",
			Down:   "6008.0",
			Ratio: "0.0",
			Status: "Up & Down",
			Name:   "Up and down torrent",
		},
	}
)

func assertTorrentStatusEqual(result *pb.TorrentStatus, expected *pb.TorrentStatus, t *testing.T) {
    if result.Id != expected.Id {
        t.Errorf("Test case failed")
        fmt.Println("Got Id:")
        fmt.Println(result.Id)
        fmt.Println("Expected Id:")
        fmt.Println(expected.Id)
    }
    if result.Done != expected.Done {
        t.Errorf("Test case failed")
        fmt.Println("Got Done:")
        fmt.Println(result.Done)
        fmt.Println("Expected Done:")
        fmt.Println(expected.Done)
    }
    if result.Have != expected.Have {
        t.Errorf("Test case failed")
        fmt.Println("Got Have:")
        fmt.Println(result.Have)
        fmt.Println("Expected Have:")
        fmt.Println(expected.Have)
    }
    if result.Eta != expected.Eta {
        t.Errorf("Test case failed")
        fmt.Println("Got Eta:")
        fmt.Println(result.Eta)
        fmt.Println("Expected Eta:")
        fmt.Println(expected.Eta)
    }
    if result.Up != expected.Up {
        t.Errorf("Test case failed")
        fmt.Println("Got Up:")
        fmt.Println(result.Up)
        fmt.Println("Expected Up:")
        fmt.Println(expected.Up)
    }
    if result.Down != expected.Down {
        t.Errorf("Test case failed")
        fmt.Println("Got Down:")
        fmt.Println(result.Down)
        fmt.Println("Expected Down:")
        fmt.Println(expected.Down)
    }
    if result.Ratio != expected.Ratio {
        t.Errorf("Test case failed")
        fmt.Println("Got Ratio:")
        fmt.Println(result.Ratio)
        fmt.Println("Expected Ratio:")
        fmt.Println(expected.Ratio)
    }
    if result.Status != expected.Status {
        t.Errorf("Test case failed")
        fmt.Println("Got Status:")
        fmt.Println(result.Status)
        fmt.Println("Expected Status:")
        fmt.Println(expected.Status)
    }
    if result.Name != expected.Name {
        t.Errorf("Test case failed")
        fmt.Println("Got Name:")
        fmt.Println(result.Name)
        fmt.Println("Expected Name:")
        fmt.Println(expected.Name)
    }
}

func TestRegex(t *testing.T) {
	testOutput := transmissionListHead + transmissionListBody + transmissionListTail
	result := parseTorrentStatusOutput(testOutput)
	for i := range expectedTorrentResults {
        fmt.Printf("Test #%d:\n", i)
        assertTorrentStatusEqual(result[i], expectedTorrentResults[i], t)
	}
}

func TestGetTorrentStatus(t *testing.T) {
    fakeExec := FakeExecCommand{
        testTarget: "TestHelperGetTorrentStatus",
    }
    execCommand = fakeExec.FakeExecCommand
    defer func() { execCommand = exec.Command }()
    os.Setenv("TRANSMISSION_HOST", "localhost")
    torrentServer := TorrentServer{IsDryRun: false}
    result := torrentServer.GetTorrentStatus()
    if len(result) != len(expectedTorrentResults) {
        t.Errorf(
            "Test case failed: Too few result Got - %d Expected - %d",
            len(result),
            len(expectedTorrentResults),
        )
        return
    }
	for i := range expectedTorrentResults {
        assertTorrentStatusEqual(result[i], expectedTorrentResults[i], t)
	}
}

func TestHelperGetTorrentStatus(t *testing.T) {
    if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
        return
    }
    expectedArgs := []string {
        "transmission",
        "localhost",
        "-l",
    }
    command := os.Args[3:]
    if reflect.DeepEqual(command, expectedArgs) {
        // TODO test transmission unexpected arguments
        // TODO test transmission cannot connect to host
        os.Exit(1)
    }
	testOutput := transmissionListHead + transmissionListBody + transmissionListTail
    fmt.Fprint(os.Stdout, testOutput)
    os.Exit(0)
}


type FakeExecCommand struct {
	testTarget string
}

func (t *FakeExecCommand) FakeExecCommand(command string, args...string) *exec.Cmd {
    cs := []string{"-test.run=" + t.testTarget}
    cs = append(cs, args...)
    cmd := exec.Command(os.Args[0], cs...)
    cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1"}
    return cmd
}
