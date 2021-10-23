package internal

import (
	"io/ioutil"
	"log"
	"net/http"
)

// Performs a HTTP GET on the path specified and returns the response or an error
func TestClient(path string) ([]byte, error) {
	resp, err := http.Get("http://localhost:8069" + path)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = resp.Body.Close()
	if err != nil {
		return nil, err
	}
	return body, nil
}
