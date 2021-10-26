// Package static is the package containing static static content
package static

import (
	"embed"
	"io/fs"
	"log"
)

//go:embed static/*
var WebFS embed.FS

var WebFSFixedPath fs.FS

func InitStatic() {
	var err error
	WebFSFixedPath, err = fs.Sub(WebFS, "static")

	if err != nil {
		log.Panicf("Error Loading Static Content: %s", err.Error())
	}
}
