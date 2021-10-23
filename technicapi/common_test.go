package technicapi

import (
	"encoding/json"
	"fmt"
	"github.com/Zaprit/Reflow/config"
	"github.com/Zaprit/Reflow/models"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/mux"
)

func TestAPIRoot(t *testing.T) {

	go startServer("/api", APIRoot)

	time.Sleep(time.Second)

	resp, err := http.Get("http://localhost:8069/api")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err.Error())
	}

	var info models.APIInfo

	err = json.Unmarshal(body, &info)

	if ! reflect.DeepEqual(info, config.DefaultInfo) {
		t.Fatal("OOF")
	}

}

func startServer(path string, handler http.HandlerFunc){
	r := mux.NewRouter()

	r.HandleFunc(path, handler)

	http.Handle("/", r)

	err := http.ListenAndServe("localhost:8069", loggingMiddleware(r))

	if err != nil {
		fmt.Println("ERROR: Something went wrong while setting up server")
		panic(err)
	}

}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		fmt.Printf("%s: %s %s\n", strings.Split(r.RemoteAddr, ":")[0], r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
