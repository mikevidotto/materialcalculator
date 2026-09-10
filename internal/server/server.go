package server

import (
    "fmt"
    "net/http"
    "log"
    "materialcalculator/internal/orders"
    "encoding/json"
    "path/filepath"
    "time"
)

func StartServer() {

    router := http.NewServeMux()
    server := http.Server {
        Addr: "localhost:8085",
        Handler: router,
        ReadTimeout: 5 * time.Second,
        WriteTimeout: 10 * time.Second,
        IdleTimeout: 60 * time.Second,
    }

    frontendPath, _ := filepath.Abs("./frontend/")

    fmt.Println("looking for frontend path, ", frontendPath)
    fs := http.FileServer(http.Dir(frontendPath))

    router.Handle("/", fs)
    router.HandleFunc("/api/calculate", NewShedHandler)

    log.Println("starting server on localhost:8085...")
    err := server.ListenAndServe()
    if err != nil {
        log.Fatal("Error starting server: ", err)
    }
}

/*
this handler should 
1. validate the json request and decode json data into OrderData{}
2. invoke the business layer function NewOrder(orderData)
3. generate a response.
*/
func NewShedHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    newOrder := orders.OrderData{}

    err := json.NewDecoder(r.Body).Decode(&newOrder)
    if err != nil {
        http.Error(w, "Error decoding json body: "+err.Error(), http.StatusBadRequest)
        return
    }

    result, err := orders.NewOrder(newOrder)
    if err != nil {
        http.Error(w, "Error calculating materials: "+err.Error(), http.StatusBadRequest)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(result); err != nil {
        log.Println("Error encoding response:", err)
    }
}
