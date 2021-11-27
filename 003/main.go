package main

import (
	"003/config"
	"003/httpserver"
	"003/logger"
	"os"
)

func main()  {
	logger.LogInit(os.Getenv("LOGINFO"))
	err := config.LoadConfig()
	if err != nil{
		logger.Error("server config get failed")
		panic(err)
	}
	httpserver.HttpServer()
}