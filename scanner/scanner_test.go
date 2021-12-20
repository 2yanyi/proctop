package scanner

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func Test(t *testing.T) {
	output := ""
	processes := Processes()
	for i, proc := range processes {
		cmdline, _ := proc.Process.Cmdline()
		ppid, _ := proc.Process.Ppid()
		output += fmt.Sprintf("%3d) %5d %5d %32s CMD %s\n",
			i+1,
			ppid,
			proc.Process.Pid,
			proc.Name,
			cmdline,
		)
	}
	_ = ioutil.WriteFile("process.txt", []byte(output), 0666)
}
