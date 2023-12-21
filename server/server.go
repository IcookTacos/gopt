package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`
}

func loadConfig() (string, string){
  conf, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("Error reading configuration file: %v", err)
	}

	var config Config
	err = yaml.Unmarshal(conf, &config)
	if err != nil {
		log.Fatalf("Error unmarshalling YAML: %v", err)
	}

	host := config.Server.Host
	port := config.Server.Port
  return host, port
}

func logRequest(req *http.Request){
  uagent := req.Header.Get("User-Agent")
  log.Printf("Request from: %s \n", uagent)
}

func status(w http.ResponseWriter, req *http.Request) {
  logRequest(req)
  response := map[string]string{"server": "running", "status" : "200 OK"}
  jsonResponse, err := json.Marshal(response)
  if(err != nil){
    http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
  }
  w.WriteHeader(http.StatusOK)
  w.Write(jsonResponse)
}

func list(w http.ResponseWriter, req *http.Request){
  if(req.Method != http.MethodGet){
    response := fmt.Sprintf("Incorrect method\nGot     : %s\nRequire : %s", req.Method, http.MethodGet)
    http.Error(w, response, http.StatusBadRequest)
    return
  }
  response := map[string]string{"data": " ", "status" : "200 OK"}
  jsonResponse, _ := json.Marshal(response)
  w.WriteHeader(http.StatusOK)
  w.Write(jsonResponse)
}

func store(w http.ResponseWriter, req *http.Request){
  //TODO: implement
}

func StartServer(){
  host, port := loadConfig()
  address := fmt.Sprintf("%s:%s", host, port)
  fmt.Printf("Serving on %s\n", address) 
  http.HandleFunc("/api/status", status)
  http.HandleFunc("/api/list", list)
  http.ListenAndServe(address,nil)
}
