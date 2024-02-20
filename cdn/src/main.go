package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/starshine-bcit/bby-buohub/cdn/router"
	"github.com/starshine-bcit/bby-buohub/cdn/service"
	"github.com/starshine-bcit/bby-buohub/cdn/util"
)

func main() {
	cfg := util.Load_config()
	util.InfoLogger.Println("Successfully parsed config and env")
	db, err := service.ConnectToMariaDB(cfg)
	util.Cfg = cfg
	if err != nil {
		util.ErrorLogger.Fatalf("Could not connect to database %v\n", err)
	}
	service.Migrate(db)
	router.Db = db
	util.InfoLogger.Println("Successfully made database connection and migrated")
	util.InitDirs()
	util.InfoLogger.Println("Successfully initialized storage dirs")
	r := router.SetupRouter()
	env := os.Getenv("SERVER_ENV")
	if env == "dev" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	r.Run(fmt.Sprintf("%v:%v", cfg.Server.Host, cfg.Server.Port))
}
