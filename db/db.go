package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	// DbUser ...
	DbUser = os.Getenv("PSQL_USER")
	// DbName ...
	DbName = os.Getenv("PSQL_DB")
	// DbPass ...
	DbPass = os.Getenv("PSQL_PASSWORD")
)

// Contents table with posts, drafts, links, ...
type Contents struct {
	CID          int64     `gorm:"PRIMARY_KEY;type:serial;NOT NULL;"`
	Title        string    `gorm:"type:varchar(200);DEFAULT:NULL"`
	Slug         int64     `gorm:"type:varchar(200);DEFAULT:NULL;INDEX:contents_slug"`
	CreatedAt    time.Time `gorm:"type:timestamp with time zone;INDEX:contents_created"`
	UpdatedAt    time.Time `gorm:"type:timestamp with time zone"`
	DeletedAt    time.Time `gorm:"type:timestamp with time zone"`
	Text         string    `gorm:"type:text"`
	Order        int64     `gorm:"DEFAULT:0"`
	AuthorID     int64     `gorm:"DEFAULT:0"`
	Template     string    `gorm:"type:varchar(32);DEFAULT:NULL"`
	Type         string    `gorm:"type:varchar(16);DEFAULT:'post'"`
	Status       string    `gorm:"type:varchar(16);DEFAULT:'publish'"`
	AllowComment bool      `gorm:"type:boolean;DEFAULT:true;NOT NULL"`
}

// Metas table with short descriptions of the contents
type Metas struct {
	MID         int64  `gorm:"PRIMARY_KEY;type:serial;NOT NULL"`
	Name        string `gorm:"type:varchar(150);DEFAULT:NULL"`
	Slug        string `gorm:"type:varchar(150);DEFAULT:NULL;UNIQUE_INDEX:metas_slug"`
	Type        string `gorm:"type:varchar(32);NOT NULL"`
	Description string `gorm:"type:varchar(150);DEFAULT:NULL"`
	Count       int64  `gorm:"DEFAULT:0"`
	Order       int64  `gorm:"DEFAULT:0"`
}

// Relationships of contents and metas
type Relationships struct {
	CID int64 `gorm:"PRIMARY_KEY;type:serial;NOT NULL"`
	MID int64 `gorm:"PRIMARY_KEY;type:serial;NOT NULL"`
}

//Users ...
type Users struct {
	UID       int64     `gorm:"PRIMARY_KEY;type:serial;NOT NULL"`
	Name      string    `gorm:"type:varchar(32);DEFAULT:NULL;UNIQUE_INDEX:users_name"`
	Password  string    `gorm:"type:varchar(64);DEFAULT:NULL"`
	CreatedAt time.Time `gorm:"type:timestamp with time zone;INDEX:users_created"`
	DeletedAt time.Time `gorm:"type:timestamp with time zone"`
	Logged    time.Time `gorm:"type:timestamp with time zone"`
	Group     string    `gorm:"type:varchar(16);DEFAULT:'vistor'"`
}

// Init connects to the PostgreSQL database.
func Init() {
	dbinfo := fmt.Sprintf("user=%s dbname=%s password=%s", DbUser, DbName, DbPass)
	db, err := gorm.Open("postgres", dbinfo)
	db.AutoMigrate(&Contents{}, &Metas{}, &Relationships{}, &Users{})
	checkErr(err)
	defer db.Close()
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
