package solderAPI

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Zaprit/Reflow/config"
	"github.com/Zaprit/Reflow/database"
	"github.com/Zaprit/Reflow/models"
	"github.com/gorilla/mux"
)

// APIRoot is the root function that identifies a solder compatible api
// Stock solder returns:
// 	{"api":"TechnicSolder","version":"v0.7.7","stream":"DEV"}
// I've never seen it not be DEV and as far as I can tell it doesn't matter what the api or version attributes are.
// (yes I did ask the Technic devs)
func APIRoot(w http.ResponseWriter, _ *http.Request) {
	out, _ := json.Marshal(config.DefaultInfo)
	_, err := w.Write(out)

	if err != nil {
		fmt.Printf("Error In /ApiRoot: %s", err.Error())
	}
}

// VerifyKey is the endpoint that verifies a key
func VerifyKey(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	key, err := database.GetKey(vars["key"])

	if err != nil {
		_, err = w.Write(models.APIErrorJSON("Invalid key provided."))
		if err != nil {
			panic(err)
		}
		return
	}

	out, err := json.Marshal(models.APIKeyVerifyResponse{
		Valid:     "Key validated.",
		Name:      key.Name,
		CreatedAt: key.CreatedAt,
	})

	if err != nil {
		panic(err.Error())
	}

	_, err = w.Write(out)

	if err != nil {
		panic(err.Error())
	}
}
