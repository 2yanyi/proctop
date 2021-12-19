package main

import (
	"fmt"
	"os"
	"os/exec"
	"r/flagset"
	"r/owspace"
	"r/top"
	"runtime"
	"time"
)

func main() {
	if !flagset.Init(BuildID) {
		return
	}

	clear()
	fmt.Printf(" Num count  mem   CPU%%  Name  (Total: %d%%)\n"+
		"------------------------------------------------------------\n", runtime.NumCPU()*100)

	owspace.New(func(w *owspace.Writer) {
		t := time.NewTicker(time.Millisecond * 2000)
		w.Write(top.Call(*flagset.Limit))
		for range t.C {
			w.Write(top.Call(*flagset.Limit))
		}
	})
}

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	_ = cmd.Run()
}

var BuildID = "0"
