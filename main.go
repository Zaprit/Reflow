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
	"os"
	"path/filepath"
)

//go:embed web/*
var static embed.FS

//go:embed conf/reflow.conf.sample
var defaultConfig string

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
	fmt.Printf("Reflow %s API: \"%s\"\n", Models.DefaultInfo.Version, APIName)

	_, err := os.Stat("conf")
	if err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir("conf", 0755)
			if err != nil {
				panic(err)
			}
		}
		info, _ := os.Stat("conf")

		if !info.IsDir() {
			err = os.Remove("conf")
			if err != nil {
				panic(err.Error())
			}

			err = os.Mkdir("conf", 0755)
			if err != nil {
				panic(err.Error())
			}
		}
		conf, err := os.Create(filepath.FromSlash("./conf/reflow.conf"))
		if err != nil {
			panic(err.Error())
		}
		_, err = conf.WriteString(defaultConfig)
		if err != nil {
			panic(err.Error())
		}
		err = conf.Sync()
		if err != nil {
			panic(err.Error())
		}
		err = conf.Close()
		if err != nil {
			panic(err.Error())
		}

	}

	r := mux.NewRouter()
	// Static content from web folder
	contentStatic, _ := fs.Sub(static, "web")

	r.PathPrefix("/static").Handler(http.FileServer(http.FS(contentStatic)))
	r.HandleFunc("/api", TechnicAPI.ApiRoot)
	r.HandleFunc("/api/mod", TechnicAPI.GetMods)
	r.NotFoundHandler = http.HandlerFunc(notFound)
	http.Handle("/", r)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("ERROR: Something went wrong while setting up server")
		panic(err)
	}
}
