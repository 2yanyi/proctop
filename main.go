package main

import (
	"fmt"
	"github.com/gosuri/uilive"
	"os"
	"os/exec"
	"r/flagset"
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
		"------------------------------------------------------------\n",
		runtime.NumCPU()*100)

	canvas := uilive.New()
	canvas.Start()
	t := time.NewTicker(time.Millisecond * 2000)
	_, _ = fmt.Fprintf(canvas, "%s", top.Call(*flagset.Limit))
	for range t.C {
		_, _ = fmt.Fprintf(canvas, "%s", top.Call(*flagset.Limit))
	}
	canvas.Stop()
}

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	_ = cmd.Run()
}

var BuildID = "0"
