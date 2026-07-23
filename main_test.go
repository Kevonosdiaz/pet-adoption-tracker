package main

// This contains all tests and test helpers for this app

import (
	"database/sql"
	"errors"
	"fmt"
	"math/rand/v2"
	_ "modernc.org/sqlite"
	"regexp"
	"slices"
	"testing"
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
	// IntN returns in interval [0,n) so we won't need len-1 to avoid OOB
	maxEnumVal := len(allAnimalTypes)
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

// Test data struct, stores slice of animals to add then check for
type FetchData struct {
	animals []Animal
}

// Helper to generate slices of test animal data, for 1..n animals
func generateRandomAnimals(n int) []Animal {
	// Offset to avoid 0, and allow n as result
	randn := rand.IntN(n) + 1
	res := make([]Animal, randn)
	for i := range res {
		res[i] = Animal{getRandomAnimalType(), generateRandomString()}
	}
	return res
}

// Test successful fetch of added animals into DB
func TestGetAdoptedAnimalSuccess(t *testing.T) {
	// Insert simple + random animal test data, then determine if oldest added animals are returned
	testData := []FetchData{
		{
			animals: []Animal{
				Animal{Dog, "dog1"},
				Animal{Dog, "dog2"},
			},
		},
		{
			animals: []Animal{
				Animal{Dog, "dog1"},
				Animal{Dog, "dog2"},
				Animal{Cat, "cat1"},
				Animal{Cat, "cat2"},
				Animal{Dog, "dog3"},
				Animal{Cat, "cat3"},
			},
		},
	}
	// Append 20 random FetchData structs containing anywhere from 1..20 Animals each
	for _ = range 20 {
		testData = append(testData, FetchData{generateRandomAnimals(20)})
	}

	// Use test data in "table driven" testing approach/loop
	numTests := len(testData)
	for i, tt := range testData {
		// Just use test case # since data is too long/complicated to print directly
		testname := fmt.Sprintf("%d/%d", i, numTests)

		// Execute "subtest" on one list of test animal data
		t.Run(testname, func(t *testing.T) {
			// Create a new in-memory DB for each run
			db := initTestDBHelper(t)
			defer db.Close()

			// First, insert data into DB directly, in same sequence as test data
			tx, err := db.Begin()
			if err != nil {
				t.Fatal(err)
			}

			// Prepare a statement for repeated usage
			stmt, err := tx.Prepare("INSERT INTO Animals(type, name) VALUES (?, ?)")
			if err != nil {
				t.Fatal(err)
			}

			defer stmt.Close()

			// Insert all animals in order, also separating out animals by type here for tracking oldest per type later
			numTypes := len(allAnimalTypes)
			// numOfEachType := make([]int, numTypes)
			animalsByType := make([][]Animal, numTypes)
			for _, a := range tt.animals {
				// Add animal to DB
				_, err = stmt.Exec(a.animal, a.name)

				// Take advantage of using AnimalType enum as index into animalsByType to filter animals by type
				animalsByType[a.animal] = append(animalsByType[a.animal], a)
			}

			// Commit the DB transaction
			err = tx.Commit()
			if err != nil {
				t.Fatal(err)
			}

			// Third, go through each type and adopt all animals of it until there are no more animals to adopt
			// Compile regex pattern for parsing result of getAdoptedAnimal() later, for alphanumeric names
			re := regexp.MustCompile(`Adopted animal is: ([a-zA-Z0-9]+)`)
			for i = range numTypes {
				animalTypeString := animalNames[AnimalType(i)]
				for _, a := range animalsByType[i] {
					resultMsg := getAdoptedAnimal(db, animalTypeString)
					// We are expecting the current animal name to match name mentioned in resultMsg (in first capture group, matches[1])
					matches := re.FindStringSubmatch(resultMsg)
					if matches[1] != a.name {
						t.Fatalf("Got %s, want %s from match", matches[1], a.name)
					}
				}
			}
		})
	}
}
