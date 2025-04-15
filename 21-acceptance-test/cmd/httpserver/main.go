package main

import (
	"log"
	"net/http"

	go_specs_greet "github.com/carlosetorresm/go-with-tests/21-acceptance-test"
)

func main() {
	handler := http.HandlerFunc(go_specs_greet.Handler)
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}
