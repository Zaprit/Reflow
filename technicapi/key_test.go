package technicapi

import (
	"encoding/json"
	"testing"

	"github.com/Zaprit/Reflow/database"
	"github.com/Zaprit/Reflow/internal"
	"github.com/Zaprit/Reflow/models"
)

const testKey = "Test Key"

// TestVerifyKey Tests the VerifyKey REST endpoint
func TestVerifyKey(t *testing.T) {

	database.GetDBInstance().Create(&models.APIKey{
		Name:   testKey,
		APIKey: "testkey2341352463",
	})

	body, err := internal.TestClient("/api/verify/testkey2341352463")
	if err != nil {
		t.Fatal(err.Error())
	}

	var key models.APIKeyVerifyResponse

	er2 := json.Unmarshal(body, &key)
	if er2 != nil {
		t.Fatal(er2.Error())
	}

	if key.Name != testKey {
		t.Fatalf("Key Mismatch, Expected: %v, Received: %v\nIn Data: %s", testKey, key.Name, body)
	}

	shouldFail, err := internal.TestClient("/api/verify/ono")
	if err != nil {
		t.Fatal(err.Error())
	}

	var errorMessage models.APIError
	err3 := json.Unmarshal(shouldFail, &errorMessage)

	if err3 != nil {
		t.Fatal(err3.Error())
	}

	if errorMessage.Message != "Invalid key provided." {
		t.Fatal("Somehow the key passed despite being invalid")
	}

}
