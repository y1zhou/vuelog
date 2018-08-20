package db

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-contrib/sessions"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // PostgreSQL driver
)

var (
	dbHost    = os.Getenv("PSQL_HOST")
	dbName    = os.Getenv("PSQL_DB")
	dbUser    = os.Getenv("PSQL_USER")
	dbPass    = os.Getenv("PSQL_PASSWORD")
	redisPass = os.Getenv("REDIS_PASSWORD")
)

// InitPostgres connects to the PostgreSQL database.
func InitPostgres() *gorm.DB {
	dbinfo := fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=disable", dbHost, dbName, dbUser, dbPass)
	db, err := gorm.Open("postgres", dbinfo)
	checkFatalError(err)
	return db
}

// InitRedis ...
func InitRedis() {
	store, _ := redis.NewStore(10, "tcp", "localhost:6379", REDIS_PASSWORD, []byte("qzbjyzbj"))
	store.Options = &sessions.Options{
		MaxAge: 86400,
		Secure: true,
		HttpOnly: true
	}
	return store
}

// UpdateSchema creates tables, missing columns and missing indexes, but won't change existing column's type or delete unused columns.
func UpdateSchema() {
	db := InitPostgres()
	defer db.Close()
	db.AutoMigrate(&Contents{}, &Metas{}, &Relationships{}, &Users{})
}

func checkFatalError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
