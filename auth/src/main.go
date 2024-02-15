package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/starshine-bcit/bby-buohub/auth/router"
	"github.com/starshine-bcit/bby-buohub/auth/service"
	"github.com/starshine-bcit/bby-buohub/auth/util"
)

func main() {
	cfg := util.Load_config()
	util.InfoLogger.Println("Successfully parsed config and env")
	util.LoadKey()
	util.InfoLogger.Println("Successfully parsed key files")
	db, err := service.ConnectToMariaDB(cfg)
	if err != nil {
		util.ErrorLogger.Fatalf("Could not connect to database %v\n", err)
	}
	service.Migrate(db)
	router.Db = db
	util.InfoLogger.Println("Successfully connected to DB and initialized tables")
	r := router.SetupRouter()
	env := os.Getenv("SERVER_ENV")
	if env == "dev" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	r.Run(fmt.Sprintf("%v:%v", cfg.Server.Host, cfg.Server.Port))
}
