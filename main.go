package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
)

type StartRequest struct {
	Action string `json:"action"`
}

type StartResponse struct {
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

func main() {
	// write the MAC of the device that you want to WOL to environment variables or write your mac in line 25 to replace "XX"
	mac := os.Getenv("MAC_ADDRESS")
	if mac == "" {
		mac = "XX:XX:XX:XX:XX:XX"
		log.Println("Warning: MAC_ADDRESS not set, using default", mac)
	}
	cmdName := os.Getenv("WAKE_CMD")
	if cmdName == "" {
		cmdName = "wakeonline"
	}
	port := os.Getenv("PORT")//default port is 11451
	if port == "" {
		port = "11451"
	}

	// API 
	http.HandleFunc("/start", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		var req StartRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
		log.Printf("Received start signal, action: %s", req.Action)

		cmd := exec.Command(cmdName, mac)
		output, err := cmd.CombinedOutput()
		if err != nil {
			errMsg := fmt.Sprintf("Failed: %v, output: %s", err, string(output))
			log.Println(errMsg)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(StartResponse{Error: errMsg})
			return
		}
		log.Printf("Wake command executed: %s", string(output))
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(StartResponse{Message: "Wake-on-LAN signal sent"})
	})

	// static file service
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	// start the server
	addr := "0.0.0.0:" + port
	log.Printf("Server starting on http://%s", addr)
	log.Printf("Using MAC: %s, command: %s", mac, cmdName)

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal("Server failed:", err)
	}
}