package reflection

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{"struct with one string field",
			struct {
				Name string
			}{"Carlos"}, []string{"Carlos"}},
		{"struct with two string field",
			struct {
				Name string
				City string
			}{"Carlos", "Mexico"}, []string{"Carlos", "Mexico"}},
		{"struct with non string field",
			struct {
				Name string
				Age  int
			}{"Carlos", 29}, []string{"Carlos"}},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})
			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %q, want %q", got, test.ExpectedCalls)
			}
		})
	}
}
