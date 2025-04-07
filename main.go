package main

import(
  "encoding/json"
  "fmt"
  "log"
  "net/http"
  "os"
)

type Signal struct {
  UserID string `json:"user_id"`
  EventType string `json:"event_type"`
  Timestamp int64 `json:"timestamp"`
}

var signalFile = "signals.log"

func handleSignal(w http.ResponseWriter, r *http.Request){
  if r.Method != http.MethodPost{
    http.Error(w,"Only POST allowed", http.StatusMethodNotAllowed)
    return
  }
  var sig Signal
  if err := json.NewDecoder(r.Body).Decode(&sig); err!=nil{
    http.Error(w, "Invalid JSON", http.StatusBadRequest)
    return
  }
  defer file.Close()
  data, _:= json.Marshal(sig)
  file.Write(data)
  file.WriteString("\n")

  log.Printf("Received signal: %+v\n",sig)
  fmt.Fprintf(w,"Signal recorded!\n")
}

func main(){
  http.HandleFunc("/signal", handleSignal)
  log.Println("Server started at :8080")
  log.Fatal(http.ListenAndServe(":8080",nil))
}
