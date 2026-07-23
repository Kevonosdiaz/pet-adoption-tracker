package main

// Contains functions to create complex widgets/UI elements

import (
	"database/sql"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	_ "github.com/mattn/go-sqlite3"
)

// Create UI element allowing user to add newly surrendered animal to database
func makeSurrenderForm(db *sql.DB) fyne.CanvasObject {
	// Create widgets for name text field and animal type selector dropdown
	allAnimalStrings := getAllAnimalStrings()
	animalName := widget.NewEntry()
	animalTypeSelector := widget.NewSelect(allAnimalStrings, func(s string) {})
	// Default to first animal type in list
	animalTypeSelector.Selected = allAnimalStrings[0]

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
		},
		SubmitText: "Add Animal",
	}

	// Return the form followed by resultText label below it in a container
	return container.New(layout.NewVBoxLayout(), form, resultText)

}

// Create UI element allowing user to get adopted animal of selected type
func makeAdoptionForm(db *sql.DB) fyne.CanvasObject {
	allAnimalStrings := getAllAnimalStrings()
	animalTypeSelector := widget.NewSelect(allAnimalStrings, func(s string) {})
	// Default to first animal type in list
	animalTypeSelector.Selected = allAnimalStrings[0]

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
		},
		SubmitText: "Adopt Pet",
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
	}

	// Return the form followed by resultText label below it in a container
	return container.New(layout.NewVBoxLayout(), form, resultText)
}
