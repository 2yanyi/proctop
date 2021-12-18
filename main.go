package main

import (
	"fmt"
	"github.com/gosuri/uilive"
	"os"
	"os/exec"
	"r/top"
	"runtime"
	"strconv"
	"time"
)

func main() {
	var limit int
	if len(os.Args) > 1 {
		limit, _ = strconv.Atoi(os.Args[1])
	}
	if limit == 0 {
		limit = 10
	}

	clear()
	fmt.Printf(" Num count  mem   CPU%%  Name  (Total: %d%%)\n"+
		"------------------------------------------------------------\n",
		runtime.NumCPU()*100)

	canvas := uilive.New()
	canvas.Start()
	t := time.NewTicker(time.Millisecond * 1000)
	_, _ = fmt.Fprintf(canvas, "%s", top.Call(limit))
	for range t.C {
		_, _ = fmt.Fprintf(canvas, "%s", top.Call(limit))
	}
	canvas.Stop()
}

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	_ = cmd.Run()
}
