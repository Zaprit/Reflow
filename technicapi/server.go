package technicapi

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	"github.com/Zaprit/Reflow/config"
	"github.com/Zaprit/Reflow/database"
	"github.com/Zaprit/Reflow/models"
)

// InitConfig Initializes the configuration objects and loads the config required for the server to start
func InitConfig() {
	config.LoadConfig()
	config.LoadRepoConfig()
}

// InitDB Initializes and migrates the DB ready for use by the server
func InitDB() {
	err := database.GetDBInstance().AutoMigrate(
		&models.Mod{}, &models.ModVersion{},
		&models.APIKey{}, &models.Modpack{},
		&models.ModpackBuild{}, &models.BuildModversion{})
	if err != nil {
		panic("Failed to migrate tables")
	}
}

// StartServer Starts the http server ready to accept connections
func StartServer(listenAddress string) {
	r := mux.NewRouter()

	// TODO: Fix static loading, this was broken when main split up, make static package and add an extractor class

	// Static content from web folder
	// contentStatic, _ := fs.Sub(static, "web")

	// r.PathPrefix("/static").Handler(http.FileServer(http.FS(contentStatic)))
	r.HandleFunc("/api", APIRoot)
	r.HandleFunc("/api/verify/{key}", VerifyKey)

	r.HandleFunc("/api/mod", GetMods)
	r.HandleFunc("/api/mod/{slug}", GetMod)
	r.HandleFunc("/api/mod/{slug}/{version}", GetModVersion)

	r.HandleFunc("/api/modpack", GetModpacks)
	r.HandleFunc("/api/modpack/", GetModpacks)
	r.HandleFunc("/api/modpack/{slug}", GetModpack)
	r.HandleFunc("/api/modpack/{slug}/{build}", GetBuild)
	// r.NotFoundHandler = http.HandlerFunc(notFound)

	http.Handle("/", r)

	err := http.ListenAndServe(listenAddress, loggingMiddleware(r))

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

//func notFound(w http.ResponseWriter, _ *http.Request) {
//	file, err := static.ReadFile("web/404.html")
//
//	w.WriteHeader(404)
//
//	if err != nil {
//		fmt.Println("Error, 404 page not found")
//
//		_, err2 := w.Write([]byte("404 not found"))
//
//		if err2 != nil {
//			panic("failed making a 404 page, somehow both the 404 page from static and the 404 page i just made with text both failed?!?!?!")
//		}
//
//		return
//	}
//
//	_, err = w.Write(file)
//
//	if err != nil {
//		panic("failed to display static 404 page")
//	}
//}
