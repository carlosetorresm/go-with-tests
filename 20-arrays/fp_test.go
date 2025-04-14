package sum

import (
	"strings"
	"testing"
)

func TestFind(t *testing.T) {
	t.Run("find first even number", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		firstEverNumber, found := Find(numbers, func(x int) bool {
			return x%2 == 0
		})
		AssertTrue(t, found)
		AssertEqual(t, firstEverNumber, 2)
	})

	type Person struct {
		Name string
	}
	t.Run("Find the best programmer", func(t *testing.T) {
		theBest := Person{Name: "Carlos Torres"}
		people := []Person{
			{Name: "Kent Beck"},
			{Name: "Martin Fowler"},
			theBest,
		}

		isBest := func(p Person) bool {
			return strings.Contains(p.Name, "Carlos")
		}
		king, found := Find(people, isBest)

		AssertTrue(t, found)
		AssertEqual(t, king, theBest)
	})
}

func TestReduce(t *testing.T) {
	t.Run("multiplication of all elements", func(t *testing.T) {
		multiply := func(x, y int) int {
			return x * y
		}
		AssertEqual(t, Reduce([]int{1, 2, 3}, multiply, 1), 6)
	})

	t.Run("concatenate strings", func(t *testing.T) {
		concatenate := func(x, y string) string {
			return x + y
		}
		AssertEqual(t, Reduce([]string{"a", "b", "c"}, concatenate, ""), "abc")
	})
}

func TestBadBank(t *testing.T) {
	var (
		riya   = Account{Name: "Riya", Balance: 100}
		carlos = Account{Name: "Carlos", Balance: 75}
		adil   = Account{Name: "Adil", Balance: 200}
	)
	transactions := []Transaction{
		NewTransaction(carlos, riya, 100),
		NewTransaction(adil, carlos, 25),
	}

	newBalanceFor := func(acc Account) float64 {
		return NewBalanceFor(acc, transactions).Balance
	}
	AssertEqual(t, newBalanceFor(riya), 200)
	AssertEqual(t, newBalanceFor(carlos), 0)
	AssertEqual(t, newBalanceFor(adil), 175)

}

func AssertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v want %v given", got, want)
	}
}

func AssertTrue(t *testing.T, got bool) {
	t.Helper()
	if !got {
		t.Errorf("got %v, want true", got)
	}
}
