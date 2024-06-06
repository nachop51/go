package pets

import "fmt"

type PetInterface interface {
	SetSpecies(string) *Pet
	SetBreed(string) *Pet
	SetMinWeight(int) *Pet
	SetMaxWeight(int) *Pet
	SetWeight(int) *Pet
	SetDescription(string) *Pet
	SetLifeSpan(string) *Pet
	SetGeographicOrigin(string) *Pet
	SetColor(string) *Pet
	SetAge(int) *Pet
	SetAgeEstimated(bool) *Pet
	Build() (*Pet, error)
}

func NewPetBuilder() PetInterface {
	return &Pet{}
}

func (p *Pet) SetSpecies(species string) *Pet {
	p.Species = species
	return p
}

func (p *Pet) SetBreed(breed string) *Pet {
	p.Breed = breed
	return p
}

func (p *Pet) SetMinWeight(minWeight int) *Pet {
	p.MinWeight = minWeight
	return p
}

func (p *Pet) SetMaxWeight(maxWeight int) *Pet {
	p.MaxWeight = maxWeight
	return p
}

func (p *Pet) SetWeight(weight int) *Pet {
	p.Weight = weight
	return p
}

func (p *Pet) SetDescription(description string) *Pet {
	p.Description = description
	return p
}

func (p *Pet) SetLifeSpan(lifeSpan string) *Pet {
	p.LifeSpan = lifeSpan
	return p
}

func (p *Pet) SetGeographicOrigin(geographicOrigin string) *Pet {
	p.GeographicOrigin = geographicOrigin
	return p
}

func (p *Pet) SetColor(color string) *Pet {
	p.Color = color
	return p
}

func (p *Pet) SetAge(age int) *Pet {
	p.Age = age
	return p
}

func (p *Pet) SetAgeEstimated(ageEstimated bool) *Pet {
	p.AgeEstimated = ageEstimated
	return p
}

func (p *Pet) Build() (*Pet, error) {
	// Validate the Pet struct
	if p.MinWeight > p.MaxWeight {
		return nil, fmt.Errorf("minimum weight cannot be greater than maximum weight")
	}

	if p.Age < 0 {
		return nil, fmt.Errorf("age cannot be negative")
	}

	p.AverageWeight = (p.MinWeight + p.MaxWeight) / 2

	// Return the Pet struct
	return p, nil
}
