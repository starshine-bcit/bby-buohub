package service

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/starshine-bcit/bby-buohub/cdn/util"
	"gorm.io/gorm"
)

var numProccessing uint

const maxProcessing uint = 5
const gpacExec = "gpac"

func ProcessIncoming(video *Video, db *gorm.DB) {
	idstr := video.UUID.String()
	workingDir := filepath.Join(util.StagingDir, idstr)
	videoPath := filepath.Join(workingDir, video.OriginalFilename.String)
	resultDir := filepath.Join(util.ReadyDir, idstr)
	ssPath := filepath.Join(resultDir, "thumb.png")
	mpdName := fmt.Sprintf("%v.mpd", strings.TrimSuffix(
		video.OriginalFilename.String,
		filepath.Ext(video.OriginalFilename.String)),
	)
	resultMpdPath := filepath.Join(resultDir, mpdName)
	if err := os.Mkdir(resultDir, 0770); err != nil {
		util.ErrorLogger.Printf("Could not make results unique directory while processing, err: %v\n", err.Error())
	}
	for numProccessing >= maxProcessing {
		util.InfoLogger.Println("Import process sleeping, max workers reached")
		time.Sleep(time.Second * 10)
	}
	numProccessing += 1
	statsStd := new(bytes.Buffer)
	statsErr := new(bytes.Buffer)
	statsCmd := exec.Command(gpacExec, "-i", videoPath, "inspect:xml")
	statsCmd.Dir = workingDir
	statsCmd.Stdout = statsStd
	statsCmd.Stderr = statsErr
	if err := statsCmd.Run(); err != nil {
		util.WarningLogger.Printf("Error executing gpac command. err: %v\n", statsErr.String())
		numProccessing -= 1
		return
	}
	stats := &util.GPACInspect{}
	if err := xml.Unmarshal(statsStd.Bytes(), stats); err != nil {
		util.ErrorLogger.Printf("Cannot unmarshal xml. err %v\n", err.Error())
		numProccessing -= 1
		return
	}
	util.InfoLogger.Println("Successfully retrieved stats, now parsing.")
	for _, el := range stats.PIDConfigures {
		util.InfoLogger.Println(el)
		switch el.StreamType {
		case "Audio":
			if !video.AudioCodec.Valid {
				video.AudioCodec.String = el.CodecID
				video.AudioCodec.Valid = true
			}
		case "Visual":
			if !video.VideoCodec.Valid {
				video.VideoCodec.String = el.CodecID
				video.VideoCodec.Valid = true
			}
			if !video.Height.Valid {
				video.Width.Int32 = int32(el.Height)
				video.Width.Valid = true
			}
			if !video.Width.Valid {
				video.Height.Int32 = int32(el.Width)
				video.Height.Valid = true
			}
			if !video.RunTime.Valid {
				video.RunTime.Int32 = int32(el.GetDurationSeconds())
				video.RunTime.Valid = true
			}
		}
	}
	procStd := new(bytes.Buffer)
	procErr := new(bytes.Buffer)
	procCmd := exec.Command(gpacExec, "-i", videoPath, "-o", resultMpdPath)
	procCmd.Dir = resultDir
	procCmd.Stdout = procStd
	procCmd.Stderr = procErr
	if err := procCmd.Run(); err != nil {
		util.WarningLogger.Printf("Error executing gpac command. err: %v\n", statsErr.String())
		numProccessing -= 1
		return
	}
	ssCMD := exec.Command(
		"ffmpeg", "-y", "-loglevel", "fatal",
		"-hide_banner", "-ss", fmt.Sprintf("%v", video.RunTime.Int32/4),
		"-i", videoPath, "-frames:v", "1", "-q:v", "2", "-update", "1",
		ssPath)
	if err := ssCMD.Run(); err != nil {
		util.WarningLogger.Printf("Did not successfully take screenshot. err %v\n", err.Error())
	}
	if err := os.Remove(videoPath); err != nil {
		util.WarningLogger.Printf("Could not remove the input video file. err %v\n", err.Error())
	}
	if err := os.Remove(workingDir); err != nil {
		util.WarningLogger.Printf("Could not remove the input video folder. err %v\n", err.Error())
	}
	video.PosterFilename.String = "thumb.png"
	video.PosterFilename.Valid = true
	video.ProcessComplete = true
	video.ProcessedAt.Time = time.Now()
	video.ProcessedAt.Valid = true
	video.ManifestName.String = mpdName
	video.ManifestName.Valid = true
	UpdateVideo(db, video)
	numProccessing -= 1
}
