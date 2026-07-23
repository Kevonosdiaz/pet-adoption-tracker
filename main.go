package main

import (
	"errors"
	// "fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"slices"
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
		// Send off form data to other function to handle database interaction
		// and clear fields to prepare for another submission
		OnSubmit: func() {
			addSurrenderedAnimal(animalName.Text, animalTypeSelector.Selected)
			animalName.SetText("")
			animalTypeSelector.ClearSelected()
		},
		SubmitText: "Add Animal",
		// Runs on form change to ensure animalType must be given before submitting
		Validator: func() error {
			allAnimalTypeStrings := getAllAnimalStrings()
			if slices.Contains(allAnimalTypeStrings, animalTypeSelector.Selected) {
				return nil
			}
			return errors.New("")
		},
	}

	// Ensure form Validator is re-run on selection change, since only Entry widget re-triggers Validator
	animalTypeSelector.OnChanged = func(s string) {
		form.Refresh()
	}

	form.Refresh()
	return form

}

// Create UI element allowing user to get adopted animal of selected type
func makeAdoptionForm() fyne.CanvasObject {
	animalTypeSelector := widget.NewSelect(getAllAnimalStrings(), func(s string) {})
	animalTypeSelector.PlaceHolder = "Select animal type"

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Animal Type", Widget: animalTypeSelector, Required: true},
		},
		// Send off form data to other function to handle database interaction
		// and clear fields to prepare for another submission
		OnSubmit: func() {
			getAdoptedAnimal(animalTypeSelector.Selected)
			animalTypeSelector.ClearSelected()
		},
		SubmitText: "Adopt Pet",
		Validator: func() error {
			allAnimalTypeStrings := getAllAnimalStrings()
			if slices.Contains(allAnimalTypeStrings, animalTypeSelector.Selected) {
				return nil
			}
			return errors.New("")
		},
	}

	// Ensure form Validator is re-run on selection change, since only Entry widget re-triggers Validator
	animalTypeSelector.OnChanged = func(s string) {
		form.Refresh()
	}

	form.Refresh()
	return form
}

func main() {
	// Setup Fyne GUI boilerplate
	myApp := app.New()
	window := myApp.NewWindow("Pet Adoption Tracker")

	// Create Fyne CanvasObject objects and corresponding labels for each UI element to add to the main window
	addSurrenderedAnimalText := widget.NewLabel("Add a New Surrendered Animal")
	addSurrenderedAnimalSection := makeSurrenderForm()
	getAdoptedAnimalText := widget.NewLabel("Pick an Animal to Adopt")
	getAdoptedAnimalSection := makeAdoptionForm()

	// Layout UI elements in vertical list
	windowContent := container.New(layout.NewVBoxLayout(), addSurrenderedAnimalText, addSurrenderedAnimalSection, getAdoptedAnimalText, getAdoptedAnimalSection)
	window.SetContent(windowContent)

	window.ShowAndRun()
}
