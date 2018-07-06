package db

import "time"

type (
	// Contents table with posts, drafts, links, ...
	Contents struct {
		CID          int64     `gorm:"PRIMARY_KEY;type:serial;NOT NULL;"`
		Title        string    `gorm:"type:varchar(200);DEFAULT:NULL"`
		Slug         int64     `gorm:"type:varchar(200);DEFAULT:NULL;INDEX:contents_slug"`
		CreatedAt    time.Time `gorm:"type:timestamp with time zone;INDEX:contents_created"`
		UpdatedAt    time.Time `gorm:"type:timestamp with time zone"`
		DeletedAt    time.Time `gorm:"type:timestamp with time zone;DEFAULT:NULL"`
		Text         string    `gorm:"type:text"`
		Order        int64     `gorm:"DEFAULT:0"`
		AuthorID     int64     `gorm:"DEFAULT:0"`
		Template     string    `gorm:"type:varchar(32);DEFAULT:NULL"`
		Type         string    `gorm:"type:varchar(16);DEFAULT:'post'"`
		Status       string    `gorm:"type:varchar(16);DEFAULT:'publish'"`
		AllowComment bool      `gorm:"type:boolean;DEFAULT:true;NOT NULL"`
	}

	// Metas table with metadata of the contents
	Metas struct {
		MID         int64  `gorm:"PRIMARY_KEY;type:serial;NOT NULL"`
		Name        string `gorm:"type:varchar(150);DEFAULT:NULL"`
		Slug        string `gorm:"type:varchar(150);DEFAULT:NULL;UNIQUE_INDEX:metas_slug"`
		Type        string `gorm:"type:varchar(32);NOT NULL"`
		Description string `gorm:"type:varchar(150);DEFAULT:NULL"`
		Count       int64  `gorm:"DEFAULT:0"`
		Order       int64  `gorm:"DEFAULT:0"`
	}

	// Relationships of contents and metas
	Relationships struct {
		CID int64 `gorm:"PRIMARY_KEY;type:serial;NOT NULL"`
		MID int64 `gorm:"PRIMARY_KEY;type:serial;NOT NULL"`
	}

	//Users ...
	Users struct {
		UID       int64     `gorm:"PRIMARY_KEY;type:serial;NOT NULL"`
		Name      string    `gorm:"type:varchar(32);DEFAULT:NULL;UNIQUE;UNIQUE_INDEX:users_name"`
		Password  string    `gorm:"type:varchar(64);DEFAULT:NULL"`
		CreatedAt time.Time `gorm:"type:timestamp with time zone;INDEX:users_created"`
		UpdatedAt time.Time `gorm:"type:timestamp with time zone"`
		DeletedAt time.Time `gorm:"type:timestamp with time zone;DEFAULT:NULL"`
		Group     string    `gorm:"type:varchar(16);DEFAULT:'vistor'"`
		Email     string    `gorm:"type:varchar(200);UNIQUE;DEFAULT:NULL"`
	}
)
