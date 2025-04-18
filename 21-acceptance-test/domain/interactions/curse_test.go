package interactions_test

import (
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/carlosetorresm/go-with-tests/21-acceptance-test/domain/interactions"
	"github.com/carlosetorresm/go-with-tests/21-acceptance-test/specifications"
)

func TestCurse(t *testing.T) {
	specifications.CurseSpecification(t, specifications.CurseAdapter(interactions.Curse))

	t.Run("default name to world if it's an empty string", func(t *testing.T) {
		assert.Equal(t, "Yucky World!", interactions.Curse(""))
	})
}
