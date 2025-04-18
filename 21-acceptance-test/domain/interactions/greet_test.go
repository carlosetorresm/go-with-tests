package interactions_test

import (
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/carlosetorresm/go-with-tests/21-acceptance-test/domain/interactions"
	"github.com/carlosetorresm/go-with-tests/21-acceptance-test/specifications"
)

func TestGreet(t *testing.T) {
	specifications.GreetSpecification(t, specifications.GreetAdapter(interactions.Greet))

	t.Run("default name to world if it's an empty string", func(t *testing.T) {
		assert.Equal(t, "Hello, World", interactions.Greet(""))
	})
}
