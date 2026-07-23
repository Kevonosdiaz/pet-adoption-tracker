package main

// Contains data models and related functionality

// Define an enum to represent each animal type
// NOTE: AnimalType, allAnimalTypes, and animalNames should be updated together
// when adding/removing animal types
type AnimalType int

const (
	Dog AnimalType = iota
	Cat
)

// Keep track of all types so we can loop over them
var allAnimalTypes = []AnimalType{Dog, Cat}

// Allow mapping/conversion from AnimalType to string for printing as string
var animalNames = map[AnimalType]string{
	Dog: "dog",
	Cat: "cat",
}

// Implement String() for AnimalType for printing
func (animal AnimalType) String() string {
	return animalNames[animal]
}

func getAllAnimalStrings() []string {
	var animalTypeStrings []string
	for _, animal := range allAnimalTypes {
		animalTypeStrings = append(animalTypeStrings, animal.String())
	}
	return animalTypeStrings
}

// Data model
type Animal struct {
	animal AnimalType
	name   string
}
