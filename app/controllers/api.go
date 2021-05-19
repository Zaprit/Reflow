package controllers

import (
	"reflow/app"

	"github.com/revel/revel"
)

type TechnicAPIController struct {
	App
	//MyMappedData map[string]interface{}
}

func (c TechnicAPIController) ApiRoot() revel.Result {
	return c.RenderJSON(app.DefaultInfo)
}

func (c TechnicAPIController) GetMods() revel.Result {
	if c.Params.Route == nil {
		return c.RenderText("you probably went to /api/mods")
	} else if c.Params.Route.Get("version") == "" {
		return c.RenderJSON(c.Txn.SelectOne(&app.Mod{}, c.Db.SqlStatementBuilder.Select("*").From("Mods").Where("name=?", c.Params.Route.Get("slug"))))
		//return c.RenderJSON(app.Mod{ID: 67, Name: "forestry", DisplayName: "Forestry", Author: "SirSengir", Description: "Forestry deals with farming, renewable energy production as well as the breeding of trees, bees and butterflies in Minecraft.", Link: `https://forestryforminecraft.info`, Versions: []string{"1.12.2-5.8.2.387"}})
	} else {
		return c.RenderText("you probably went to /api/mods/something/version where something is " + c.Params.Route.Get("slug") + " and version is " + c.Params.Route.Get("version"))
	}
}

func (c TechnicAPIController) AddMod() revel.Result {
	//c.DB.Create(app.Mod{Name: "forestry", DisplayName: "Forestry", Author: "SirSengir", Description: "Forestry deals with farming, renewable energy production as well as the breeding of trees, bees and butterflies in Minecraft.", Link: "https://forestryforminecraft.info", Versions: []string{"1.12.2-5.8.2.387"}})
	//c.DB.Create(app.Mod{Name: c.Params.Form.Get("name"), DisplayName: c.Params.Form.Get("displayname"), Author: c.Params.Form.Get("author"), Description: c.Params.Form.Get("description"), Link: c.Params.Form.Get("link")})
	return c.Render()
}
