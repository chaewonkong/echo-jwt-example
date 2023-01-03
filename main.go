package main

import (
	"go-jwt/modules"
)

func main() {
	server, err := modules.InitializeServer()
	if err != nil {
		panic(err)
	}
	server.Start()
}
