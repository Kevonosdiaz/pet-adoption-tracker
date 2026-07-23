package main

// Contains data models and related functionality

// NOTE: AnimalType, allAnimalTypes, animalNames, and stringToAnimalName should be updated together
// when adding/removing animal types
// Defines an enum to represent each animal type
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

var stringToAnimalType = map[string]AnimalType{
	"dog": Dog,
	"cat": Cat,
}

// NOTE: Consider using official stringer tool to implement String() automatically
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
