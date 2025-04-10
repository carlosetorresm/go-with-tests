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
		{"pointers to things",
			&Person{"Carlos", Profile{29, "Mexico"}}, []string{"Carlos", "Mexico"}},
		{"slices",
			[]Profile{
				{29, "Mexico"},
				{34, "Reykjavik"},
			}, []string{"Mexico", "Reykjavik"}},
		{"arrays",
			[2]Profile{
				{29, "Mexico"},
				{34, "Reykjavik"},
			}, []string{"Mexico", "Reykjavik"}},
		{"maps",
			map[string]string{
				"Cow":   "Moo",
				"Sheep": "Baa",
			}, []string{"Moo", "Baa"}},
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

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Cow":   "Moo",
			"Sheep": "Baa",
		}
		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Moo")
		assertContains(t, got, "Baa")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{29, "Mexico"}
			aChannel <- Profile{33, "Katowice"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Mexico", "Katowice"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %v, to contain %q but it didn't", haystack, needle)
	}
}
