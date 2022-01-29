package solderAPI

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

const testKey = "test123!!!"

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

// TestMain Bootstraps the tests for solderAPI by creating a database instance and starting a http server
func TestMain(m *testing.M) {
	database.CreateDBInstance(&database.DBConfig{
		Driver: "sqlite",
		DBName: "file::memory:?cache=shared",
	})

	database.InitDB()

	code := m.Run()

	os.Exit(code)
}

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
