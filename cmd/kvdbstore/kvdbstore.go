package main

import (
  "log"
	"github.com/zeidlitz/kvdbstore/pkg/server"
)

func main() {
	configPath := "config.yaml"
  err, address := server.LoadConfig(configPath)
  if err != nil{
    log.Print("Unable to parse ", configPath, "\nGot :", err.Error())
    address = "localhost:8090"
  }
	server.StartServer(address)
}
