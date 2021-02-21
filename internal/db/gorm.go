package db

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Initialize the sqllite database
func Initialize() {
	// github.com/mattn/go-sqlite3
	file, err := os.Create("sqlite-database.db") // Create SQLite file
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()
}

// GetDBInstance will return the db instance
func GetDBInstance() (*gorm.DB, error) {
	return gorm.Open(sqlite.Open("sqlite-database.db"), &gorm.Config{})
}
