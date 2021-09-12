package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"github.com/Zaprit/Reflow/Models"
	"github.com/Zaprit/Reflow/TechnicAPI"
	"github.com/gorilla/mux"
	"io/fs"
	"net/http"
)

//go:embed web/*
var static embed.FS

func notFound(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(404)
	file, err := static.ReadFile("web/404.html")
	if err != nil {
		fmt.Println("Error, 404 page not found")
		_, err2 := w.Write([]byte("404 not found"))
		if err2 != nil {
			panic("failed making a 404 page, somehow both the 404 page from static and the 404 page i just made with text both failed?!?!?!")
		}
		return
	}
	_, err = w.Write(file)
	if err != nil {
		panic("failed to display static 404 page")
	}
}

func main() {
	APIName, _ := json.Marshal(Models.DefaultInfo)
	fmt.Printf("Reflow %s API: \"%s\"", Models.DefaultInfo.Version, APIName)
	r := mux.NewRouter()

	// Static content from web folder
	contentStatic, _ := fs.Sub(static, "web")

	r.PathPrefix("/").Handler(http.FileServer(http.FS(contentStatic)))
	r.HandleFunc("/api", TechnicAPI.ApiRoot)
	r.HandleFunc("/api/mod", TechnicAPI.GetMods)
	r.NotFoundHandler = http.HandlerFunc(notFound)
	http.Handle("/", r)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("ERROR: Something went wrong while setting up server")
		panic(err)
	}
}