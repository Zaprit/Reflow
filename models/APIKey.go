package models

import "time"

// APIKey is the struct that represents an API key for the technic platform
type APIKey struct {
	DBStructTemplate
	Name   string
	APIKey string ` gorm:"api_key" `
}

// TableName is the tabler interface for GORM to specify a custom name for a table
func (APIKey) TableName() string {
	return "keys"
}

// APIKeyVerifyResponse is the response if an api key successfully verifies
type APIKeyVerifyResponse struct {
	Valid     string    `json:"valid"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}
