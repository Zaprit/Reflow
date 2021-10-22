// Package models contains the structs for the various things
package models

import (
	"encoding/json"
	"time"
)

// APIInfo is the information sent to the technic client to verify that we're definitely talking to a solder server.
// In reality the client isn't very picky.
type APIInfo struct {
	Name    string ` json:"api" `
	Version string ` json:"version" `
	Stream  string ` json:"stream" `
}

// APIError is the representation of an error sent by solder as a struct
// An Example of an error:
//  {"error":"Mod does not exist"}
type APIError struct {
	Message string ` json:"error" `
}

// APIErrorJSON takes an error string, marshals it to JSON and returns it in a solder compatible format
func APIErrorJSON(e string) []byte {
	out, _ := json.Marshal(APIError{Message: e})
	return out
}

// DBStructTemplate most all the tables in the database have these fields.
// This is similar to the GORM one but without the deletedAt field.
type DBStructTemplate struct {
	ID        int32     ` gorm:"primaryKey" `
	CreatedAt time.Time ` json:"-" `
	UpdatedAt time.Time ` json:"-" `
}