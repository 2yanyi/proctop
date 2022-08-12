package main

import (
	"fmt"
	"r/pkg/owspace"
	"time"
)

func main() {
	owspace.New(func(w *owspace.Writer) {
		for i := 0; i <= 100; i++ {
			w.Write(fmt.Sprintf("Downloading... (%d/%d)\n", i, 100))
			time.Sleep(time.Millisecond * 10)
		}
	})
}
