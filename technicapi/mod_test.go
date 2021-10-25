package technicapi

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
	database.CreateDBInstance(&database.DBConfig{
		Driver: "sqlite",
		DBName: "file::memory:?cache=shared",
	})

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

func TestGetModVersion(t *testing.T) {

}

func TestGetMods(t *testing.T) {
	database.CreateDBInstance(&database.DBConfig{
		Driver: "sqlite",
		DBName: "file::memory:?cache=shared",
	})

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

	var modlist models.ModList

	er2 := json.Unmarshal(body, &modlist)
	if er2 != nil {
		t.Fatal(er2.Error())
	}

	mods := map[string]string{
		dbMod.Name:  dbMod.DisplayName,
		dbMod2.Name: dbMod2.DisplayName,
		dbMod3.Name: dbMod3.DisplayName,
	}

	for name, prettyName := range modlist.Mods {
		if mods[name] != prettyName {
			t.Errorf("While checking modlist response unexpected mod, expected {%s:%s}, got {%s:%s}",
				name, mods[name], name, prettyName)
		}
	}
}
