package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/starshine-bcit/bby-buohub/cdn/router"
	"github.com/starshine-bcit/bby-buohub/cdn/service"
	"github.com/starshine-bcit/bby-buohub/cdn/util"
)

func main() {

	http.HandleFunc("/player/", func(w http.ResponseWriter, r *http.Request) {
		// Define allowed origins
		allowedOrigins := "http://localhost:8999, http://127.0.0.1:8999"

		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", allowedOrigins)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Serve your routes
		http.DefaultServeMux.ServeHTTP(w, r)
	})

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
