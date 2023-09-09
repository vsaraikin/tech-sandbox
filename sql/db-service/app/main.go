package main

import (
	"bufio"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"math/rand"
	"os"
	"time"
)

const (
	host     = "database"
	port     = 5432
	user     = "user"
	password = "password"
	dbname   = "db"
)

func main() {
	numberRows := 10000

	// Establish a connection to the PostgreSQL database
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create the "balances" table if it doesn't exist
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS balances (
			PersonID SERIAL PRIMARY KEY,
			LastName varchar(255),
			FirstName varchar(255),
			City varchar(255),
			Balance bigint
		);
	`)
	if err != nil {
		log.Fatal(err)
	}

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Load CSV data for realistic names and cities
	lastNames, err := loadCSV("last_names.csv")
	if err != nil {
		log.Fatal(err)
	}

	firstNames, err := loadCSV("first_names.csv")
	if err != nil {
		log.Fatal(err)
	}

	cities, err := loadCSV("cities.csv")
	if err != nil {
		log.Fatal(err)
	}

	// Generate and insert random data into the "balances" table with permuted names and cities
	for i := 1; i <= numberRows; i++ {
		permutedLastNames := permuteSlice(lastNames)
		permutedFirstNames := permuteSlice(firstNames)
		permutedCities := permuteSlice(cities)

		lastName := randomItem(permutedLastNames)
		firstName := randomItem(permutedFirstNames)
		city := randomItem(permutedCities)
		balance := rand.Int63n(10000)

		_, err := db.Exec("INSERT INTO balances (LastName, FirstName, City, Balance) VALUES ($1, $2, $3, $4)", lastName, firstName, city, balance)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Inserted record %d\n", i)
	}

	fmt.Printf("Completed!")
}

func loadCSV(filename string) ([]string, error) {
	var data []string
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return data, nil
}

func permuteSlice(slice []string) []string {
	permuted := make([]string, len(slice))
	perm := rand.Perm(len(slice))
	for i, randIndex := range perm {
		permuted[i] = slice[randIndex]
	}
	return permuted
}

func randomItem(items []string) string {
	index := rand.Intn(len(items))
	return items[index]
}
