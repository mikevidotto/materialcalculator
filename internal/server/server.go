package server

import (
    "fmt"
    "net/http"
    "log"
    "materialcalculator/internal/materialcalculator"
    "encoding/json"
)

func StartServer() {
    http.HandleFunc("/calculate", CalculateMaterials)
    err := http.ListenAndServe("localhost:8085", nil)
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

    lumber := materialcalculator.CalculateMaterials(shed)

    fmt.Printf("RECEIVED DATA: \n\n\n%#v\n", lumber)

    fmt.Println("ok??")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("materials data received."))
}

