package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Signal struct {
	UserID    string `json:"user_id"`
	EventType string `json:"event_type"`
	Timestamp int64  `json:"timestamp"`
}

var signalFile = "signals.log"

func handleSignal(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	var sig Signal
	err := json.NewDecoder(r.Body).Decode(&sig)
	if err != nil {
		log.Printf("Error decoding JSON: %v", err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	file, err := os.OpenFile(signalFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("Error opening file: %v", err)
		http.Error(w, "Could not write signal", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	data, _ := json.Marshal(sig)
	file.Write(data)
	file.WriteString("\n")

	log.Printf("✅ Received signal: %+v\n", sig)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Signal recorded!")
}

func main() {
	http.HandleFunc("/signal", handleSignal)
	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
