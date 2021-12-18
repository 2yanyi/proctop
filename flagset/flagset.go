package flagset

import (
	"flag"
	"fmt"
	"os"
)

const (
	binary  = "proctop"
	version = "0.1"
)

var fs = flag.NewFlagSet(binary, flag.ExitOnError)
var (
	__version = fs.Bool("version", false, "show version information")
	Limit     = fs.Int("l", 10, "limit")
)

func Init(BuildID string) (_ bool) {
	if BuildID == "" {
		BuildID = "dev"
	}

	// 打印程序相关信息
	if len(os.Args) == 1 {
		os.Args = append(os.Args, []string{"-l", "10"}...)
	}
	_ = fs.Parse(os.Args[1:])

	// 查看程序版本
	if *__version {
		fmt.Printf("v%s.%s \n", version, BuildID)
		return
	}

	return true
}
