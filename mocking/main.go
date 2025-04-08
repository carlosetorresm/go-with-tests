package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	finalWord     = "Go!"
	coundownStart = 3
)

func Countdown(w io.Writer) {
	for i := coundownStart; i > 0; i-- {
		fmt.Fprintln(w, i)
		time.Sleep(1 * time.Second)
	}
	fmt.Fprint(w, finalWord)
}

func main() {
	Countdown(os.Stdout)
}
