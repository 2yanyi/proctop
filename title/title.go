package title

import (
	"fmt"
	"github.com/matsuwin/cat"
	"github.com/utilgo/execve"
	"r/colors"
	"strings"
)

func Show() {
	info := cat.SystemInfo()
	logo := strings.Join([]string{"\u001B[1;30;42m", " procTop ", "\u001B[0m"}, "")
	address := fmt.Sprintf("(%s) %s ", execve.Args("", []string{"whoami"}), strings.Join(info.LanAddress, ", "))
	fmt.Printf("%s (%s) %s %s / %s\n%s\n", logo, info.Kernel, info.Perf,
		colors.Fuchsia("<"+info.Processor+">", colors.Italic),
		info.Name, address)

	fmt.Print(colors.White("\n Num     PPID  Count  Memory                             "+
		"Name    CPU%  Thread  FD             FIO(r/w)                                                             \n", colors.Underscore))
}
