package httpserver

import (
	"fmt"
	"net/http"

	"github.com/carlosetorresm/go-with-tests/21-acceptance-test/domain/interactions"
)

const (
	greetPath = "/greet"
	cursePath = "/curse"
)

func NewHandler() http.Handler {
	serverMux := http.NewServeMux()

	serverMux.HandleFunc(greetPath, replyWith(interactions.Greet))
	serverMux.HandleFunc(cursePath, replyWith(interactions.Curse))

	return serverMux
}

func replyWith(f func(name string) (interaction string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		fmt.Fprint(w, f(name))
	}
}
