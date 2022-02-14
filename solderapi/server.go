package solderapi

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Zaprit/Reflow/dashboard"

	"github.com/Zaprit/Reflow/static"

	"github.com/gorilla/mux"
)

// StartServer Starts the http server ready to accept connections
func StartServer(listenAddress string) {
	r := mux.NewRouter()

	r.PathPrefix("/static").Handler(http.FileServer(http.FS(static.WebFS)))
	r.HandleFunc("/", dashboard.WebDashboardRoot)
	r.HandleFunc("/index.html", dashboard.WebDashboardRoot)
	r.HandleFunc("/mod", dashboard.WebDashboardModList)
	r.HandleFunc("/api", APIRoot)
	r.HandleFunc("/api/verify/{key}", VerifyKey)

	r.HandleFunc("/api/mod", GetMods)
	r.HandleFunc("/api/mod/{slug}", GetMod)
	r.HandleFunc("/api/mod/{slug}/{version}", GetModVersion)

	r.HandleFunc("/api/modpack", GetModpacks)
	r.HandleFunc("/api/modpack/", GetModpacks)
	r.HandleFunc("/api/modpack/{slug}", GetModpack)
	r.HandleFunc("/api/modpack/{slug}/{build}", GetBuild)
	r.NotFoundHandler = http.HandlerFunc(notFound)

	http.Handle("/", r)

	err := http.ListenAndServe(listenAddress, loggingMiddleware(r))

	if err != nil {
		fmt.Println("ERROR: Something went wrong while setting up server")
		panic(err.Error())
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

func notFound(w http.ResponseWriter, _ *http.Request) {
	file, err := static.WebFS.ReadFile("static/404.html")

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
