package flagset

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"time"
)

const binary = "proctop"

var fs = flag.NewFlagSet(binary, flag.ExitOnError)
var (
	__version       = fs.Bool("version", false, "show version information")
	__diskWriteRate = fs.Bool("diskw", false, "Disk write rate Test")
	Limit           = fs.Int("l", 10, "limit")
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
		fmt.Printf("v%s \n", BuildID)
		return
	}

	if *__diskWriteRate {
		fmt.Printf("Write rate %.1fMB/s\n", diskWriteRate()/1024/1024)
		return
	}

	return true
}

func diskWriteRate() float64 {
	size := 1024 * 1024 * 128
	buf := bytes.Buffer{}
	for i := 0; i < size; i++ {
		buf.Write([]byte{'0'})
	}
	start := time.Now()
	if err := os.WriteFile(".bytes", buf.Bytes(), 0666); err != nil {
		println("data write error: " + err.Error())
		return 0
	}
	duration := time.Now().Sub(start)
	second := float64(size) / (float64(duration) / float64(time.Second))
	_ = os.Remove(".bytes")
	return second
}
