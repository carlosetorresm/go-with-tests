package main

import (
	"os"
	"time"

	"example.com/hello/16-math/svg"
)

func main() {
	t := time.Now()
	svg.Write(os.Stdout, t)
}
