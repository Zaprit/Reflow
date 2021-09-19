package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"strings"

	"io/fs"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/Zaprit/Reflow/config"
	"github.com/Zaprit/Reflow/database"
	"github.com/Zaprit/Reflow/models"
	"github.com/Zaprit/Reflow/technicapi"
)

//go:embed web/*
var static embed.FS

func notFound(w http.ResponseWriter, _ *http.Request) {
	file, err := static.ReadFile("web/404.html")

	w.WriteHeader(404)

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
	APIName, _ := json.Marshal(models.DefaultInfo)
	fmt.Printf("Reflow %s API: \"%s\"\n", models.DefaultInfo.Version, APIName)

	config.LoadConfig()
	config.LoadRepoConfig()

	err := database.GetDBInstance().AutoMigrate(&models.Mod{}, &models.ModVersion{}, &models.APIKey{}, models.Modpack{}, models.ModpackBuild{}, models.BuildModversion{})

	if err != nil {
		return
	}

	r := mux.NewRouter()
	// Static content from web folder
	contentStatic, _ := fs.Sub(static, "web")

	r.PathPrefix("/static").Handler(http.FileServer(http.FS(contentStatic)))
	r.HandleFunc("/api", technicapi.APIRoot)
	r.HandleFunc("/api/", technicapi.APIRoot)
	r.HandleFunc("/api/mod", technicapi.GetMods)
	r.HandleFunc("/api/mod/", technicapi.GetMods)
	r.HandleFunc("/api/mod/{slug}", technicapi.GetMod)
	r.HandleFunc("/api/mod/{slug}/", technicapi.GetMod)
	r.HandleFunc("/api/mod/{slug}/{version}", technicapi.GetModVersion)
	r.HandleFunc("/api/mod/{slug}/{version}/", technicapi.GetModVersion)

	r.HandleFunc("/api/modpack", technicapi.GetModpacks)
	r.HandleFunc("/api/modpack/", technicapi.GetModpacks)
	r.HandleFunc("/api/modpack/{slug}", technicapi.GetModpack)
	r.HandleFunc("/api/modpack/{slug}/", technicapi.GetModpack)
	r.HandleFunc("/api/modpack/{slug}/{build}", technicapi.GetBuild)
	r.HandleFunc("/api/modpack/{slug}/{build}/", technicapi.GetBuild)
	r.NotFoundHandler = http.HandlerFunc(notFound)

	http.Handle("/", r)
	err = http.ListenAndServe(":8080", loggingMiddleware(r))

	if err != nil {
		fmt.Println("ERROR: Something went wrong while setting up server")
		panic(err)
	}
}

// this addresses trailing slashes and logs HTTP requests
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		fmt.Printf("%s: %s %s\n", strings.Split(r.RemoteAddr, ":")[0], r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
