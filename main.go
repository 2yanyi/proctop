package main

import (
	"os"
	"os/exec"
	"r/data/variable"
	"r/flagset"
	"r/owspace"
	"r/title"
	"r/top"
	"time"
)

func main() {

	// pprof
	// defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop() // github.com/pkg/profile

	if !flagset.Init(BuildID) {
		return
	}
	clear()
	title.Show()

	owspace.New(func(w *owspace.Writer) {
		t := time.NewTicker(time.Second * 3)
		w.Write(top.Call(*flagset.Limit))
		for range t.C {
			w.Write(top.Call(*flagset.Limit))
		}
	})
}

func clear() {
	cmd := exec.Command("clear")
	if variable.IsWin {
		cmd = exec.Command("cmd", "/c", "cls")
	}
	cmd.Stdout = os.Stdout
	_ = cmd.Run()
}

var BuildID = "0"
