package db

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // PostgreSQL driver
)

var (
	// DbHost ...
	DbHost = os.Getenv("PSQL_HOST")
	// DbName ...
	DbName = os.Getenv("PSQL_DB")
	// DbUser ...
	DbUser = os.Getenv("PSQL_USER")
	// DbPass ...
	DbPass = os.Getenv("PSQL_PASSWORD")
)

// Init connects to the PostgreSQL database.
func Init() *gorm.DB {
	dbinfo := fmt.Sprintf("host=%s dbname=%s user=%s password=%s", DbHost, DbName, DbUser, DbPass)
	db, err := gorm.Open("postgres", dbinfo)
	checkFatalError(err)
	return db
}

// UpdateSchema creates tables, missing columns and missing indexes, but won't change existing column's type or delete unused columns.
func UpdateSchema() {
	db := Init()
	defer db.Close()
	db.AutoMigrate(&Contents{}, &Metas{}, &Relationships{}, &Users{})
}

func checkFatalError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
