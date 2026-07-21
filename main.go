package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func addSurrenderedAnimal(animalName string, animalTypeString string) {
	// Do we convert back from string to AnimalType here?
}

// Create UI element allowing user to add newly surrendered animal to database
func makeSurrenderForm() fyne.CanvasObject {
	// Create widgets for name text field and animal type selector dropdown
	animalName := widget.NewEntry()
	animalTypeSelector := widget.NewSelect(getAllAnimalStrings(), func(s string) {})
	animalTypeSelector.PlaceHolder = "Select animal type"

	// Add widgets to a form
	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Animal Name", Widget: animalName, Required: true},
			{Text: "Animal Type", Widget: animalTypeSelector, Required: true},
		},
		// OnCancel will act as clear button
		OnCancel: func() {
			animalName.Refresh()
			animalTypeSelector.Refresh()
		},
		// Send off form data to other function to handle database interaction
		OnSubmit: func() {
			addSurrenderedAnimal(animalName.Text, animalTypeSelector.Selected)
		},
		CancelText: "Clear",
		SubmitText: "Submit",
	}
	form.Refresh()
	return form

}

func makeAdoptionForm() fyne.CanvasObject {
	return nil
}

func main() {
	// Setup Fyne GUI boilerplate
	myApp := app.New()
	window := myApp.NewWindow("Pet Adoption Tracker")

	window.ShowAndRun()
}
