package main

import (
	"github.com/IcookTacos/kvdbstore/pkg/server"
)

func main() {
  configPath := "config.yaml"
	server.StartServer(configPath)
}
