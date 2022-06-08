package scanner

import (
	"fmt"
	"os"
	"testing"
)

func Test(t *testing.T) {
	output := ""
	processes := Processes()
	for i := 0; i < len(processes); i++ {
		if processes[i] == nil {
			continue
		}
		output += fmt.Sprintf("%3d) %32s CMD %s\n",
			i+1,
			processes[i].Name,
			processes[i].Commandline,
		)
	}
	_ = os.WriteFile("process.txt", []byte(output), 0666)
}
