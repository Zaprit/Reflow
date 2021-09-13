package technicapi

import (
	"encoding/json"
	"fmt"
	"github.com/Zaprit/Reflow/database"
	"github.com/Zaprit/Reflow/models"
	"github.com/gorilla/mux"
	"net/http"
)

//ApiRoot is the root function that identifies a solder compatible api
//Stock solder returns:
//	{"api":"TechnicSolder","version":"v0.7.7","stream":"DEV"}
//I've never seen it not be DEV and as far as I can tell it doesn't matter what the api or version attributes are.
//(yes I did ask the Technic devs)
func ApiRoot(w http.ResponseWriter, _ *http.Request) {
	out, _ := json.Marshal(models.DefaultInfo)
	_, err := w.Write(out)
	if err != nil {
		fmt.Printf("Error In /ApiRoot: %s", err.Error())
	}
}

//GetMods gets the list of mods from the database and displays it in a JSON document with the following format
// {"mod-slug":"pretty-name",...}
//Where mod-slug is the internal name used by regular solder and pretty-name is the display name
func GetMods(w http.ResponseWriter, _ *http.Request) {
	var modMap = make(map[string]string)
	var mods []models.Mod

	database.GetDBInstance().Instance.Find(&mods)
	for _, m := range mods {
		modMap[m.Name] = m.DisplayName
	}
	out, _ := json.Marshal(models.ModList{Mods: modMap})
	_, err := w.Write(out)
	if err != nil {
		fmt.Printf("Error in GetMods: %s", err.Error())
	}
}

//GetMod gets a specific mod from the database and returns it as a JSON document with the following format
//{
// "id":123,
// "name":"test-mod",
// "description":"This is a description of a test mod","author":"Test Author",
// "link":"https://example.com",
// "versions":["1.0"],
// "pretty_name":"Test Mod"
//}
func GetMod(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	var mod models.Mod
	var versions []models.ModVersion
	database.GetDBInstance().Instance.First(&mod, "name = ?", vars["slug"])
	database.GetDBInstance().Instance.Table("modversions").Where("mod_id = ?", mod.ID).Find(&versions)

	for _, s := range versions {
		mod.Versions = append(mod.Versions, s.Version)
	}
	jsonMod, err := json.Marshal(mod)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		_, err = w.Write(jsonMod)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

//func GetModVersion(w http.ResponseWriter, req *http.Request){
//
//
//}
