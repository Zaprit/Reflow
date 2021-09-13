// Package models contains the structs for the various things
package models

import (
	"time"
)

// APIInfo is the information required to make the technic client believe that we are definitely talking to a solder server.
// In reality the client isn't very picky.
type APIInfo struct {
	Name    string ` json:"api" `
	Version string ` json:"version" `
	Stream  string ` json:"stream" `
}

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
	ID        int
	ModID     int ` json:"-" `
	Version   string
	MD5       string    ` json:"md5" `
	CreatedAt time.Time ` json:"-" `
	UpdatedAt time.Time ` json:"-" `
	Filesize  string    ` json:"filesize" `
	URL       string    ` json:"url" gorm:"-"`
}

// ModList is for the most part a bodge to get the JSON document correct for /api/mod
type ModList struct {
	Mods map[string]string `json:"mods" `
}

// APIError is the representation of an error sent by solder as a struct
// An Example of an error:
//  {"error":"Mod does not exist"}
type APIError struct {
	Message	string ` json:"error" `
}

// DefaultInfo is the default APIInfo for reflow.
var DefaultInfo = APIInfo{Name: "Reflow", Version: "v0.1", Stream: "DEV"}
