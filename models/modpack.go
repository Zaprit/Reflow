package models

import "gopkg.in/guregu/null.v4"

// ModpackList is a list of slug/name pairs and the mirror URL, don't know why it's done like this but eh what can you do
type ModpackList struct {
	Modpacks  map[string]string `json:"modpacks" `
	MirrorURL string            `json:"mirror_url"`
}

// Modpack is a struct representation of a technic modpack
type Modpack struct {
	DBStructTemplate
	Name          string   ` json:"name" gorm:"column:slug" `
	DisplayName   string   ` json:"display_name" gorm:"column:name" `
	URL           string   ` json:"url" `
	Icon          bool     ` json:"-" `
	IconURL       string   ` json:"icon" `
	IconMD5       string   ` json:"icon_md5" `
	Logo          bool     ` json:"-"`
	LogoURL       string   ` json:"logo" `
	LogoMD5       string   ` json:"logo_md5" `
	Background    bool     ` json:"-" `
	BackgroundURL string   ` json:"background" `
	BackgroundMD5 string   ` json:"background_md5" `
	Recommended   string   ` json:"recommended" `
	Latest        string   ` json:"latest" `
	Builds        []string ` json:"builds" gorm:"-" `
	Order         int32    ` json:"-" `
	Hidden        bool     ` json:"-" `
	Private       bool     ` json:"-" `
}

// ListModpack is a stripped down modpack to save bandwidth while querying.
type ListModpack struct {
	Name        string ` json:"name" gorm:"column:slug" `
	DisplayName string ` json:"display_name" gorm:"column:name" `
}

// ModpackBuild is a struct representation of a modpack build.
type ModpackBuild struct {
	DBStructTemplateID
	ModpackID    int32        ` json:"-" `
	Version      string       ` json:"-" `
	Minecraft    string       ` json:"minecraft" `
	MinecraftMD5 string       ` json:"minecraft_md5"`
	Java         string       ` json:"java" gorm:"column:min_java" `
	Memory       int32        ` json:"memory" gorm:"column:min_memory" `
	Forge        null.String  ` json:"forge" `
	Mods         []ModpackMod ` json:"mods" gorm:"-" `
	IsPublished  bool         ` json:"-"`
	Private      bool         ` json:"-" `
}

// TableName is the Tabler interface that provides a table name for GORM
func (ModpackBuild) TableName() string {
	return "builds"
}
