package solderapi

import (
	"encoding/json"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/gorilla/mux"

	"github.com/Zaprit/Reflow/database"
	"github.com/Zaprit/Reflow/internal"
	"github.com/Zaprit/Reflow/models"
)

func TestGetMod(t *testing.T) {

	dbMod := models.Mod{
		DBStructTemplate: models.DBStructTemplate{CreatedAt: time.Now()},
		Name:             "test-mod",
		Description:      "Test Mod Description",
		Author:           "Test Author",
		Link:             "https://example.com",
		DisplayName:      "Test Mod",
	}
	database.GetDBInstance().Create(&dbMod)

	var modFromDB models.Mod

	database.GetDBInstance().Limit(1).Where("name = ?", "test-mod").Take(&modFromDB)

	dbModVersion := models.ModVersion{
		ModID:    modFromDB.ID,
		Version:  "1.0",
		MD5:      "notARealMD5hash",
		Filesize: 12345677,
		URL:      "https://example.com",
	}

	database.GetDBInstance().Create(&dbModVersion)

	r := mux.NewRouter()
	r.HandleFunc("/api/mod/{slug}", GetMod)

	ts := httptest.NewServer(r)

	body, err := internal.TestClient(ts.URL + "/api/mod/test-mod")

	if err != nil {
		t.Fatal(err.Error())
	}

	var mod models.Mod

	er2 := json.Unmarshal(body, &mod)
	if er2 != nil {
		t.Fatal(er2.Error())
	}

	if reflect.DeepEqual(&dbMod, &mod) {
		t.Fatalf("Mod Mismatch, Expected: %v, Received: %v\nIn Data: %s", dbMod, mod, body)
	}
}

// TestGetModVersions tests the GetModVersions function
func TestGetModVersion(t *testing.T) {
	// Create a new mod and a mod version
	dbMod := models.Mod{
		DBStructTemplate: models.DBStructTemplate{CreatedAt: time.Now()},
		Name:             "test-mod",
		Description:      "Test Mod Description",
	}
	database.GetDBInstance().Create(&dbMod)
	dbModVersion := models.ModVersion{
		ModID:    dbMod.ID,
		Version:  "1.0",
		MD5:      "notARealMD5hash",
		Filesize: 12345677,
		URL:      "https://example.com",
	}
	database.GetDBInstance().Create(&dbModVersion)

	// Create a new router
	r := mux.NewRouter()
	r.HandleFunc("/api/mod/{slug}/version/{version}", GetModVersion)

	// Create a new test server
	ts := httptest.NewServer(r)

	// Create a new client
	resp, err := internal.TestClient(ts.URL + "/api/mod/test-mod/version/1.0")
	if err != nil {
		t.Fatal(err.Error())
	}
	// Unmarshal the response
	var modVersion models.ModVersion
	er2 := json.Unmarshal(resp, &modVersion)
	if er2 != nil {
		t.Fatal(er2.Error())
	}

	if reflect.DeepEqual(&dbModVersion, &modVersion) {
		t.Fatalf("Mod Version Mismatch, Expected: %v, Received: %v\nIn Data: %s", dbModVersion, modVersion, resp)
	}

}

func TestGetMods(t *testing.T) {

	dbMod := models.Mod{
		DBStructTemplate: models.DBStructTemplate{CreatedAt: time.Now()},
		Name:             "test-mod",
		Description:      "Test Mod Description",
		Author:           "Test Author",
		Link:             "https://example.com",
		DisplayName:      "Test Mod",
	}
	dbMod2 := models.Mod{
		DBStructTemplate: models.DBStructTemplate{CreatedAt: time.Now()},
		Name:             "test-mod-two",
		Description:      "Second Test Mod Description",
		Author:           "Second Test Author",
		Link:             "https://example.org",
		DisplayName:      "Second Test Mod",
	}
	dbMod3 := models.Mod{
		DBStructTemplate: models.DBStructTemplate{CreatedAt: time.Now()},
		Name:             "test-mod-three",
		Description:      "Third Test Mod Description",
		Author:           "Third Test Author",
		Link:             "https://example.net",
		DisplayName:      "Third Test Mod",
	}

	database.GetDBInstance().Create(&dbMod)
	database.GetDBInstance().Create(&dbMod2)
	database.GetDBInstance().Create(&dbMod3)

	r := mux.NewRouter()
	r.HandleFunc("/api/mod", GetMod)

	ts := httptest.NewServer(r)

	body, err := internal.TestClient(ts.URL + "/api/mod")

	if err != nil {
		t.Fatal(err.Error())
	}

	var modList models.ModList

	er2 := json.Unmarshal(body, &modList)
	if er2 != nil {
		t.Fatal(er2.Error())
	}

	mods := map[string]string{
		dbMod.Name:  dbMod.DisplayName,
		dbMod2.Name: dbMod2.DisplayName,
		dbMod3.Name: dbMod3.DisplayName,
	}

	for name, prettyName := range modList.Mods {
		if mods[name] != prettyName {
			t.Errorf("While checking modList response unexpected mod, expected {%s:%s}, got {%s:%s}",
				name, mods[name], name, prettyName)
		}
	}
}
