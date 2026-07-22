package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func addSurrenderedAnimal(animalName string, animalTypeString string) {
	// Do we convert back from string to AnimalType here?
}

func getAdoptedAnimal(animalTypeString string) {

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
		// Cancel will act as clear button
		// OnCancel: func() {
		// 	animalName.Refresh()
		// 	animalTypeSelector.Refresh()
		// },
		// Send off form data to other function to handle database interaction
		OnSubmit: func() {
			addSurrenderedAnimal(animalName.Text, animalTypeSelector.Selected)
		},
		// CancelText: "Clear",
		SubmitText: "Submit",
	}
	form.Refresh()
	return form

}

func makeAdoptionForm() fyne.CanvasObject {
	animalTypeSelector := widget.NewSelect(getAllAnimalStrings(), func(s string) {})
	animalTypeSelector.PlaceHolder = "Select animal type"

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Animal Type", Widget: animalTypeSelector, Required: true},
		},
		// Cancel will act as clear button
		// OnCancel: func() {
		// 	animalTypeSelector.Refresh()
		// },
		// Send off form data to other function to handle database interaction
		OnSubmit: func() {
			getAdoptedAnimal(animalTypeSelector.Selected)
		},
		// CancelText: "Clear",
		SubmitText: "Submit",
	}
	form.Refresh()
	return form
}

func main() {
	// Setup Fyne GUI boilerplate
	myApp := app.New()
	window := myApp.NewWindow("Pet Adoption Tracker")

	addSurrenderedAnimalSection := makeSurrenderForm()
	getAdoptedAnimalSection := makeAdoptionForm()

	windowContent := container.New(layout.NewVBoxLayout(), addSurrenderedAnimalSection, getAdoptedAnimalSection)
	window.SetContent(windowContent)

	window.ShowAndRun()
}
