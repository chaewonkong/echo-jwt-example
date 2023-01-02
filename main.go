package main

import (
	"go-jwt/modules"
)

func main() {
	server := modules.InitializeServer()
	server.Start()
}
