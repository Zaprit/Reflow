package technicapi

import (
	"encoding/json"
	"testing"

	"github.com/Zaprit/Reflow/database"
	"github.com/Zaprit/Reflow/internal"
	"github.com/Zaprit/Reflow/models"
)

func TestVerifyKey(t *testing.T) {

	database.CreateDBInstance(database.DBConfig{
		Driver: "sqlite",
		DBName: "file::memory:?cache=shared",
	})

	err := database.GetDBInstance().AutoMigrate(&models.APIKey{})
	if err != nil {
		t.Fatal(err.Error())
	}

	database.GetDBInstance().Create(&models.APIKey{
		Name:   "Test Key",
		APIKey: "testkey2341352463",
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
