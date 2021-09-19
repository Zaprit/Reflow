package models

// ModpackList is a list of slug/name pairs and the mirror URL, don't know why it's done like this but eh what can you do
type ModpackList struct {
	Modpacks  map[string]string `json:"modpacks" `
	MirrorURL string            `json:"mirror_url"`
}

// Modpack is a struct representation of a technic modpack
type Modpack struct {
	DBStructTemplate
	Name          string   ` json:"name" gorm:"slug" `
	DisplayName   string   ` json:"display_name" gorm:"name" `
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
	Order         int      ` json:"-" `
	Hidden        bool     ` json:"-" `
	Private       bool     ` json:"-" `
}

// ListModpack is a stripped down modpack to save bandwidth while querying.
type ListModpack struct {
	Name        string ` json:"name" gorm:"column:slug" `
	DisplayName string ` json:"display_name" gorm:"column:name" `
}

type ModpackBuild struct {
	DBStructTemplate
	ModpackID   uint         ` json:"-" `
	Version     string       ` json:"-" `
	Minecraft   string       ` json:"minecraft" `
	Java        string       ` json:"java" gorm:"min_java" `
	Memory      int          ` json:"memory" `
	Forge       string       ` json:"forge" `
	Mods        []ModpackMod ` json:"mods" gorm:"-" `
	IsPublished bool         ` json:"-"`
	Private     bool         ` json:"-" `
}

func (ModpackBuild) TableName() string {
	return "builds"
}
