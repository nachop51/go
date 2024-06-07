package adapters

import (
	"breeders/models"
	"encoding/json"
	"encoding/xml"
	"net/http"
)

type CatBreedsInterface interface {
	GetAllCatBreeds() ([]*models.CatBreed, error)
	GetCatBreedByName(string) (*models.CatBreed, error)
}

type RemoteService struct {
	Remote CatBreedsInterface
}

func (r *RemoteService) GetAllCatBreeds() ([]*models.CatBreed, error) {
	return r.Remote.GetAllCatBreeds()
}

type JSONBackend struct{}

func (j *JSONBackend) GetAllCatBreeds() ([]*models.CatBreed, error) {
	resp, err := http.Get("https://api.thecatapi.com/v1/breeds")

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var breeds []*models.CatBreed

	err = json.NewDecoder(resp.Body).Decode(&breeds)

	if err != nil {
		return nil, err
	}

	return breeds, nil
}

type XMLBackend struct{}

func (x *XMLBackend) GetAllCatBreeds() ([]*models.CatBreed, error) {
	resp, err := http.Get("https://api.thecatapi.com/v1/breeds")

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var breeds []*models.CatBreed

	err = xml.NewDecoder(resp.Body).Decode(&breeds)

	if err != nil {
		return nil, err
	}

	return breeds, nil
}

func (r *RemoteService) GetCatBreedByName(breed string) (*models.CatBreed, error) {
	return r.Remote.GetCatBreedByName(breed)
}

func (j *JSONBackend) GetCatBreedByName(breed string) (*models.CatBreed, error) {
	resp, err := http.Get("https://api.thecatapi.com/v1/breeds")

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var breeds []*models.CatBreed

	err = json.NewDecoder(resp.Body).Decode(&breeds)

	if err != nil {
		return nil, err
	}

	for _, b := range breeds {
		if b.Breed == breed {
			return b, nil
		}
	}

	return nil, nil
}

func (x *XMLBackend) GetCatBreedByName(breed string) (*models.CatBreed, error) {
	resp, err := http.Get("https://api.thecatapi.com/v1/breeds")

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var breeds []*models.CatBreed

	err = xml.NewDecoder(resp.Body).Decode(&breeds)

	if err != nil {
		return nil, err
	}

	for _, b := range breeds {
		if b.Breed == breed {
			return b, nil
		}
	}

	return nil, nil
}
