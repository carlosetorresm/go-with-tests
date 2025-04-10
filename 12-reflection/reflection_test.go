package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         any
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
		{"struct nested field",
			Person{"Carlos", Profile{29, "Mexico"}}, []string{"Carlos", "Mexico"}},
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
