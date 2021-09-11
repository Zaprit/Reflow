package main

import (
	"github.com/Zaprit/Reflow/src/Config"
	"github.com/Zaprit/Reflow/src/TechnicAPI"
	"net/http"
)

func main() {
	Config.ConfPaths = []string{"conf"}
	http.HandleFunc("/api", TechnicAPI.ApiRoot)
	http.HandleFunc("/api/mod", TechnicAPI.GetMods)
	http.ListenAndServe(":8080", nil)
}