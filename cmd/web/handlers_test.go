package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApplication_GetDogBreads(t *testing.T) {
	// Create a request
	req, _ := http.NewRequest("GET", "/api/v1/dog-breads", nil)

	// Create a response recorder
	rr := httptest.NewRecorder()

	// Create a handler
	handler := http.HandlerFunc(testApp.GetDogBreads)

	// Serve the handlerw
	handler.ServeHTTP(rr, req)

	// Check response
	if rr.Code != http.StatusOK {
		t.Errorf("Wrong response code, got %d, wanted 200.", rr.Code)
	}
}
