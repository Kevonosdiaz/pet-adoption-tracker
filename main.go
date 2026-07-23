package main

import (
	"database/sql"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func addSurrenderedAnimal(db *sql.DB, animalName string, animalTypeString string) {
	// Do we convert back from string to AnimalType here?
}

func getAdoptedAnimal(db *sql.DB, animalTypeString string) {

}

func getAnimalCount(db *sql.DB, animalTypeString string) {

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
