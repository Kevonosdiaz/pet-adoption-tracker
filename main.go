package main

import (
	"database/sql"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Setup Fyne GUI boilerplate
	myApp := app.NewWithID("pet-adoption-tracker")
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

	// Initialize Animals table (if not yet created)
	initDB(db)

	// Create Fyne CanvasObject objects and corresponding subheading-sized labels for each UI element to add to the main window
	// Pass in SQLite DB handle so widgets can access it
	addSurrenderedAnimalText := widget.NewLabel("Add a New Surrendered Animal:")
	addSurrenderedAnimalText.SizeName = theme.SizeNameSubHeadingText
	addSurrenderedAnimalSection := makeSurrenderForm(db)

	getAdoptedAnimalText := widget.NewLabel("Pick an Animal to Adopt:")
	getAdoptedAnimalText.SizeName = theme.SizeNameSubHeadingText
	getAdoptedAnimalSection := makeAdoptionForm(db)

	animalCountText := widget.NewLabel("Check How Many Animals Are Up For Adoption:")
	animalCountText.SizeName = theme.SizeNameSubHeadingText
	animalCountSection := makeAnimalCountForm(db)

	// Layout UI elements in vertical list
	windowContent := container.New(layout.NewVBoxLayout(), addSurrenderedAnimalText, addSurrenderedAnimalSection, getAdoptedAnimalText, getAdoptedAnimalSection, animalCountText, animalCountSection)
	window.SetContent(windowContent)
	window.Resize(fyne.NewSize(500, 600))
	window.SetFixedSize(true)

	// Begin showing window and exec Fyne event loop
	window.ShowAndRun()
}
