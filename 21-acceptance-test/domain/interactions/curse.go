package interactions

import "fmt"

func Curse(name string) string {
	if name == "" {
		name = "World"
	}
	return fmt.Sprintf("Yucky %s!", name)
}
