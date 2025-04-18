package httpserver

import (
	"fmt"
	"net/http"

	interactions "github.com/carlosetorresm/go-with-tests/21-acceptance-test/domain/interactions"
)

func GreetHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprint(w, interactions.Greet(name))
}

func CurseHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprint(w, interactions.Curse(name))
}
