package iteration

import "fmt"

const repeateCount = 5

func Repeat(character string) string {
	var repeated string
	for i := 0; i < repeateCount; i++ {
		repeated += character
	}
	return repeated
}

func main() {
	fmt.Println(Repeat("a"))
}
