package main

import (
    "bytes"
    "encoding/json"
    "net/http"
    "reflect"
    "testing"
)

func TestProcessDataHandler(t *testing.T) {
    testData := []int{1, 2, 3, 4, 5}
    testJSON, err := json.Marshal(testData)
    if err != nil {
        t.Fatalf("Error marshaling test data: %v", err)
    }

    resp, err := http.Post("/processData", "application/json", bytes.NewReader(testJSON))
    if err != nil {
        t.Errorf("Expected status code 200, got %d", resp.StatusCode)
    }

    var result []int
    err = json.NewDecoder(resp.Body).Decode(&result)
    if err != nil {
        t.Fatalf("Error decoding response body: %v", err)
    }

    expectedResult := []int{3, 4, 5}
    if !reflect.DeepEqual(result, expectedResult) {
        t.Errorf("Expected result %v, got %v", expectedResult, result)
    }
}
