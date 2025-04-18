package httpserver

import (
	"fmt"
	"net/http"

	go_specs_greet "github.com/carlosetorresm/go-with-tests/21-acceptance-test/domain/interactions"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprint(w, go_specs_greet.Greet(name))
}
