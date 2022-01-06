package scanner

import (
	"fmt"
	"os"
	"testing"
)

func Test(t *testing.T) {
	output := ""
	processes := Processes()
	for i, proc := range processes {
		if proc.Process == nil {
			continue
		}
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
	_ = os.WriteFile("process.txt", []byte(output), 0666)
}
