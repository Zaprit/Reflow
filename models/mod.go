package models

import (
	"time"
)

// A Mod is a singular mod roughly compliant with what comes out the database and what gets turned into JSON
type Mod struct {
	ID          uint64    ` json:"id"`
	Name        string    ` json:"name" `
	Description string    ` json:"description" `
	Author      string    ` json:"author" `
	Link        string    ` json:"link" `
	CreatedAt   time.Time ` json:"-" `
	UpdatedAt   time.Time ` json:"-" `
	Versions    []string  ` json:"versions" gorm:"-"`
	DisplayName string    ` json:"pretty_name" gorm:"column:pretty_name" `
}

// A ModVersion is a version of a mod with all the things required for JSON
type ModVersion struct {
	ID        int       ` json:"id" `
	ModID     int       ` json:"-" `
	Version   string    ` json:"-" `
	MD5       string    ` json:"md5" `
	CreatedAt time.Time ` json:"-" `
	UpdatedAt time.Time ` json:"-" `
	Filesize  string    ` json:"filesize" `
	URL       string    ` json:"url" gorm:"url"`
}

// ModList is for the most part a bodge to get the JSON document correct for /api/mod
type ModList struct {
	Mods map[string]string `json:"mods" `
}
