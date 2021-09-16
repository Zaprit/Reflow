// Package models contains the structs for the various things
package models

// APIInfo is the information required to make the technic client believe that we are definitely talking to a solder server.
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

// DefaultInfo is the default APIInfo for reflow.
var DefaultInfo = APIInfo{Name: "Reflow", Version: "v0.1", Stream: "DEV"}
