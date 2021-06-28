package controllers

import (
	"reflow/app"
	"reflow/app/models"

	"github.com/revel/revel"
)

type TechnicAPIController struct {
	App
}

func (c TechnicAPIController) ApiRoot() revel.Result {
	return c.RenderJSON(models.DefaultInfo)
}

func (c TechnicAPIController) GetMods() revel.Result {
	if c.Params.Route == nil {
		var modmap = make(map[string]string)
		var mods []models.Mod
		app.GetDBInstance().Instance.Find(&mods)
		for _, m := range mods {
			modmap[m.Name] = m.DisplayName
		}
		return c.RenderJSON(models.ModList{Mods: modmap})
	} else if c.Params.Route.Get("version") == "" {

		var mod models.Mod
		var versions []models.Modversion
		app.GetDBInstance().Instance.First(&mod, "name = ?", c.Params.Route.Get("slug"))
		app.GetDBInstance().Instance.Where("mod_id = ?", mod.ID).Find(&versions)

		for _, s := range versions {
			mod.Versions = append(mod.Versions, s.Version)
		}
		return c.RenderJSON(mod)

	} else {
		// specific mod version
		return c.RenderText("NOT YET IMPLEMENTED")
	}
}

func (c TechnicAPIController) AddMod() revel.Result {
	//c.DB.Create(app.Mod{Name: c.Params.Form.Get("name"), DisplayName: c.Params.Form.Get("displayname"), Author: c.Params.Form.Get("author"), Description: c.Params.Form.Get("description"), Link: c.Params.Form.Get("link")})
	return c.Render()
}
