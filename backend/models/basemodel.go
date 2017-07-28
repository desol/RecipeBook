package models

import (
	"log"

	"github.com/asdine/storm"
)

var db *storm.DB // Holds the connection to the storm database

// InitDB : Opens the connection to the database
func InitDB(path string) error {
	var err error
	db, err = storm.Open(path)

	return err
}

// CloseDB : Closes the connection to the database
func CloseDB() {
	err := db.Close()
	if err != nil {
		log.Fatalf("Failed to close db. Message: " + err.Error())
	}
}
