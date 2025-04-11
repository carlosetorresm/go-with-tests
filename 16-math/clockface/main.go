package main

import (
	"os"
	"time"

	clockface "example.com/hello/16-math"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
