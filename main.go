package main

import (
	"fmt"
	"os"
	"os/exec"
	"r/flagset"
	"r/owspace"
	"r/top"
	"runtime"
	"strings"
	"time"
)

func main() {
	if !flagset.Init(BuildID) {
		return
	}

	if runtime.GOOS == "linux" {
		clear()
	}
	OSName, OSVersion := OSRelease()
	fmt.Printf(" Num Count  Memory                             Name    CPU%%  / Core*%d  %s %s\n"+
		"--------------------------------------------------------------------------------------\n",
		runtime.NumCPU(), OSName, OSVersion)

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

func cat(fp string) string {
	data, _ := os.ReadFile(fp)
	return string(data)
}

func OSRelease() (name, version string) {
	for _, elem := range strings.Split(cat("/etc/os-release"), "\n") {
		if strings.HasPrefix(elem, "NAME=") {
			name = elem[5:]
			if strings.HasPrefix(name, `"`) {
				name = strings.TrimPrefix(name, `"`)
			}
			if strings.HasSuffix(name, `"`) {
				name = strings.TrimSuffix(name, `"`)
			}
		}
		if strings.HasPrefix(elem, "VERSION=") {
			version = elem[8:]
			if strings.HasPrefix(version, `"`) {
				version = strings.TrimPrefix(version, `"`)
			}
			if strings.HasSuffix(version, `"`) {
				version = strings.TrimSuffix(version, `"`)
			}
		}
	}
	return
}

var BuildID = "0"
