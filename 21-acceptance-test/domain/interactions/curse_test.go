package interactions_test

import (
	"testing"

	"github.com/carlosetorresm/go-with-tests/21-acceptance-test/domain/interactions"
	"github.com/carlosetorresm/go-with-tests/21-acceptance-test/specifications"
)

func TestCurse(t *testing.T) {
	specifications.CurseSpecification(t, specifications.CurseAdapter(interactions.Curse))
}
