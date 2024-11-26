package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// processData is the handler for the /processData endpoint.
func processData(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decode the request body
	var data []int
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	// Process the data (simply filter for values > 2 in this example)
	processedData := []int{}
	for _, value := range data {
		if value > 2 {
			processedData = append(processedData, value)
		}
	}

	// Encode the response body
	jsonResp, err := json.Marshal(processedData)
	if err != nil {
		http.Error(w, "Error encoding response body", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonResp)
	if err != nil {
		log.Printf("Error writing response: %v", err)
	}
}

func main() {
	http.HandleFunc("/processData", processData)
	fmt.Println("Server listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
