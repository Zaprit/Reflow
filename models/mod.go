package models

// A Mod is a singular mod roughly compliant with what comes out the database and what gets turned into JSON
type Mod struct {
	DBStructTemplate
	Name        string   ` json:"name" `
	Description string   ` json:"description" `
	Author      string   ` json:"author" `
	Link        string   ` json:"link" `
	Versions    []string ` json:"versions" gorm:"-"`
	DisplayName string   ` json:"pretty_name" gorm:"column:pretty_name" `
}

// A ModVersion is a version of a mod with all the things required for JSON
type ModVersion struct {
	DBStructTemplate
	ModID    int32  ` json:"-" `
	Version  string ` json:"-" `
	MD5      string ` json:"md5" `
	Filesize int32  ` json:"filesize" `
	URL      string ` json:"url" gorm:"url"`
}

// TableName is the tabler interface function for GORM
func (ModVersion) TableName() string {
	return "modversions"
}

// ModpackMod is the mod version type returned by /api/modpack/<slug>/<build>
type ModpackMod struct {
	ID       int32  `json:"id"`
	Name     string `json:"name"`
	Version  string `json:"version"`
	MD5      string `json:"md5"`
	Filesize int32  `json:"filesize"`
	URL      string `json:"url"`
}

// ModpackModFormat converts a mod and a modVersion into a modpack build compatible doodad
func ModpackModFormat(mod *Mod, modVersion *ModVersion) ModpackMod{
	return ModpackMod{
		ID:       mod.ID,
		Name:     mod.Name,
		Version:  modVersion.Version,
		MD5:      modVersion.MD5,
		Filesize: modVersion.Filesize,
		URL:      modVersion.URL,
	}
}

// BuildModversion is the DB map between modversions and modpack builds
type BuildModversion struct {
	DBStructTemplate
	ModVersionID uint ` gorm:"column:modversion_id" `
	BuildID      uint
}

// TableName is the tabler interface function for GORM
func (BuildModversion) TableName() string {
	return "build_modversion"
}

// ModList is for the most part a bodge to get the JSON document correct for /api/mod
type ModList struct {
	Mods map[string]string `json:"mods" `
}
