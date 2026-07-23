package main

// Contains functions to create complex widgets/UI elements

import (
	"database/sql"
	"errors"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	_ "github.com/mattn/go-sqlite3"
	"slices"
)

// Create UI element allowing user to add newly surrendered animal to database
func makeSurrenderForm(db *sql.DB) fyne.CanvasObject {
	// Create widgets for name text field and animal type selector dropdown
	animalName := widget.NewEntry()
	animalTypeSelector := widget.NewSelect(getAllAnimalStrings(), func(s string) {})
	animalTypeSelector.PlaceHolder = "Select animal type"

	// Display submission error/success message in this label
	resultText := widget.NewLabel("")

	// Add widgets to a form
	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Animal Name", Widget: animalName, Required: true},
			{Text: "Animal Type", Widget: animalTypeSelector, Required: true},
		},
		// Send off form data to other function to handle database interaction
		// and clear fields to prepare for another submission
		OnSubmit: func() {
			resultMsg := addSurrenderedAnimal(db, animalTypeSelector.Selected, animalName.Text)
			resultText.SetText(resultMsg)
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

	// Return the form followed by resultText label below it in a container
	return container.New(layout.NewVBoxLayout(), form, resultText)

}

// Create UI element allowing user to get adopted animal of selected type
func makeAdoptionForm(db *sql.DB) fyne.CanvasObject {
	animalTypeSelector := widget.NewSelect(getAllAnimalStrings(), func(s string) {})
	animalTypeSelector.PlaceHolder = "Select animal type"

	// Display submission error/success message in this label
	resultText := widget.NewLabel("")

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Animal Type", Widget: animalTypeSelector, Required: true},
		},
		// Send off form data to other function to handle database interaction
		// and clear fields to prepare for another submission
		OnSubmit: func() {
			resultMsg := getAdoptedAnimal(db, animalTypeSelector.Selected)
			resultText.SetText(resultMsg)
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

	// Return the form followed by resultText label below it in a container
	return container.New(layout.NewVBoxLayout(), form, resultText)
}

// Create UI element allowing user to get animal count for all animals or specific type
func makeAnimalCountForm(db *sql.DB) fyne.CanvasObject {
	// Allow user to pick "all animals" (default) to get count of all animals, or select a specific type instead
	animalTypeSelector := widget.NewSelect(append(getAllAnimalStrings(), "all animals"), func(s string) {})
	animalTypeSelector.Selected = "all animals"

	// Display resulting count in this label
	resultText := widget.NewLabel("")

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Animal Type", Widget: animalTypeSelector, Required: false},
		},
		// Send off form data to other function to handle database interaction
		// and clear fields to prepare for another submission
		OnSubmit: func() {
			resultMsg := getAnimalCount(db, animalTypeSelector.Selected)
			resultText.SetText(resultMsg)
		},
		SubmitText: "Get Count",
		// NOTE: No validator needed here since all options (incl. default/placeholder) are valid
	}

	// Return the form followed by resultText label below it in a container
	return container.New(layout.NewVBoxLayout(), form, resultText)
}
