package main

import (
	"database/sql"
	"errors"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"strconv"
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
	// First, find oldest animal matching type (and check if it exists)
	var animalId int
	var animalResult Animal
	animalType := stringToAnimalType[animalTypeString]
	selectQuery := "SELECT id, type, name FROM Animals WHERE type = ? ORDER BY time_added ASC LIMIT 1;"

	// Scan in resulting query info into appropriate variables
	err := db.QueryRow(selectQuery, animalType).Scan(&animalId, &animalResult.animal, &animalResult.name)

	// Error handling for above query
	// Check for specifically no animal returned; this is fine and we return an error msg
	if errors.Is(err, sql.ErrNoRows) {
		return "Sorry, but we don't have any animal types to adopt."
	}
	if err != nil {
		log.Fatal(err)
	}

	// Second, remove the found animal from the DB, using resulting animalId
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

// Query database for total number of animals, returning status message containing count
func getAllAnimalCount(db *sql.DB) string {
	var countResult int
	query := "SELECT COUNT(*) FROM Animals;"
	err := db.QueryRow(query).Scan(&countResult)

	if err != nil {
		log.Fatal(err)
	}

	// Make sure resulting int is converted to a string first to allow concatenation
	return "Total number of animals: " + strconv.Itoa(countResult)
}

// Query database for number of specified animal type, returning status message containing count
func getSpecificAnimalCount(db *sql.DB, animalTypeString string) string {
	var countResult int
	animalType := stringToAnimalType[animalTypeString]
	query := "SELECT COUNT(*) FROM Animals WHERE type = ?;"
	err := db.QueryRow(query, animalType).Scan(&countResult)

	if err != nil {
		log.Fatal(err)
	}

	return "Number of " + animalTypeString + "s: " + strconv.Itoa(countResult)
}

// Gets either total number of animals if animalTypeString is "all animals", else count for specific animal
func getAnimalCount(db *sql.DB, animalTypeString string) string {
	// Prepare our query, adding extra condition if animalTypeString is a specific type
	// Consider empty string from placeholder as "all animals" too
	if animalTypeString == "all animals" || animalTypeString == "" {
		return getAllAnimalCount(db)
	} else {
		return getSpecificAnimalCount(db, animalTypeString)
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
