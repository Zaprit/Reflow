package app

type APIInfo struct {
	Name    string ` json:"api" `
	Version string ` json:"version" `
	Stream  string ` json:"stream" `
}

type Mod struct {
	ID          int      ` json:"id" db:"id" `
	Name        string   ` json:"name" db:"name" `
	DisplayName string   ` json:"pretty_name" db:"pretty_name" `
	Author      string   ` json:"author" db:"author" `
	Description string   ` json:"description" db:"description" `
	Link        string   ` json:"link" db:"link" `
	Versions    []string ` json:"versions" db:"name" `
}

type ModVersion struct {
	ID       int    ` json:"id" `
	MD5      string ` json:"md5" `
	Filesize string ` json:"filesize" `
	URL      string ` json:"url" `
}

var Mods []Mod
var DefaultInfo = APIInfo{Name: "Reflow", Version: "v0.1", Stream: "DEV"}
