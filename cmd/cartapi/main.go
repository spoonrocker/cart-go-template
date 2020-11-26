package main

import (
	base "cartapi"
	"cartapi/cartapi"
	"cartapi/cartapi/logger"
)

func main() {
	logger.SetUp()
	cartapi.LoadConfig()
	server := base.InitServer()
	server.Start()
}
