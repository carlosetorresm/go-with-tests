package iteration

import (
	"fmt"
	"strings"
)

const repeateCount = 5

func Repeat(character string) string {
	var repeated strings.Builder
	for i := 0; i < repeateCount; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()
}

func main() {
	fmt.Println(Repeat("a"))
}
