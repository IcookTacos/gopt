package server

import (
	"encoding/json"
	"fmt"
	"github.com/zeidlitz/kvdbstore/pkg/storage"
	"gopkg.in/yaml.v3"
	"io"
	"log"
	"net/http"
	"os"
)

type Config struct {
	Server struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`
}

type Payload struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func LoadConfig(configPath string) (error, string) {
	conf, err := os.ReadFile(configPath)
	if err != nil {
		return err, ""
	}

	var config Config
	err = yaml.Unmarshal(conf, &config)
	if err != nil {
		return err, ""
	}

	host := config.Server.Host
	port := config.Server.Port
	address := fmt.Sprintf("%s:%s", host, port)
	return nil, address
}

func logRequest(req *http.Request) {
	uagent := req.Header.Get("User-Agent")
	log.Printf("Request from: %s \n", uagent)
}

func apiHandler(w http.ResponseWriter, req *http.Request) {
	method := req.Method

	if method == http.MethodGet {
		get(w, req)
	}

	if method == http.MethodPost {
		post(w, req)
	}
}

func status(w http.ResponseWriter, req *http.Request) {
	logRequest(req)
	response := map[string]string{"server": "running", "status": "200 OK"}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func get(w http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
	}

	var data Payload
	err = json.Unmarshal(body, &data)
	if err != nil {
		http.Error(w, "Error unmarshalling request payload", http.StatusBadRequest)
	}

	err, result := storage.List(data.Key)
	if err != nil {
		http.Error(w, "Bad request: key not found", http.StatusBadRequest)
		return
	}

	response := map[string]interface{}{"data": result, "status": "200 OK"}
	jsonResponse, _ := json.Marshal(response)
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func post(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		response := fmt.Sprintf("Incorrect method\nGot     : %s\nRequire : %s", req.Method, http.MethodPost)
		http.Error(w, response, http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
	}

	var data Payload
	err = json.Unmarshal(body, &data)
	if err != nil {
		http.Error(w, "Error decoding JSON data", http.StatusBadRequest)
		return
	}

	err = storage.Store(data.Key, data.Value)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{data.Key: data.Value, "status": "200 OK"}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func StartServer(address string) {
	// host, port := loadConfig(configPath)
	http.HandleFunc("/api/status", status)
	http.HandleFunc("/api", apiHandler)
	fmt.Printf("Serving on %s\n", address)
	http.ListenAndServe(address, nil)
}
