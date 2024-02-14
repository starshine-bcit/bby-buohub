package util

import (
	"os"
	"path/filepath"
)

var StagingDir, ReadyDir string

func init() {
	ex, _ := os.Executable()
	uploadDir := filepath.Join(filepath.Dir(filepath.Dir(ex)), "upload")
	StagingDir = filepath.Join(uploadDir, "staging")
	ReadyDir = filepath.Join(uploadDir, "ready")
	_ = os.MkdirAll(StagingDir, 0770)
	_ = os.MkdirAll(ReadyDir, 0770)
}

type ErrorResponse struct {
	ErrorName string `json:"errorName"`
	ErrorText string `json:"errorText"`
}

type GenericResponse struct {
	Message string `json:"message"`
}

type PidVideo struct {

}

type PidAudio struct {

}

type PidText struct {

}

type GpacInspect struct {
	
}