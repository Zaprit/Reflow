package technicapi

import (
	"encoding/json"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	"github.com/Zaprit/Reflow/internal"

	"github.com/gorilla/mux"

	"github.com/Zaprit/Reflow/config"
	"github.com/Zaprit/Reflow/database"
	"github.com/Zaprit/Reflow/models"
)

func TestAPIRoot(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/api", APIRoot)
	ts := httptest.NewServer(r)

	resp, err := internal.TestClient(ts.URL + "/api")
	if err != nil {
		t.Errorf("Expected nil, Received %s", err.Error())
	}

	ts.Close()

	var info models.APIInfo

	er2 := json.Unmarshal(resp, &info)
	if er2 != nil {
		t.Errorf("Error while Unmarshaling JSON, %s", er2.Error())
	}

	if !reflect.DeepEqual(info, config.DefaultInfo) {
		t.Errorf("API Info Mismatch Expected: %v, Received: %v", config.DefaultInfo, info)
	}
}

// TestMain Bootstraps the tests for technicapi by creating a database instance and starting a http server
func TestMain(m *testing.M) {
	database.CreateDBInstance(&database.DBConfig{
		Driver: "sqlite",
		DBName: "file::memory:?cache=shared",
	})

	InitDB()

	code := m.Run()

	os.Exit(code)
}
