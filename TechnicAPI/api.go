package TechnicAPI

import (
	"encoding/json"
	"github.com/Zaprit/Reflow/Database"
	"github.com/Zaprit/Reflow/Models"
	"net/http"
)


func ApiRoot(w http.ResponseWriter, req *http.Request) {
	out, _ := json.Marshal(Models.DefaultInfo)
	w.Write(out)
}

func GetMods(w http.ResponseWriter, req *http.Request){
		var modmap = make(map[string]string)
		var mods []Models.Mod

		Database.GetDBInstance().Instance.Find(&mods)
		for _, m := range mods {
			modmap[m.Name] = m.DisplayName
		}
		out, _ := json.Marshal(Models.ModList{modmap})
		w.Write(out)
	//else if c.Params.Route.Get("version") == "" {
	//
	//	var mod Models.Mod
	//	var versions []Models.Modversion
	//	Database.GetDBInstance().Instance.First(&mod, "name = ?", c.Params.Route.Get("slug"))
	//	Database.GetDBInstance().Instance.Where("mod_id = ?", mod.ID).Find(&versions)
	//
	//	for _, s := range versions {
	//		mod.Versions = append(mod.Versions, s.Version)
	//	}
	//	return c.RenderJSON(mod)
	//
	//} else {
	//	// specific mod version TODO: Implement this
	//
	//	return c.RenderText("NOT YET IMPLEMENTED")
	//}
}