package server

import (
    "fmt"
    "net/http"
    "log"
    "materialcalculator/internal/materialcalculator"
    "encoding/json"
    "path/filepath"
)

func StartServer() {
    filepath, _ := filepath.Abs("./frontend/")
    fmt.Println("looking for frontend path, ", filepath) 
    fs := http.FileServer(http.Dir(filepath))


    http.Handle("/", http.StripPrefix("/", fs))
    http.HandleFunc("/api/calculate", CalculateMaterials)

    log.Println("starting server on localhost:8085...")
    err := http.ListenAndServe(":8085", nil)
    if err != nil {
        log.Fatal("error starting server: ", err)
    }
}


func CalculateMaterials(w http.ResponseWriter, r *http.Request) {
    fmt.Println("fdjksa")
    if r.Method != http.MethodPost{
        http.Error(w, "Method now allowed", http.StatusMethodNotAllowed)
    }
    shed := materialcalculator.ShedData{}

    err := json.NewDecoder(r.Body).Decode(&shed)
    if err != nil {
        http.Error(w, "Error decoding json body: "+err.Error(), http.StatusBadRequest)
        return
    }


    fmt.Printf("\n\nShed Received:\n%#v\n", shed, "\n")
    shed.Length = shed.Length * 12
    shed.Width = shed.Width * 12
    shed.Height = shed.Height * 12

    lumber := materialcalculator.CalculateMaterials(shed)

    fmt.Printf("%#v\n", lumber)

    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Shed Materials Data received."))
}

