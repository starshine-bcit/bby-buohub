package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	gin.DisableConsoleColor()
    r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
        c.String(http.StatusOK, "pong")
    })

	r.POST("/login", HandleLogin)

	r.POST("/create", HandleCreate)

	r.POST("/auth", HandleAuth)

	r.POST("/refresh", HandleRefresh)

	return r
}