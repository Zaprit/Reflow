package TechnicAPI

import (
	"encoding/json"
	"fmt"
	"github.com/Zaprit/Reflow/Database"
	"github.com/Zaprit/Reflow/Models"
	"github.com/gorilla/mux"
	"net/http"
)


func ApiRoot(w http.ResponseWriter, _ *http.Request) {
	out, _ := json.Marshal(Models.DefaultInfo)
	_, err := w.Write(out)
	if err != nil {
		fmt.Printf("Error In /ApiRoot: %s", err.Error())
	}
}

func GetMods(w http.ResponseWriter, _ *http.Request) {
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
}

func GetMod(w http.ResponseWriter, req *http.Request){
	vars := mux.Vars(req)
	var mod Models.Mod
	var versions []Models.ModVersion
	Database.GetDBInstance().Instance.First(&mod, "name = ?", vars["slug"])
	Database.GetDBInstance().Instance.Table("modversions").Where("mod_id = ?", mod.ID).Find(&versions)

	for _, s := range versions {
		mod.Versions = append(mod.Versions, s.Version)
	}
	jsonMod, err := json.Marshal(mod)
	json.
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