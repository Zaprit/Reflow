package technicapi

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/Zaprit/Reflow/config"
	"github.com/Zaprit/Reflow/internal"
	"github.com/Zaprit/Reflow/models"
)

func TestAPIRoot(t *testing.T) {
	go internal.StartTestServer("/api", APIRoot)

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
