package main

// Define an enum to represent each animal type
type AnimalType int

const (
	Dog AnimalType = iota
	Cat
)

// Allow mapping/conversion from AnimalType to string for printing as string
var animalNames = map[AnimalType]string{
	Dog: "dog",
	Cat: "cat",
}

// Implement String() for AnimalType for printing
func (animal AnimalType) String() string {
	return animalNames[animal]
}

// Data model
type Animal struct {
	animal AnimalType
	name   string
}
