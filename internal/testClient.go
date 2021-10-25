// Internal is the place where I dump my test classes
package internal

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Performs a HTTP GET on the path specified and returns the response or an error
func TestClient(path string) ([]byte, error) {
	resp, err := http.Get(path) //nolint:gosec
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Unexpected Status From Server: %s", resp.Status)
	}

	err = resp.Body.Close()
	if err != nil {
		return nil, err
	}

	return body, nil
}
