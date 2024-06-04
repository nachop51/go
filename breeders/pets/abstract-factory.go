package pets

import (
	"breeders/models"
	"errors"
)

type AnimalInterface interface {
	Show() string
}

type DogFromFactory struct {
	Pet *models.Dog
}

func (d *DogFromFactory) Show() string {
	return "This animal is a " + d.Pet.Breed.Breed
}

type CatFromFactory struct {
	Pet *models.Cat
}

func (c *CatFromFactory) Show() string {
	return "This animal is a " + c.Pet.Breed.Breed
}

type PetFactoryInterface interface {
	newPet(breed string) AnimalInterface
}

type DogAbstractFactory struct{}

func (d *DogAbstractFactory) newPet() AnimalInterface {
	return &DogFromFactory{
		Pet: &models.Dog{},
	}
}

type CatAbstractFactory struct{}

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
