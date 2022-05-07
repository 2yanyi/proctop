package title

import (
	"fmt"
	"github.com/matsuwin/cat"
	"github.com/utilgo/execve"
	"r/colors"
	"strings"
)

func Show() {
	logo := strings.Join([]string{"\u001B[1;30;42m", " procTop ", "\u001B[0m"}, "")
	info := cat.SystemInfo()
	address := fmt.Sprintf("(%s) %s ", execve.Args("", []string{"whoami"}), strings.Join(info.LanAddress, ", "))
	if info.Graphics != "" {
		info.Processor = colors.Fuchsia("<"+info.Processor+" & "+info.Graphics+">", colors.Italic)
	} else {
		info.Processor = colors.Fuchsia("<"+info.Processor+">", colors.Italic)
	}
	fmt.Printf("%s (%s) %s %s / %s\n%s\n%s", logo, info.Kernel, info.Perf, info.Processor, info.Name, address, tableHead)
}

var tableHead = colors.White("\n Num     PPID  Count  Memory                             "+
	"Name    CPU%  Thread  FD             FIO(r/w)                                                             \n", colors.Underscore)
