package main

import (
	"github.com/starshine-bcit/bby-buohub/cdn/router"
	"github.com/starshine-bcit/bby-buohub/cdn/service"
	"github.com/starshine-bcit/bby-buohub/cdn/util"
)

func main() {
	cfg := util.Load_config()
	util.InfoLogger.Println("Successfully parsed config and env")
	db, err := service.ConnectToMariaDB(cfg)
	if err != nil {
		util.ErrorLogger.Fatalf("Could not connect to database %v\n", err)
	}
	service.Migrate(db)
	util.InfoLogger.Println("Successfully made database connection and migrated")
	router.Db = db
}
