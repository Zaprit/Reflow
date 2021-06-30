package main

import (
	"fmt"
	"math/rand"
)

/*
The code from this was taken from the revel cmd new.go file
*/

// Used to generate a new secret key
const alphaNumeric = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

// Generate a secret key
func generateSecret() string {
	chars := make([]byte, 64)
	for i := 0; i < 64; i++ {
		chars[i] = alphaNumeric[rand.Intn(len(alphaNumeric))]
	}
	return string(chars)
}
func main() {
	fmt.Println(generateSecret())
}
