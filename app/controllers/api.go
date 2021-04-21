package controllers

import (
	"github.com/revel/revel"
)

type APIInfo struct {
	Name    string ` json:"api" `
	Version string ` json:"version" `
	Stream  string ` json:"stream" `
}

type Mod struct {
	ID          int      ` json:"id" `
	Name        string   ` json:"name" `
	DisplayName string   ` json:"pretty_name" `
	Author      string   ` json:"author" `
	Description string   ` json:"description" `
	Link        string   ` json:"link" `
	Versions    []string ` json:"versions" `
}

type ModVersion struct {
	ID       int    ` json:"id" `
	MD5      string ` json:"md5" `
	Filesize string ` json:"filesize" `
	URL      string ` json:"url" `
}

var defaultInfo = APIInfo{Name: "Reflow", Version: "v0.1", Stream: "DEV"}

type TechnicAPIController struct {
	*revel.Controller
	MyMappedData map[string]interface{}
}

func (c TechnicAPIController) ApiRoot() revel.Result {
	return c.RenderJSON(defaultInfo)
}

func (c TechnicAPIController) GetMods() revel.Result {
	if c.Params.Route == nil {
		return c.RenderText("you probably went to /api/mods")
	} else if c.Params.Route.Get("version") == "" {
		return c.RenderJSON(Mod{ID: 67, Name: "forestry", DisplayName: "Forestry", Author: "SirSengir", Description: "Forestry deals with farming, renewable energy production as well as the breeding of trees, bees and butterflies in Minecraft.", Link: `https://forestryforminecraft.info`, Versions: []string{"1.12.2-5.8.2.387"}})
	} else {
		return c.RenderText("you probably went to /api/mods/something/version where something is " + c.Params.Route.Get("slug") + " and version is " + c.Params.Route.Get("version"))
	}
}
