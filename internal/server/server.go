package server

import (
	"encoding/json"
    "strings"
    "time"
	"fmt"
	"log"
	"materialcalculator/internal/orders"
	"net/http"
	"os"
	"path/filepath"
)

func StartServer() {

	router := http.NewServeMux()
	server := http.Server{
		Addr:    "localhost:8085",
		Handler: router,
	}

	filepath, _ := filepath.Abs("./frontend/")

	fmt.Println("looking for frontend path, ", filepath)
	fs := http.FileServer(http.Dir(filepath))

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
		http.Error(w, "Method now allowed", http.StatusMethodNotAllowed)
	}

	newOrder := orders.OrderData{}

	err := json.NewDecoder(r.Body).Decode(&newOrder)
	if err != nil {
		http.Error(w, "Error decoding json body: "+err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("%#v\n", newOrder)

	err = orders.NewOrder(newOrder)
	if err != nil {
		http.Error(w, "Error creating list file: "+err.Error(), http.StatusInternalServerError)
	}

	filePath := "./list.txt"
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to read file: %s", err)
        http.Error(w, "Failed to read file: "+err.Error(), http.StatusInternalServerError)
	}
	fileContent := string(bytes)

    http.ServeContent(w, r, "list.txt", time.Now(), strings.NewReader(fileContent))

    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Shed Materials Data received."))
}
