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

}
