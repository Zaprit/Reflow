// Package technicapi is the package that defines the solder compatible api
package technicapi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/Zaprit/Reflow/config"
	"github.com/Zaprit/Reflow/database"
	"github.com/Zaprit/Reflow/models"
)

// GetMods gets the list of mods from the database and displays it in a JSON document with the following format
//  {"mod-slug":"pretty-name",...}
// Where mod-slug is the internal name used by regular solder and pretty-name is the display name
func GetMods(w http.ResponseWriter, _ *http.Request) {
	var modMap = make(map[string]string)

	var mods []models.Mod

	database.GetDBInstance().Instance.Find(&mods)

	for i := range mods {
		modMap[mods[i].Name] = mods[i].DisplayName
	}

	out, _ := json.Marshal(models.ModList{Mods: modMap})
	_, err := w.Write(out)

	if err != nil {
		fmt.Printf("Error in GetMods: %s", err.Error())
	}
}

// GetMod gets a specific mod from the database and returns it as a JSON document with the following format
//  {
//   "id":123,
//   "name":"test-mod",
//   "description":"This is a description of a test mod","author":"Test Author",
//   "link":"https://example.com",
//   "versions":["1.0"],
//   "pretty_name":"Test Mod"
//  }
// If the mod can't be found it returns this error
func GetMod(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	var mod models.Mod

	var versions []models.ModVersion

	database.GetDBInstance().Instance.First(&mod, "name = ?", vars["slug"])
	database.GetDBInstance().Instance.Table("modversions").Where("mod_id = ?", mod.ID).Find(&versions)

	for s := range versions {
		mod.Versions = append(mod.Versions, versions[s].Version)
	}

	var output []byte

	var err error

	if mod.Name == "" {
		output, err = json.Marshal(models.APIError{Message: "Mod does not exist"})
	} else {
		output, err = json.Marshal(mod)
	}

	if err != nil {
		fmt.Println(err.Error())
	} else {
		_, err = w.Write(output)

		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

// GetModVersion is the API endpoint that returns a specific version of a mod in the following JSON format
// {
//  "id":1,
//  "md5":"949b3066566657167bc3da57fd1b0a83",
//  "filesize":462,
//  "url":"http:\/\/127.0.0.1:8080\/mods\/test-mod\/test-mod-1.0.zip"
// }
// the strange escaping on the url is a holdover from solder
func GetModVersion(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	var mod models.Mod

	var version models.ModVersion

	database.GetDBInstance().Instance.First(&mod, "name = ?", vars["slug"])
	database.GetDBInstance().Instance.Table("modversions").Where("mod_id = ? AND version = ?", mod.ID, vars["version"]).First(&version)

	if version.URL == "" {
		version.URL = fmt.Sprintf("%s/%s/%s", config.RepoURL, mod.Name, version.Version)
	}

	out, err := json.Marshal(version)

	if err != nil {
		w.WriteHeader(500)
		_, err = w.Write([]byte("Internal Server Error"))

		if err != nil {
			panic(err.Error())
		}
	}

	_, err = w.Write(out)
	if err != nil {
		panic(err.Error())
	}
}
