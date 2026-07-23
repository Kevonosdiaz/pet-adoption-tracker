package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

// Functions that interact with the SQLite DB

// Add new animal with type animalTypeString and name animalName to the DB, returning status message as string
func addSurrenderedAnimal(db *sql.DB, animalTypeString string, animalName string) string {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	// Protect against an early exit without Commit() being called successfully
	defer tx.Rollback()

	// Prepare our query
	query := "INSERT INTO Animals(type, name) VALUES (?, ?)"

	// Execute query with values; animalTypeString is converted back to an enum (int) first
	_, err = tx.Exec(query, stringToAnimalType[animalTypeString], animalName)
	if err != nil {
		log.Fatal(err)
	}

	// Commit it to DB
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
	return "Successfully added animal!"
}

func getAdoptedAnimal(db *sql.DB, animalTypeString string) {
	// tx, err := db.Begin()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// // First check if there is an animal we can get
}

// Gets either total number of animals if animalTypeString is "all animals", else count for specific animal
func getAnimalCount(db *sql.DB, animalTypeString string) {
	// Prepare our query, adding extra condition if animalTypeString is a specific type
	if animalTypeString == "all animals" {
		// return getAllAnimalCount(db)
	} else {
		// return getSpecificAnimalCount(db, animalTypeString)
	}
}

// Create main table if it does not exist
// time_added is stored as integer to easily determine oldest animal
func initDB(db *sql.DB) {
	sqlStmt := `CREATE TABLE IF NOT EXISTS Animals (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		type INTEGER NOT NULL,
		name TEXT NOT NULL,
		time_added INTEGER DEFAULT (unixepoch())
	);
	`
	_, err := db.Exec(sqlStmt)
	if err != nil {
		log.Fatal(err)
	}
}
