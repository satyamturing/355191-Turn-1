package main

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandleSum(t *testing.T) {
	rec := new(record)
	//Create a large dataset for testing
	largeData := []int{1, 2, 3..... 10000}
	w := new(testing.ResponseRecorder)
	r := new(testing.Request)
	// Here, you would marshal the largeData into JSON and set it on the request body.
	// For simplicity, we'll skip that in this initial step.

	rec.HandleSum(w, r)

	require.Equal(t, http.StatusOK, w.Code)
	var response struct {
		Sum int `json:"sum"`
	}
	err := json.NewDecoder(w.Body).Decode(&response)
	require.NoError(t, err)

	// Test the sum of large dataset
	assert.Equal(t, calculateSum(largeData), response.Sum)
}
