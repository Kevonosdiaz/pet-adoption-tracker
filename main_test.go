package main

import (
	// "database/sql"
	// "errors"
	// "log"
	// _ "modernc.org/sqlite"
	// "os"
	// "regexp"
	"slices"
	// "strconv"
	"testing"
)

// Quick checks to see if AnimalType enum and related helpers are consistent
func TestEnums(t *testing.T) {
	var animalStrings []string
	// We should be able to go back and forth and match original AnimalType
	for _, animalType := range allAnimalTypes {
		animalStr := animalNames[animalType]
		animalTypeFromConversion := stringToAnimalType[animalStr]
		// Also check all strings from animalNames matches getAllAnimalStrings() later
		animalStrings = append(animalStrings, animalStr)
		if animalType != animalTypeFromConversion {
			t.Errorf("Got %d AnimalType from conversion; wanted %d", animalTypeFromConversion, animalType)
		}
	}

	// Check if strings in animalNames match up with getAllAnimalStrings()
	allAnimalStrings := getAllAnimalStrings()
	for _, animalName := range animalStrings {
		if !slices.Contains(allAnimalStrings, animalName) {
			t.Errorf("Did not find %s animal name in result of getAllAnimalStrings()", animalName)
		}
	}
}

// Spin up a new test DB to use for each test
// func initTestDBHelper() *sql.DB {
// 	os.Remove("./test.db")
// 	// Setup SQLite DB connection and confirm connection is working
// 	db, err := sql.Open("sqlite", "./test.db")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	pingErr := db.Ping()
// 	if pingErr != nil {
// 		log.Fatal(pingErr)
// 	}
// 	defer db.Close()
// 	return db
// }
