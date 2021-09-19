package technicapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Zaprit/Reflow/database"
	"github.com/Zaprit/Reflow/models"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
)

// VerifyKey is the endpoint that verifies a key
func VerifyKey(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)

	var key models.APIKey

	result := database.GetDBInstance().Take(&key, "api_key = ?", vars["key"])

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		fmt.Errorf("this is not nessecerally an error, it just means that someone used the wrong API key, " +
			"check it's been entered correctly and try again\n")
		_, err := w.Write(models.APIErrorJSON("Invalid key provided."))

		if err != nil {
			panic(err.Error())
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
