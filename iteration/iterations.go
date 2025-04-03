package iteration

import (
	"strings"
)

func Repeat(character string, repeateCount int) string {
	var repeated strings.Builder
	for i := 0; i < repeateCount; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()
}
