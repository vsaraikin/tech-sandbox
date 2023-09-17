package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	_ "github.com/lib/pq"
	"inserter/data"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"
)

const (
	host     = "database"
	port     = 5432
	user     = "user"
	password = "password"
	dbname   = "db"
)

var numberRows = 150000

func main() {

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	conn, err := pgx.Connect(context.Background(), connStr)

	if err != nil {
		log.Fatal(err)
	}
	defer func(conn *pgx.Conn, ctx context.Context) {
		err := conn.Close(ctx)
		if err != nil {

		}
	}(conn, context.Background())

	q := data.New(conn)

	_, err = conn.Exec(context.Background(), `
	CREATE TABLE IF NOT EXISTS balances (
		PersonID SERIAL PRIMARY KEY,
		LastName varchar(255),
		FirstName varchar(255),
		City varchar(255),
		Balance bigint);
	`)
	if err != nil {
		log.Fatal(err)
	}

	rand.Seed(824)

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

	var wg sync.WaitGroup
	var m sync.Mutex
	var dataToInsert []data.InsertBalancesParams

	permutedLastNames := permuteSlice(lastNames)
	permutedFirstNames := permuteSlice(firstNames)
	permutedCities := permuteSlice(cities)

	start := time.Now()
	for i := 1; i <= numberRows; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			m.Lock()
			defer m.Unlock()
			dataToInsert = append(dataToInsert, data.InsertBalancesParams{
				Firstname: randomItem(permutedFirstNames),
				Lastname:  randomItem(permutedLastNames),
				City:      randomItem(permutedCities),
				Balance:   rand.Int63n(10000),
			},
			)

			if err != nil {
				log.Fatal(err)
			}
		}(i)
	}

	wg.Wait()
	result := q.InsertBalances(context.Background(), dataToInsert)
	result.Exec(func(i int, err error) {
		if err != nil {
			log.Printf("error updating book %d: %s\n", i, err)
		}
	})
	log.Printf("Completed in %s\n", time.Since(start))
}

func loadCSV(filename string) ([]string, error) {
	var outputData []string
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		outputData = append(outputData, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return outputData, nil
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
