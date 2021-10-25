package technicapi

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"

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

	r := mux.NewRouter()
	r.HandleFunc("/api/verify/{key}", VerifyKey)

	ts := httptest.NewServer(r)

	body, err := internal.TestClient(ts.URL + "/api/verify/testkey2341352463")
	if err != nil {
		t.Errorf("Expected nil, Received %s", err.Error())
	}

	var key models.APIKeyVerifyResponse

	er2 := json.Unmarshal(body, &key)
	if er2 != nil {
		t.Fatal(er2.Error())
	}

	if key.Name != testKey {
		t.Fatalf("Key Mismatch, Expected: %v, Received: %v\nIn Data: %s", testKey, key.Name, body)
	}

	shouldFail, err := internal.TestClient(ts.URL + "/api/verify/ono")
	if err != nil {
		t.Fatal(err.Error())
	}

	ts.Close()

	var errorMessage models.APIError
	err3 := json.Unmarshal(shouldFail, &errorMessage)

	if err3 != nil {
		t.Fatal(err3.Error())
	}

	if errorMessage.Message != "Invalid key provided." {
		t.Fatal("Somehow the key passed despite being invalid")
	}
}
