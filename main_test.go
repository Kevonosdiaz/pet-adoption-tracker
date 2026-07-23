package main

import (
	"database/sql"
	"errors"
	"fmt"
	"math/rand/v2"
	_ "modernc.org/sqlite"
	"slices"
	"testing"
	// "regexp"
	// "strconv"
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

// Test helper functions

// Spin up a new DB in memory to use for tests, should be closed by caller
func initTestDBHelper(t *testing.T) *sql.DB {
	// Declare as a helper function for tests
	t.Helper()

	// Setup SQLite DB connection and confirm connection is working
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	pingErr := db.Ping()
	if pingErr != nil {
		t.Fatal(pingErr)
	}

	// Create table here
	initDB(db)

	return db
}

// Helper to generate random strings for test names
func generateRandomString() string {
	// Allow length anywhere from 1..64 characters and just stick to alpha chars
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	n := rand.N(63) + 1
	b := make([]byte, n)
	// Pick a (pseudo)random letter from letterBytes for each char in b
	for i := range b {
		b[i] = letterBytes[rand.Int64()%int64(len(letterBytes))]
	}

	return string(b)

}

type InsertData struct {
	animalTypeString string
	animalName       string
}

// Helper to randomly pick AnimalType for tests
func getRandomAnimalType() AnimalType {
	// Account for iota beginning from 0
	maxEnumVal := len(allAnimalTypes) - 1
	n := rand.IntN(maxEnumVal)
	return AnimalType(n)
}

// Test insertion of new animals into DB
func TestAddSurrenderedAnimal(t *testing.T) {
	db := initTestDBHelper(t)
	defer db.Close()

	// Insert simple + random animal test data and expect to be able to SELECT it in DB
	testData := []InsertData{
		{"dog", "dog1"},
		{"dog", "dog2"},
		{"dog", "dog3"},
		{"dog", "dog4"},
		{"cat", "cat1"},
		{"cat", "cat2"},
		{"cat", "cat3"},
		{"cat", "cat4"},
		{"dog", "dog5"},
		{"cat", "cat5"},
	}
	// Append 100 random animal data elements
	for _ = range 100 {
		testData = append(testData, InsertData{generateRandomString(), getRandomAnimalType().String()})
	}

	// Use test data in "table driven" testing approach/loop
	for _, tt := range testData {
		testname := fmt.Sprintf("%s,%s", tt.animalTypeString, tt.animalName)

		// Execute "subtest" on one data element from testData
		t.Run(testname, func(t *testing.T) {
			// Call addSurrenderedAnimal() on test data
			addSurrenderedAnimal(db, tt.animalTypeString, tt.animalName)

			// Confirm animal was added into DB
			animalType := stringToAnimalType[tt.animalTypeString]
			query := "SELECT * FROM Animals WHERE type = ? AND name = ?;"
			// Ignore returned value, just ensure any row was returned
			err := db.QueryRow(query, animalType, tt.animalName).Scan()
			if errors.Is(err, sql.ErrNoRows) {
				t.Error("Unable to find inserted animal in DB")
			}
		})
	}
}
