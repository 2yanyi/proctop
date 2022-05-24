package internal

import (
	"bytes"
	"fmt"
	"golang.org/x/sys/unix"
	"strconv"
	"strings"
)

func commandMerge(stdout, stderr *bytes.Buffer) string {
	if stderr.Len() != 0 {
		stdout.WriteString("\n")
		stdout.Write(stderr.Bytes())
	}
	return stdout.String()
}

func (it *Environment) vendor() *Environment {
	fp := "/sys/class/dmi/id/sys_vendor"
	if FileExist(fp) {
		it.Vendor = strings.TrimSpace(String(&fp))
	}
	return it
}

func (it *Environment) kernel() *Environment {
	var uname unix.Utsname
	if err := unix.Uname(&uname); err != nil {
		it.Kernel = "unknown"
		return it
	}
	it.Kernel = fmt.Sprintf("%s", uname.Release)
	return it
}

func (it *Environment) release() *Environment {
	var NAME, VERSION, ID, ID_LIKE string
	var v = func(s string, l int) string {
		return strings.Trim(s[l:], "\"")
	}
	fp := "/etc/os-release"
	if !FileExist(fp) {
		return it
	}
	for _, elem := range strings.Split(String(&fp), "\n") {
		switch {
		case strings.HasPrefix(elem, "NAME="):
			NAME = v(elem, 5)
		case strings.HasPrefix(elem, "VERSION="):
			VERSION = v(elem, 8)
		case strings.HasPrefix(elem, "ID="):
			ID = v(elem, 3)
		case strings.HasPrefix(elem, "ID_LIKE="):
			ID_LIKE = v(elem, 8)
		}
	}
	it.Name = strings.Join([]string{NAME, VERSION}, " ")
	it.Platform = ID
	if _, has := releaseSet[ID]; !has {
		if ID_LIKE != "" {
			it.Platform = strings.Fields(ID_LIKE)[0]
		}
	}
	return it
}

func (it *Environment) storage() *Environment {
	text := CommandArgs("", []string{"df", "/"})
	for i, elem := range strings.Split(text, "\n") {
		if i == 1 {
			values := strings.Fields(elem)
			if len(values) != 6 {
				continue
			}
			avail, _ := strconv.ParseInt(values[3], 0, 64)
			size, _ := strconv.ParseInt(values[1], 0, 64)
			it.Perf += fmt.Sprintf(" DF=[ Avail:%s / %s ]", SizeFormat(float64(avail*1000)), SizeFormat(float64(size*1000)))
			break
		}
		if i == 2 {
			values := strings.Fields(elem)
			if len(values) != 5 {
				continue
			}
			avail, _ := strconv.ParseInt(values[2], 0, 64)
			size, _ := strconv.ParseInt(values[0], 0, 64)
			it.Perf += fmt.Sprintf(" DF=[ Avail:%s / %s ]", SizeFormat(float64(avail*1000)), SizeFormat(float64(size*1000)))
			break
		}
	}
	return it
}

func (it *Environment) android() *Environment {
	it.Platform = strings.ToLower(CommandArgs("", []string{"uname", "-o"}))
	if it.Platform == "android" {
		it.Processor = CommandArgs("", []string{"getprop", "ro.config.cpu_info_display"})
		it.Vendor = CommandArgs("", []string{"getprop", "ro.product.manufacturer"})
		it.Name = "Android " + CommandArgs("", []string{"getprop", "ro.system.build.version.release"})
	}
	return it
}

var releaseSet = map[string]*struct{}{
	"fedora": nil,
	"rhel":   nil,
	"centos": nil,
	"debian": nil,
	"ubuntu": nil,
}
