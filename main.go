package main

import (
	"database/sql"
	"errors"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"slices"
)

func addSurrenderedAnimal(db *sql.DB, animalName string, animalTypeString string) {
	// Do we convert back from string to AnimalType here?
}

func getAdoptedAnimal(db *sql.DB, animalTypeString string) {

}

func getAnimalCount(db *sql.DB, animalTypeString string) {

}

// Create UI element allowing user to add newly surrendered animal to database
func makeSurrenderForm(db *sql.DB) fyne.CanvasObject {
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
			addSurrenderedAnimal(db, animalName.Text, animalTypeSelector.Selected)
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
func makeAdoptionForm(db *sql.DB) fyne.CanvasObject {
	animalTypeSelector := widget.NewSelect(getAllAnimalStrings(), func(s string) {})
	animalTypeSelector.PlaceHolder = "Select animal type"

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Animal Type", Widget: animalTypeSelector, Required: true},
		},
		// Send off form data to other function to handle database interaction
		// and clear fields to prepare for another submission
		OnSubmit: func() {
			getAdoptedAnimal(db, animalTypeSelector.Selected)
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

// Create UI element allowing user to get animal count for all animals or specific type
func makeAnimalCountForm(db *sql.DB) fyne.CanvasObject {
	// Allow user to pick "all animals" (default) to get count of all animals, or select a specific type instead
	animalTypeSelector := widget.NewSelect(append(getAllAnimalStrings(), "all animals"), func(s string) {})
	animalTypeSelector.PlaceHolder = "all animals"

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Animal Type", Widget: animalTypeSelector, Required: false},
		},
		// Send off form data to other function to handle database interaction
		// and clear fields to prepare for another submission
		OnSubmit: func() {
			getAnimalCount(db, animalTypeSelector.Selected)
			animalTypeSelector.ClearSelected()
		},
		SubmitText: "Get Count",
		// NOTE: No validator needed here since all options (incl. default/placeholder) are valid
	}

	form.Refresh()
	return form
}

func main() {
	// Setup Fyne GUI boilerplate
	myApp := app.New()
	window := myApp.NewWindow("Pet Adoption Tracker")

	// Setup SQLite DB connection and confirm connection is working
	db, err := sql.Open("sqlite3", "./animals.db")
	if err != nil {
		log.Fatal(err)
	}
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	defer db.Close()

	// Create Fyne CanvasObject objects and corresponding labels for each UI element to add to the main window
	// Pass in SQLite DB handle so widgets can access it
	addSurrenderedAnimalText := widget.NewLabel("Add a New Surrendered Animal:")
	addSurrenderedAnimalSection := makeSurrenderForm(db)

	getAdoptedAnimalText := widget.NewLabel("Pick an Animal to Adopt:")
	getAdoptedAnimalSection := makeAdoptionForm(db)

	animalCountText := widget.NewLabel("Check How Many Animals Are Up For Adoption:")
	animalCountSection := makeAnimalCountForm(db)

	// Layout UI elements in vertical list
	windowContent := container.New(layout.NewVBoxLayout(), addSurrenderedAnimalText, addSurrenderedAnimalSection, getAdoptedAnimalText, getAdoptedAnimalSection, animalCountText, animalCountSection)
	window.SetContent(windowContent)

	// Begin showing window and exec Fyne event loop
	window.ShowAndRun()
}
