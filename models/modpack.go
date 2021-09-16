package models

// ModpackList is a list of slug/name pairs and the mirror URL, don't know why it's done like this but eh what can you do
type ModpackList struct {
	Modpacks  map[string]string `json:"modpacks" `
	MirrorURL string            `json:"mirror_url"`
}

type Modpack struct {
	Name          string   `json:"name"`
	DisplayName   string   `json:"display_name"`
	URL           string   `json:"url"`
	Icon          string   `json:"icon"`
	IconMd5       string   `json:"icon_md5"`
	Logo          string   `json:"logo"`
	LogoMd5       string   `json:"logo_md5"`
	Background    string   `json:"background"`
	BackgroundMd5 string   `json:"background_md5"`
	Recommended   string   `json:"recommended"`
	Latest        string   `json:"latest"`
	Builds        []string `json:"builds"`
}
