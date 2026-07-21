package main

import (
	"fmt"
	"fyne.io/fyne/v2/app"
	// "fyne.io/fyne/v2/widget"
)

func main() {
	// Setup Fyne GUI boilerplate
	myApp := app.New()
	window := myApp.NewWindow("Pet Adoption Tracker")

	fmt.Println(Dog)

	window.ShowAndRun()
}
