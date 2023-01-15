package main

import (
	"database/sql"
)

// Create tables
func initDB() error {
	db, err := connectToDB()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE movies (id SERIAL PRIMARY KEY, title TEXT, year INTEGER, genre TEXT, rating REAL)`)
	if err != nil {
		return err
	}

	return nil
}

// Database connection
func connectToDB() (*sql.DB, error) {

	connStr := "user=gerard password=12345 dbname=postgres sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}