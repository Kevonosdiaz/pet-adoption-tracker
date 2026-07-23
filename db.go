package main

import (
	"database/sql"
	"errors"
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

	// Create our query
	query := "INSERT INTO Animals(type, name) VALUES (?, ?)"

	// Execute query with values; animalTypeString is converted back to an enum (int) first
	_, err = tx.Exec(query, stringToAnimalType[animalTypeString], animalName)
	if err != nil {
		log.Fatal(err)
	}

	// Commit the write transaction to DB
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
	return "Successfully added animal!"
}

// Fetch oldest animal of specified type and remove it (if animal of type exists), returning status message either with animal details or error
func getAdoptedAnimal(db *sql.DB, animalTypeString string) string {
	// First find oldest animal matching type (and check if it exists)
	var animalId int
	var animalResult Animal
	animalType := stringToAnimalType[animalTypeString]
	selectQuery := "SELECT id, type, name FROM Animals WHERE type = ? ORDER BY time_added ASC LIMIT 1;"
	err := db.QueryRow(selectQuery, animalType).Scan(&animalId, &animalResult.animal, &animalResult.name)

	// Error handling for above query
	if errors.Is(err, sql.ErrNoRows) {
		return "Sorry, but we don't have any animal types to adopt."
	}
	if err != nil {
		log.Fatal(err)
	}

	// Remove the found animal from the DB, using resulting animalId
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	defer tx.Rollback()

	// Create our query
	deleteQuery := "DELETE FROM Animals WHERE id = ?;"

	// Execute deletion query on found animal
	_, err = tx.Exec(deleteQuery, animalId)
	if err != nil {
		log.Fatal(err)
	}

	// Commit the write transaction to DB
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	return "Adopted animal is: " + animalResult.name
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
