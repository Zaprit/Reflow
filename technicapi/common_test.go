package technicapi

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"

	"github.com/Zaprit/Reflow/config"
	"github.com/Zaprit/Reflow/database"
	"github.com/Zaprit/Reflow/internal"
	"github.com/Zaprit/Reflow/models"
)

func TestAPIRoot(t *testing.T) {
	body, err := internal.TestClient("/api")

	if err != nil {
		t.Fatal(err.Error())
	}

	var info models.APIInfo

	er2 := json.Unmarshal(body, &info)
	if er2 != nil {
		t.Fatal(er2.Error())
	}

	if !reflect.DeepEqual(info, config.DefaultInfo) {
		t.Fatalf("API Info Mismatch Expected: %v, Received: %v", config.DefaultInfo, info)
	}
}

// TestMain Bootstraps the tests for technicapi by creating a database instance and starting a http server
func TestMain(m *testing.M) {
	database.CreateDBInstance(&database.DBConfig{
		Driver: "sqlite",
		DBName: "file::memory:?cache=shared",
	})

	InitDB()

	go StartServer("localhost:8069")

	code := m.Run()

	os.Exit(code)
}
