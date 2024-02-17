package util

import (
	"encoding/xml"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var StagingDir, ReadyDir string

func InitDirs() {
	ex, _ := os.Executable()
	uploadDir := filepath.Join(filepath.Dir(filepath.Dir(ex)), "uploads")
	StagingDir = filepath.Join(uploadDir, "staging")
	ReadyDir = filepath.Join(uploadDir, "ready")
	if err := os.MkdirAll(StagingDir, 0770); err != nil {
		ErrorLogger.Fatalf("Could not make staging dir: err %v\n", err.Error())
	}
	if err := os.MkdirAll(ReadyDir, 0770); err != nil {
		ErrorLogger.Fatalf("Could not make ready dir: err %v\n", err.Error())
	}
}

type ErrorResponse struct {
	ErrorName string `json:"errorName"`
	ErrorText string `json:"errorText"`
}

type GenericResponse struct {
	Message string `json:"message"`
}

type PIDConfigure struct {
	XMLName        xml.Name `xml:"PIDConfigure"`
	PID            uint     `xml:"PID,attr"`
	Name           string   `xml:"name,attr"`
	StreamType     string   `xml:"StreamType,attr"`
	ID             uint     `xml:"ID,attr"`
	Timescale      uint     `xml:"Timescale,attr"`
	ClockID        uint     `xml:"ClockID,attr"`
	Duration       string   `xml:"Duration,attr"`
	SAR            string   `xml:"SAR,attr"`
	Language       string   `xml:"Language,attr"`
	CodecID        string   `xml:"CodecID,attr"`
	PlaybackMode   string   `xml:"PlaybackMode,attr"`
	DecoderConfig  string   `xml:"DecoderConfig,attr"`
	Width          uint     `xml:"Width,attr"`
	Height         uint     `xml:"Height,attr"`
	FullRange      bool     `xml:"FullRange,attr"`
	ColorPrimaries string   `xml:"ColorPrimaries,attr"`
	ColorTransfer  string   `xml:"ColorTransfer,attr"`
	ColorMatrix    string   `xml:"ColorMatrix,attr"`
	ChromaLoc      string   `xml:"ChromaLoc,attr"`
	FPS            string   `xml:"FPS,attr"`
	URL            string   `xml:"URL,attr"`
	Cached         bool     `xml:"Cached,attr"`
	IsDefault      bool     `xml:"IsDefault,attr"`
	ChapTimes      string   `xml:"ChapTimes,attr"`
	ChapNames      string   `xml:"ChapNames,attr"`
	Delay          int      `xml:"Delay,attr"`
	Unframed       bool     `xml:"Unframed,attr"`
	StreamSubtype  string   `xml:"StreamSubtype,attr"`
}

type GPACInspect struct {
	XMLName       xml.Name       `xml:"GPACInspect"`
	PIDConfigures []PIDConfigure `xml:"PIDConfigure"`
}

func (p *PIDConfigure) GetDurationSeconds() uint {
	parts := strings.Split(p.Duration, "/")
	if len(parts) != 2 {
		return 0
	}
	p1, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0
	}
	p2, err := strconv.Atoi(parts[1])
	if err != nil || p2 == 0 {
		return 0
	}
	return uint(p1 / p2)
}
