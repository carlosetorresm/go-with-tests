package specifications

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

type MeanGreeter interface {
	Curse(name string) (string, error)
}

func CurseSpecification(t testing.TB, meany MeanGreeter) {
	got, err := meany.Curse("Carlos")
	assert.NoError(t, err)
	assert.Equal(t, got, "Yucky Carlos!")
}
