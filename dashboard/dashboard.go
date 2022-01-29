// Package dashboard contains the web ui for reflow, for the static files see Package static
package dashboard

import (
	"html/template"
	"net/http"

	"github.com/Zaprit/Reflow/database"
	"github.com/Zaprit/Reflow/models"

	"github.com/Zaprit/Reflow/static"
)

// WebDashboardRoot is the page handler for the mod list
func WebDashboardRoot(w http.ResponseWriter, _ *http.Request) {
	_, err := w.Write([]byte("NOT YET IMPLEMENTED"))
	if err != nil {
		panic(err.Error())
	}
}

// WebDashboardModList is the page handler for the mod list
func WebDashboardModList(w http.ResponseWriter, _ *http.Request) {
	tmpl, err := template.ParseFS(static.WebFS, "static/mod.html")
	if err != nil {
		panic(err.Error())
	}
	var pageMods []CompiledMod //nolint:prealloc The array will contain an unknown number of elements

	var mods []models.Mod

	database.GetDBInstance().Find(&mods)

	for i := range mods {
		var compMod CompiledMod
		compMod.Mod = mods[i]
		database.GetDBInstance().Find(&compMod.ModVersions, "mod_id = ?", mods[i].ID)
		pageMods = append(pageMods, compMod)
	}

	pageData := ModPageData{
		ModCount: len(pageMods),
		Mods:     pageMods,
	}

	err = tmpl.Execute(w, pageData)
	if err != nil {
		panic(err.Error())
	}
}
