package service

import (
	"bytes"
	"encoding/xml"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/starshine-bcit/bby-buohub/cdn/util"
)

var numProccessing uint
var gpacExec string

const maxProcessing uint = 5

func init() {
	ex, _ := os.Executable()
	gpacExec = filepath.Join(filepath.Dir(ex))
}

func ProcessIncoming(video *Video) {
	idstr := video.UUID.String()
	workingDir := filepath.Join(util.StagingDir, idstr)
	// resultDir := filepath.Join(util.ReadyDir, idstr)
	// inspect := filepath.Join(util.StagingDir, "inspect")
	for numProccessing >= maxProcessing {
		util.InfoLogger.Println("Import process sleeping, max workers reached")
		time.Sleep(time.Second * 10)
	}
	numProccessing += 1
	statsStd := new(bytes.Buffer)
	statsErr := new(bytes.Buffer)
	statsCmd := exec.Command(gpacExec, "arg1", "arg2")
	statsCmd.Dir = workingDir
	statsCmd.Stdout = statsStd
	statsCmd.Stderr = statsErr
	if err := statsCmd.Run(); err != nil {
		util.WarningLogger.Printf("Error executing gpac command. err: %v\n", statsErr.String())
	}
	stats := &util.GPACInspect{}
	if err := xml.Unmarshal(statsStd.Bytes(), stats); err != nil {
		util.ErrorLogger.Fatalf("Cannot unmarshal xml. err %v\n", err.Error())
	}
	util.InfoLogger.Println("Successfully retrieved stats, now parsing.")
	for _, el := range stats.PIDConfigures {
		switch el.StreamType {
		case "Audio":
			if !video.AudioCodec.Valid {
				video.AudioCodec.String = el.CodecID
			}
		case "Video":
			if !video.VideoCodec.Valid {
				video.VideoCodec.String = el.CodecID
			}
			if !video.Height.Valid {
				video.Width.Int32 = int32(el.Height)
			}
			if !video.Width.Valid {
				video.Height.Int32 = int32(el.Width)
			}
			if !video.Height.Valid {
				video.Width.Int32 = int32(el.GetDurationSeconds())
			}
		}
	}
	numProccessing -= 1
}
