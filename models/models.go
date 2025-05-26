package models

import "time"

type DogBreed struct {
	ID               int    `json:"id"`
	Breed            string `json:"breed"`
	WeightLowLbs     int    `json:"weight_low_lbs"`
	WeightHighLbs    int    `json:"weight_high_lbs"`
	AverageWeight    int    `json:"average_weight"`
	LifeSpan         int    `json:"lifespan"`
	Description      string `json:"description"`
	AlternateNames   string `json:"alternate_names"`
	GeographicOrigin string `json:"geographic_origin"`
}

type CatBreed struct {
	ID               int    `json:"id"`
	Breed            string `json:"breed"`
	WeightLowLbs     int    `json:"weight_low_lbs"`
	WeightHighLbs    int    `json:"weight_high_lbs"`
	AverageWeight    int    `json:"average_weight"`
	LifeSpan         int    `json:"lifespan"`
	Description      string `json:"description"`
	AlternateNames   string `json:"alternate_names"`
	GeographicOrigin string `json:"geographic_origin"`
}

type Dog struct {
	ID               int       `json:"id"`
	BreedID          int       `json:"breed_id"`
	DogName          string    `json:"dog_name"`
	Color            string    `json:"color"`
	BreederID        int       `json:"breeder_id"`
	DateOfBirth      time.Time `json:"date_of_birth"`
	SpayedOrNeutered bool      `json:"spayed_or_neutered"`
	Description      string    `json:"description"`
	Weight           int       `json:"weight"`
	Breed            DogBreed  `json:"breed"`
	Breerder         Breeder   `json:"breeder"`
}

type Cat struct {
	ID               int       `json:"id"`
	BreedID          int       `json:"breed_id"`
	CatName          string    `json:"cat_name"`
	Color            string    `json:"color"`
	BreederID        int       `json:"breeder_id"`
	DateOfBirth      time.Time `json:"date_of_birth"`
	SpayedOrNeutered bool      `json:"spayed_or_neutered"`
	Description      string    `json:"description"`
	Weight           int       `json:"weight"`
	Breed            DogBreed  `json:"breed"`
	Breerder         Breeder   `json:"breeder"`
}

type Breeder struct {
	ID            int         `json:"id"`
	BreederName   string      `json:"breeder_name"`
	Address       string      `json:"address"`
	City          string      `json:"city"`
	ProvState     string      `json:"prov_state"`
	Country       string      `json:"country"`
	ZipPostalCode string      `json:"zip_postal_code"`
	Phone         string      `json:"phone"`
	Email         string      `json:"email"`
	Active        bool        `json:"active"`
	DogBreeds     []*DogBreed `json:"dog_breeds"`
	CatBreeds     []*CatBreed `json:"cat_breeds"`
}

type Pet struct {
	Species     string `json:"species"` // "dog" or "cat" or something else
	Breed       string `json:"breed"`
	MinWeight   int    `json:"min_weight"`
	MaxWeight   int    `json:"max_weight"`
	Description string `json:"description"`
	LifeSpan    int    `json:"lifespan"`
}
