package main

import (
	"github.com/zeidlitz/kvdbstore/pkg/server"
)

func main() {
	configPath := "config.yaml"
	server.StartServer(configPath)
}
