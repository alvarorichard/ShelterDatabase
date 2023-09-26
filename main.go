package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	//"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"strconv"
)

func insertIntoDatabase(pool *pgxpool.Pool, id string, folderName string) {
    q := "INSERT into logs (id, folder_name) values ($1, $2);"
    _, err := pool.Exec(context.Background(), q, id, folderName)
    if err != nil {
        panic(err)
    }
}
func getNextID(pool *pgxpool.Pool) string {
    var maxID int
    err := pool.QueryRow(context.Background(), "SELECT COALESCE(MAX(id)::int, 0) FROM logs").Scan(&maxID)
    if err != nil {
        panic(err)
    }
    return strconv.Itoa(maxID + 1)
}


func main() {
	pool, err := pgxpool.New(context.Background(), "postgres://postgres:123@localhost:5432/bd")
	if err != nil {
		panic(err)
	}
	defer pool.Close()

	// Get the current working directory
	// dir, err := os.Getwd()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // Read all files and folders in the current directory
	// files, err := ioutil.ReadDir(dir)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	dir := "/home/krone/Documents"

	// Read all files and folders in the specified directory
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	// Loop through each item and insert its name into the database
	// for _, file := range files {
	// 	fmt.Println("Inserting:", file.Name())
	// 	nextID := getNextID(pool)
	// 	insertIntoDatabase(pool, nextID, file.Name())

	for _, file := range files {
		if file.IsDir() { // Check if the item is a directory
			fmt.Println("Inserting:", file.Name())
			nextID := getNextID(pool)  // Retrieve the next ID
			insertIntoDatabase(pool, nextID, file.Name())
		}
	}
	

	fmt.Println("Done!")
}


