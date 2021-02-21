package db

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Initialize the sqllite database
func Initialize(model ...interface{}) {
	// github.com/mattn/go-sqlite3

	if _, err := os.Stat("sqlite-database.db"); os.IsNotExist(err) {
		// your file does not exist
		file, err := os.Create("sqlite-database.db") // Create SQLite file
		if err != nil {
			log.Fatal(err.Error())
		}
		file.Close()

		db, err := GetDBInstance()
		if err != nil {
			log.Fatalln("Coudln't get database instance")
		}

		for _, m := range model {
			db.AutoMigrate(m)
		}
	}

}

// GetDBInstance will return the db instance
func GetDBInstance() (*gorm.DB, error) {
	return gorm.Open(sqlite.Open("sqlite-database.db"), &gorm.Config{})
}
