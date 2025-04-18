package main

import (
	"log"
	"net/http"

	"github.com/carlosetorresm/go-with-tests/21-acceptance-test/adapters/httpserver"
)

func main() {

	serverMux := http.NewServeMux()

	serverMux.HandleFunc("/greet", httpserver.GreetHandler)
	serverMux.HandleFunc("/curse", httpserver.CurseHandler)

	if err := http.ListenAndServe(":8080", serverMux); err != nil {
		log.Fatal(err)
	}
}
