package service

import (
	"bytes"
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
	resultDir := filepath.Join(util.ReadyDir, idstr)
	inspect := filepath.Join(util.StagingDir, "inspect")
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
		util.WarningLogger.Printf("Error executing gpac command. err: %v stderr: %v\n", statsErr.String())
	}
	numProccessing -= 1
}
