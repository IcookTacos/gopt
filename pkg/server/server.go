package server

import (
	"encoding/json"
	"fmt"
	"github.com/zeidlitz/kvdbstore/pkg/storage"
	"github.com/gorilla/mux"
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

type Storage struct {
	Value string `json:"value"`
}

func loadConfig(configPath string) (string, string) {
	conf, err := os.ReadFile(configPath)
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

func logRequest(req *http.Request) {
	uagent := req.Header.Get("User-Agent")
	log.Printf("Request from: %s \n", uagent)
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
	if req.Method != http.MethodGet {
		response := fmt.Sprintf("Incorrect method\nGot     : %s\nRequire : %s", req.Method, http.MethodGet)
		http.Error(w, response, http.StatusBadRequest)
		return
	}

	vars := mux.Vars(req)
	key := vars["key"]
	err, result := storage.List(key)

	if err != nil {
		http.Error(w, "Bad request, key not found", http.StatusBadRequest)
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

	var data Storage

	err = json.Unmarshal(body, &data)

	if err != nil {
		http.Error(w, "Error decoding JSON data", http.StatusBadRequest)
		return
	}

	vars := mux.Vars(req)
	key := vars["key"]

	err = storage.Store(key, data.Value)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{key: data.Value, "status": "200 OK"}
	jsonResponse, _ := json.Marshal(response)
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func StartServer(configPath string) {
	host, port := loadConfig(configPath)
	address := fmt.Sprintf("%s:%s", host, port)
	router := mux.NewRouter()
	router.HandleFunc("/api/status", status).Methods("GET")
	router.HandleFunc("/api/{key}", get).Methods("GET")
	router.HandleFunc("/api/{key}", post).Methods("POST")
	http.Handle("/", router)
	fmt.Printf("Serving on %s\n", address)
	http.ListenAndServe(address, nil)
}
