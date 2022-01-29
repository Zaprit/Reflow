package dashboard

import "github.com/Zaprit/Reflow/models"

type ModPageData struct {
	PageTitle string
	ModCount  int
	Mods      []CompiledMod
}

type CompiledMod struct {
	Mod         models.Mod
	ModVersions []models.ModVersion
}
