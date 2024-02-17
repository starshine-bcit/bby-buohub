package router

import (
	"errors"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/ryanfowler/uuid"
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
		c.JSON(http.StatusInternalServerError, &util.ErrorResponse{
			ErrorName: "ParsingFormError",
			ErrorText: "Could not get file from multipart/form-data",
		})
		return
	}
	uuidstr, exist := c.GetPostForm("uuid")
	if !exist || uuidstr == "" {
		util.WarningLogger.Println("Could not parse uuid field from multipart=formdata")
		c.JSON(http.StatusInternalServerError, &util.ErrorResponse{
			ErrorName: "ParsingFormError",
			ErrorText: "Could not get uuid from multipart/form-data",
		})
		return
	}
	uuid, err := uuid.ParseString(uuidstr)
	if err != nil {
		util.WarningLogger.Println("Could not parse uuid from string")
		c.JSON(http.StatusInternalServerError, &util.ErrorResponse{
			ErrorName: "ParsingFormError",
			ErrorText: "Could not convert incoming uuid to uuid type",
		})
		return
	}
	video, err := service.GetByUUID(Db, uuid)
	if err != nil {
		util.WarningLogger.Println("Could not database entry for id")
		c.JSON(http.StatusInternalServerError, &util.ErrorResponse{
			ErrorName: "DBLookupError",
			ErrorText: "Could not get the corresponding database entry for current id",
		})
		return
	}
	fname := filepath.Base(file.Filename)
	tempDir := filepath.Join(util.StagingDir, video.UUID.String())
	os.Mkdir(tempDir, 0770)
	c.SaveUploadedFile(file, filepath.Join(tempDir, fname))
	if ok := checkFileExists(filepath.Join(tempDir, fname)); !ok {
		util.WarningLogger.Println("Could not find saved file")
		c.JSON(http.StatusInternalServerError, &util.ErrorResponse{
			ErrorName: "FileNotExistError",
			ErrorText: "Could not save/find the corresponding uploaded file to disk",
		})
		return
	}
	video.OriginalFilename.String = fname
	video.OriginalFilename.Valid = true
	util.InfoLogger.Println("Incoming file upload successful, ready for processing")
	go service.ProcessIncoming(video, Db)
	msg := &util.GenericResponse{Message: "Upload successful, queued for processing"}
	c.JSON(http.StatusCreated, msg)
}

func checkFileExists(filePath string) bool {
	_, error := os.Stat(filePath)
	return !errors.Is(error, os.ErrNotExist)
}
