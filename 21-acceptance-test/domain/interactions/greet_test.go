package interactions_test

import (
	"testing"

	go_specs_greet "github.com/carlosetorresm/go-with-tests/21-acceptance-test/domain/interactions"
	"github.com/carlosetorresm/go-with-tests/21-acceptance-test/specifications"
)

func TestGreet(t *testing.T) {
	specifications.GreetSpecification(t, specifications.GreetAdapter(go_specs_greet.Greet))
}
