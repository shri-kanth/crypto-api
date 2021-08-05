package main

import (
	"crypto-api/server"
)

func main() {
	// config := config.GetConfig()

	server := &server.App{}
	server.Initialize()
	server.Run(":3000")
}