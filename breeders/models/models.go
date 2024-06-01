package models

import "time"

type DogBreeds struct {
	Id               int    `json:"id"`
	Breed            string `json:"breed"`
	WeightLowLbs     int    `json:"weight_low_lbs"`
	WeightHighLbs    int    `json:"weight_high_lbs"`
	AverageWeight    int    `json:"average_weight"`
	LifeSpan         string `json:"average_lifespan"`
	Details          string `json:"details"`
	AlternateNames   string `json:"alternate_names"`
	GeographicOrigin string `json:"geographic_origin"`
}
type CatBreeds struct {
	Id               int    `json:"id"`
	Breed            string `json:"breed"`
	WeightLowLbs     int    `json:"weight_low_lbs"`
	WeightHighLbs    int    `json:"weight_high_lbs"`
	AverageWeight    int    `json:"average_weight"`
	LifeSpan         string `json:"average_lifespan"`
	Details          string `json:"details"`
	AlternateNames   string `json:"alternate_names"`
	GeographicOrigin string `json:"geographic_origin"`
}

type Dog struct {
	Id               int       `json:"id"`
	Name             string    `json:"name"`
	BreedId          int       `json:"breed_id"`
	BreederId        int       `json:"breeder_id"`
	Color            string    `json:"color"`
	DateOfBirth      time.Time `json:"date_of_birth"`
	SpayedOrNeutered bool      `json:"spayed_or_neutered"`
	Descrption       string    `json:"description"`
	Weight           int       `json:"weight"`
	Breed            DogBreeds `json:"breed"`
	Breeder          Breeder   `json:"breeder"`
}

type Cat struct {
	Id               int       `json:"id"`
	Name             string    `json:"name"`
	BreedId          int       `json:"breed_id"`
	BreederId        int       `json:"breeder_id"`
	Color            string    `json:"color"`
	DateOfBirth      time.Time `json:"date_of_birth"`
	SpayedOrNeutered bool      `json:"spayed_or_neutered"`
	Descrption       string    `json:"description"`
	Weight           int       `json:"weight"`
	Breed            CatBreeds `json:"breed"`
	Breeder          Breeder   `json:"breeder"`
}

type Breeder struct {
	Id          int          `json:"id"`
	BreederName string       `json:"breeder_name"`
	Address     string       `json:"address"`
	City        string       `json:"city"`
	ProvState   string       `json:"prov_state"`
	Country     string       `json:"country"`
	Zip         string       `json:"zip"`
	Phone       string       `json:"phone"`
	Email       string       `json:"email"`
	Active      bool         `json:"active"`
	DogBreeds   []*DogBreeds `json:"dog_breeds"`
	CatBreeds   []*CatBreeds `json:"cat_breeds"`
}

type Pet struct {
	Species     string `json:"species"`
	Breed       string `json:"breed"`
	MinWeight   int    `json:"min_weight"`
	MaxWeight   int    `json:"max_weight"`
	Description string `json:"description"`
	LifeSpan    string `json:"lifespan"`
}
