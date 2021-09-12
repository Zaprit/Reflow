package TechnicAPI

import (
	"encoding/json"
	"fmt"
	"github.com/Zaprit/Reflow/Database"
	"github.com/Zaprit/Reflow/Models"
	"net/http"
)


func ApiRoot(w http.ResponseWriter, _ *http.Request) {
	out, _ := json.Marshal(Models.DefaultInfo)
	_, err := w.Write(out)
	if err != nil {
		fmt.Printf("Error In /ApiRoot: %s", err.Error())
	}
}

func GetMods(w http.ResponseWriter, _ *http.Request){
	var modMap = make(map[string]string)
	var mods []Models.Mod

	Database.GetDBInstance().Instance.Find(&mods)
	for _, m := range mods {
		modMap[m.Name] = m.DisplayName
	}
	out, _ := json.Marshal(Models.ModList{Mods: modMap})
	_, err := w.Write(out)
	if err != nil {
		fmt.Printf("Error in GetMods: %s", err.Error())
	}
	// TODO: Rewrite this into separate functions
	//else if c.Params.Route.Get("version") == "" {
	//
	//	var mod Models.Mod
	//	var versions []Models.ModVersion
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