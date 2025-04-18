package specifications

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

type Greeter interface {
	Greet(name string) (string, error)
}

func GreetSpecification(t testing.TB, greeter Greeter) {
	got, err := greeter.Greet("Carlos")
	assert.NoError(t, err)
	assert.Equal(t, got, "Hello, Carlos")
}

type MeanGreeter interface {
	Curse(name string) (string, error)
}

func CurseSpecification(t testing.TB, meany MeanGreeter) {
	got, err := meany.Curse("Carlos")
	assert.NoError(t, err)
	assert.Equal(t, got, "Yucky Carlos!")
}
