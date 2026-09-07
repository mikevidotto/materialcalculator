package main

import (
	"encoding/json"
	"log"
	"net/http"
	"path/filepath"
)

type ShedDimensions struct {
	Length float64 `json:"length"`
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}

func CalculateMaterials(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var dims ShedDimensions
	if err := json.NewDecoder(r.Body).Decode(&dims); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	// Replace this sample response with your real shed material calculation.
	response := map[string]any{
		"length": dims.Length,
		"width":  dims.Width,
		"height": dims.Height,
		"message": "dimensions received",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	frontendPath, err := filepath.Abs("../../frontend")
	if err != nil {
		log.Fatal(err)
	}

	fs := http.FileServer(http.Dir(frontendPath))

	http.HandleFunc("/calculate", CalculateMaterials)
	http.Handle("/", fs)

	log.Println("starting server on http://localhost:8085")
	log.Fatal(http.ListenAndServe(":8085", nil))
}
