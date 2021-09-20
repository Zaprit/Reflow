package technicapi

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"

	"github.com/Zaprit/Reflow/database"
	"github.com/Zaprit/Reflow/models"
)

// VerifyKey is the endpoint that verifies a key
func VerifyKey(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	var key models.APIKey

	result := database.GetDBInstance().Take(&key, "api_key = ?", vars["key"])

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
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
