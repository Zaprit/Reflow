package models

import "time"

type APIInfo struct {
	Name    string ` json:"api" `
	Version string ` json:"version" `
	Stream  string ` json:"stream" `
}

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

type Modversion struct {
	ID       int    ` json:"id" `
	MD5      string ` json:"md5" `
	Filesize string ` json:"filesize" `
	URL      string ` json:"url" `
}

var Mods []Mod
var DefaultInfo = APIInfo{Name: "Reflow", Version: "v0.1", Stream: "DEV"}
