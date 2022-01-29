// Package static is the package containing static content
package static

import (
	"embed"
)

//go:embed static/*
var WebFS embed.FS
