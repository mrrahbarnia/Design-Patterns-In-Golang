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

func (app *application) CreateDogFromAbstractFactory(w http.ResponseWriter, r *http.Request) {
	dog, err := pets.NewPetFromAbstractFactory("dog")
	if err != nil {
		http.Error(w, "Failed to create dog from abstract factory", http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(dog); err != nil {
		http.Error(w, "Failed to encode dog", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (app *application) CreateCatFromAbstractFactory(w http.ResponseWriter, r *http.Request) {
	cat, err := pets.NewPetFromAbstractFactory("cat")
	if err != nil {
		http.Error(w, "Failed to create cat from abstract factory", http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(cat); err != nil {
		http.Error(w, "Failed to encode cat", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (app *application) GetDogBreads(w http.ResponseWriter, r *http.Request) {
	dogBreeds, err := app.Models.DogBreeds.All()
	if err != nil {
		http.Error(w, "Failed to get dog breeds", http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(dogBreeds); err != nil {
		http.Error(w, "Failed to encode dog breeds", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
