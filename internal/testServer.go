package internal

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// StartTestServer Creates a http server on the specified path with the specified handler, this is really only for use in tests
func StartTestServer(path string, handler http.HandlerFunc) {
	r := mux.NewRouter()

	r.HandleFunc(path, handler)

	http.Handle("/", r)

	err := http.ListenAndServe("localhost:8069", nil)

	if err != nil {
		fmt.Println("ERROR: Something went wrong while setting up server")
		panic(err)
	}

}
