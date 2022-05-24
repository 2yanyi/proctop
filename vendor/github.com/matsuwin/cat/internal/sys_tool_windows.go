package internal

import (
	"bytes"
	"github.com/jaypipes/ghw"
	"github.com/shirou/gopsutil/v3/host"
	"golang.org/x/text/encoding/simplifiedchinese"
)

func commandMerge(stdout, stderr *bytes.Buffer) string {
	if stderr.Len() != 0 {
		stdout.WriteString("\n")
		stdout.Write(stderr.Bytes())
	}
	text, _ := simplifiedchinese.GBK.NewDecoder().String(stdout.String())
	return text
}

func (it *Environment) vendor() *Environment {
	if product, _ := ghw.Product(); product != nil {
		it.Vendor = product.Vendor
	}
	return it
}

func (it *Environment) kernel() *Environment {
	return it
}

func (it *Environment) release() *Environment {
	info, err := host.Info()
	if err != nil {
		it.Name = "unknown"
		it.Kernel = "unknown"
		return it
	}
	it.Name = info.Platform
	it.Kernel = info.KernelVersion
	return it
}

func (it *Environment) storage() *Environment {
	return it
}

func (it *Environment) android() *Environment {
	return it
}
