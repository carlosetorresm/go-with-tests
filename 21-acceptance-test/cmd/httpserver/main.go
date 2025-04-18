package main

import (
	"log"
	"net/http"

	"github.com/carlosetorresm/go-with-tests/21-acceptance-test/adapters/httpserver"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", httpserver.NewHandler()))
}
