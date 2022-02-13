package database

import (
	"errors"
	"fmt"

	"github.com/Zaprit/Reflow/models"
	"gorm.io/gorm"
)

func GetKey(key string) (models.APIKey, error) {
	var out models.APIKey

	result := GetDBInstance().Take(&out, "api_key = ?", key)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return models.APIKey{}, fmt.Errorf("invalid API Key")
	}
	return out, nil
}
