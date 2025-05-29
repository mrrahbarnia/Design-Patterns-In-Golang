package pets

import "errors"

type PetInterface interface {
	SetSpecies(s string) *Pet
	SetBreed(s string) *Pet
	SetMinWeight(s int) *Pet
	SetMaxWeight(s int) *Pet
	SetWeight(s int) *Pet
	SetDescription(s string) *Pet
	SetLifeSpan(s int) *Pet
	SetGeographicOrigin(s string) *Pet
	SetColor(s string) *Pet
	SetAge(s int) *Pet
	SetAgeEstimated(ageEstimated bool) *Pet
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

func (p *Pet) SetLifeSpan(lifeSpan int) *Pet {
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
	if p.MinWeight > p.MaxWeight {
		return nil, errors.New("Min weight must be less than max weight.")
	}
	p.AverageWeight = (p.MinWeight + p.MaxWeight) / 2
	return p, nil
}
