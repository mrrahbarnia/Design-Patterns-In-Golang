package main

import (
	"encoding/json"
	"net/http"

	"github.com/mrrahbarnia/Design-Patterns-In-Golang/pets"
)

func (app *application) CheckHealth(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Everything is healthy!"))
}

func (app *application) CreateDogFromFactory(w http.ResponseWriter, r *http.Request) {
	dog := pets.NewPet("dog")
	if err := json.NewEncoder(w).Encode(dog); err != nil {
		http.Error(w, "Failed to encode dog", http.StatusInternalServerError)
		return
	}
}
