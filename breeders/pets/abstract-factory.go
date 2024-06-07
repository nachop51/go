package pets

import (
	"breeders/config"
	"breeders/models"
	"errors"
)

type AnimalInterface interface {
	Show() string
}

type DogFromFactory struct {
	Pet *models.Dog
}

type CatFromFactory struct {
	Pet *models.Cat
}

func (d *DogFromFactory) Show() string {
	return "This animal is a " + d.Pet.Breed.Breed
}

func (c *CatFromFactory) Show() string {
	return "This animal is a " + c.Pet.Breed.Breed
}

type PetFactoryInterface interface {
	newPet(breed string) AnimalInterface
	newPetWithBreed(breed string) AnimalInterface
}

type DogAbstractFactory struct{}
type CatAbstractFactory struct{}

func (d *DogAbstractFactory) newPet() AnimalInterface {
	return &DogFromFactory{
		Pet: &models.Dog{},
	}
}

func (d *DogAbstractFactory) newPetWithBreed(b string) AnimalInterface {
	app := config.GetInstance()

	breed, err := app.Models.DogBreed.GetBreedByName(b)

	if err != nil {
		return nil
	}

	return &DogFromFactory{
		Pet: &models.Dog{
			Breed: *breed,
		},
	}
}

func (c *CatAbstractFactory) newPetWithBreed(b string) AnimalInterface {
	// app := config.GetInstance()

	// breed, err := app.catBreed.GetBreedByName(b)

	// if err != nil {
	// 	return nil, err
	// }

	// return &CatFromFactory{
	// 	Pet: &models.Cat{
	// 		Breed: breed,
	// 	},
	// }
	return nil
}

func (c *CatAbstractFactory) newPet() AnimalInterface {
	return &CatFromFactory{
		Pet: &models.Cat{},
	}
}

func NewPetFromAbstractFactory(species string) (AnimalInterface, error) {
	switch species {
	case "dog":
		return (&DogAbstractFactory{}).newPet(), nil
	case "cat":
		return (&CatAbstractFactory{}).newPet(), nil
	default:
		return nil, errors.New("invalid species")
	}
}

func NewPetWithBreedFromAbstractFactory(species string, breed string) (AnimalInterface, error) {
	// switch species {
	// case "dog":
	// 	return (&DogAbstractFactory{}).newPetWithBreed(), nil
	// case "cat":
	// 	return (&CatAbstractFactory{}).newPetWithBreed(), nil
	// default:
	// 	return nil, errors.New("invalid species")
	// }
	return nil, nil
}
