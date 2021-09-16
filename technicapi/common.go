package technicapi

import (
	"encoding/json"
	"fmt"
	"github.com/Zaprit/Reflow/models"
	"net/http"
)

// APIRoot is the root function that identifies a solder compatible api
// Stock solder returns:
// 	{"api":"TechnicSolder","version":"v0.7.7","stream":"DEV"}
// I've never seen it not be DEV and as far as I can tell it doesn't matter what the api or version attributes are.
// (yes I did ask the Technic devs)
func APIRoot(w http.ResponseWriter, _ *http.Request) {
	out, _ := json.Marshal(models.DefaultInfo)
	_, err := w.Write(out)

	if err != nil {
		fmt.Printf("Error In /ApiRoot: %s", err.Error())
	}
}
