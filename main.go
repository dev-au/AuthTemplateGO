package main

import (
	"AuthTemplate/src"
	"AuthTemplate/src/api"
	"AuthTemplate/src/resources"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
)

func main() {

	src.Config.SetupEnv()
	resources.SetupDB()
	_ = resources.DB.AutoMigrate(src.Models...)
	resources.SetupRedis()
	if len(os.Args) > 1 {
		resources.InitCommands(os.Args[1])
		return
	}

	gin.SetMode(src.Config.ApplicationMode)
	ginEngine := api.GinEngine()

	s := &http.Server{
		Addr:           ":" + src.Config.Port,
		Handler:        ginEngine,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	_ = s.ListenAndServe()

}
