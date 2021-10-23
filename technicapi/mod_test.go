package technicapi

import (
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/Zaprit/Reflow/database"
	"github.com/Zaprit/Reflow/internal"
	"github.com/Zaprit/Reflow/models"
)

func TestGetMod(t *testing.T) {
	database.CreateDBInstance(database.DBConfig{
		Driver: "sqlite",
		DBName: "file::memory:?cache=shared",
	})

	err := database.GetDBInstance().AutoMigrate(&models.Mod{})
	if err != nil {
		t.Fatal(err.Error())
	}

	database.GetDBInstance().Create(&models.Mod{
		DBStructTemplate: models.DBStructTemplate{CreatedAt: time.Now()},
		Name:             "Test-Mod",
		Description:      "Test Mod Description",
		Author:           "Test Author",
		Link:             "http://example.com",
		Versions:         nil,
		DisplayName:      "Test Mod",
	})

	go internal.StartTestServer("/api/verify/{key}", VerifyKey)

	body, err := internal.TestClient("/api/verify/testkey2341352463")

	if err != nil {
		t.Fatal(err.Error())
	}

	var key models.APIKeyVerifyResponse

	er2 := json.Unmarshal(body, &key)
	if er2 != nil {
		t.Fatal(er2.Error())
	}

	if key.Name != "Test Key" {
		t.Fatalf("Key Mismatch, Expected: %v, Received: %v\nIn Data: %s", "Test Key", key.Name, body)
	}
}

func TestGetModVersion(t *testing.T) {
	type args struct {
		w   http.ResponseWriter
		req *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestGetMods(t *testing.T) {
	type args struct {
		w   http.ResponseWriter
		in1 *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
