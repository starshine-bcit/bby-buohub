package router

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/starshine-bcit/bby-buohub/cdn/service"
	"github.com/starshine-bcit/bby-buohub/cdn/util"
	"gorm.io/gorm"
)

var Db *gorm.DB

func HandleUpload(c *gin.Context) {
	util.InfoLogger.Println("Handling incoming file upload")
	file, err := c.FormFile("file")
	if err != nil {
		util.WarningLogger.Printf("Could not get file from multipart-formdata. err: %v\n", err.Error())
		// return error http
		return
	}
	uuidstr, exist := c.GetPostForm("uuid")
	if !exist || uuidstr == "" {
		util.WarningLogger.Println("Could not parse uuid field from multipart=formdata")
	}
	id, err := uuid.FromBytes([]byte(uuidstr))
	if err != nil {
		util.ErrorLogger.Printf("Could not parse uuid from uuid string. err: %v\n", err.Error())
		// return error http
		return
	}
	video, err := service.GetByUUID(Db, id)
	if err != nil {
		// return error http
		return
	}
	fname := filepath.Base(file.Filename)
	tempDir := filepath.Join(util.StagingDir, id.String())
	os.Mkdir(tempDir, 0770)
	c.SaveUploadedFile(file, filepath.Join(tempDir, fname))
	video.OriginalFilename.String = fname
	video.OriginalFilename.Valid = true
	util.InfoLogger.Println("Incoming file upload successfull, ready for processing")
	go service.ProcessIncoming(video)
	msg := &util.GenericResponse{Message: "Upload successfull, queued for processing"}
	c.JSON(http.StatusCreated, msg)
}

func HandleStream(c *gin.Context) {
	//
}
