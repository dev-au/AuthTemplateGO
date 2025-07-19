package main

import (
	"AuthTemplate/src"
	"AuthTemplate/src/api"
	"AuthTemplate/src/resources"
	"os"
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

	webEngine := api.WebEngine()
	_ = webEngine.Run(":8000")
}
