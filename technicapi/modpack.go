package technicapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"

	"github.com/Zaprit/Reflow/config"
	"github.com/Zaprit/Reflow/database"
	"github.com/Zaprit/Reflow/models"
)

// GetModpacks gets the modpack list from the database and returns it as JSON
func GetModpacks(w http.ResponseWriter, _ *http.Request) {
	var out models.ModpackList

	var modpacks []models.ListModpack

	out.Modpacks = make(map[string]string)
	out.MirrorURL = config.RepoURL

	database.GetDBInstance().Model(&models.Modpack{}).Find(&modpacks)

	for i := range modpacks {
		out.Modpacks[modpacks[i].Name] = modpacks[i].DisplayName
	}

	outJSON, _ := json.Marshal(out)
	_, err := w.Write(outJSON)

	if err != nil {
		fmt.Printf("Error in GetMods: %s", err.Error())
	}
}

// GetModpack gets a specific modpack and its builds and returns them in a JSON document
func GetModpack(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	var modpack models.Modpack

	var builds []models.ModpackBuild

	result := database.GetDBInstance().First(&modpack).Where("name = ?", vars["slug"])

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		_, err := w.Write(models.APIErrorJSON("Modpack does not exist/Build does not exist"))
		if err != nil {
			panic(err.Error())
		}
		return
	}

	database.GetDBInstance().Find(&builds).Where("name = ?", vars["slug"])

	for i := range builds {
		modpack.Builds = append(modpack.Builds, builds[i].Version)
	}

	out, err := json.Marshal(modpack)

	if err != nil {
		panic(err.Error())
	}

	_, err = w.Write(out)

	if err != nil {
		panic(err.Error())
	}
}

func GetBuild(w http.ResponseWriter, req *http.Request) {

}
