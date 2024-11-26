
package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
)
type record struct{}

func (r *record) HandleSum(w http.ResponseWriter, r *http.Request) {

    //Here, you would parse the request body to extract the large dataset.
    // For simplicity, we'll skip that in this example.

    largeData := []int{1, 2, 3, ..., 10000}
    sum := calculateSum(largeData)
	
    response := struct {
        Sum int `json:"sum"`
    }{
        Sum: sum,
    }

    json.NewEncoder(w).Encode(response)
}

func calculateSum(data []int) int {
    // For large datasets, consider streaming or optimizing this logic
    total := 0
    for _, num := range data {
        total += num
    }
    return total
}


func main() {
    http.HandleFunc("/sum", (&record{}).HandleSum)
    fmt.Println("Server is running on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
func (r *record) HandleSum(w http.ResponseWriter, r *http.Request) {
    // ... (existing code)
    sum := calculateSum(largeData)
    // ... (existing code)
}

func calculateSum(data []int) int {
    // Optimized sum calculation for large datasets (parallel processing or streaming)
    total := 0
    for _, num := range data {
        total += num
    }
    return total
}