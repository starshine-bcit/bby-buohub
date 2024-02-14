package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() {
	gin.DisableConsoleColor()
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.POST("/upload", HandleUpload)

	r.GET("/stream", HandleStream)
}
