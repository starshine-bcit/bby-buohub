package router

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/starshine-bcit/bby-buohub/cdn/util"
)

func SetupRouter() *gin.Engine {
	gin.DisableConsoleColor()
	r := gin.Default()

	r.Use(cors.Default())

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.POST("/upload", HandleUpload)

	r.Static("/stream", util.ReadyDir)

	return r
}
