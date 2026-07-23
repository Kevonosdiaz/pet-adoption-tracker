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
	// Prepare our query
	stmt, err := tx.Prepare("INSERT INTO Animals(type, name) VALUES (?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	// Execute query with values; animalTypeString is converted back to an enum (int) first
	_, err = stmt.Exec(stringToAnimalType[animalTypeString], animalName)
	if err != nil {
		log.Fatal(err)
	}
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
	return "Successfully added animal!"
}

func getAdoptedAnimal(db *sql.DB, animalTypeString string) {

}

func getAnimalCount(db *sql.DB, animalTypeString string) {

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
