package internal

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"os/exec"
	"runtime"
	"strings"
)

type Environment struct {
	Vendor     string   `json:"vendor"`
	Name       string   `json:"name"`
	Perf       string   `json:"perf"`
	Processor  string   `json:"processor"`
	Platform   string   `json:"platform"`
	Kernel     string   `json:"kernel"`
	Init       string   `json:"init,omitempty"`
	LanAddress []string `json:"lanAddress"`
}

func SystemInfo() *Environment {
	it := &Environment{}
	it.Perf, it.Processor = cpuTitle()
	it.Vendor = vendor()
	stat, _ := host.Info()
	if stat == nil {
		return it
	}
	switch runtime.GOOS {
	case "windows":
		{
			it.Name = stat.Platform
			it.Platform = runtime.GOOS
			it.Kernel = "NT " + strings.Fields(stat.KernelVersion)[0]
		}
	case "linux":
		{
			it.Name = release()
			it.Platform = stat.Platform
			it.Kernel = "Linux " + strings.Split(stat.KernelVersion, "-")[0]
			if fp, _ := exec.LookPath("systemctl"); fp != "" {
				it.Init = "systemd"
			} else if fp, _ = exec.LookPath("service"); fp != "" {
				it.Init = "upstart" // sysvinit
			} else {
				it.Init = "no init"
			}
		}
	}
	it.LanAddress = LanAddress()

	return it
}

func cpuTitle() (perf, processor string) {
	stat, _ := cpu.Info()
	if len(stat) != 0 {
		perf = fmt.Sprintf("Hertz(%.2fG).T%d", stat[0].Mhz/1000, runtime.NumCPU())
		processor = strings.TrimSpace(stat[0].ModelName)
		info, _ := mem.VirtualMemory()
		if info != nil {
			perf += fmt.Sprintf(" Memory(%s)", SizeFormat(float64(info.Total)))
		}
	}
	return
}

func release() string {
	var name, version string
	for _, elem := range strings.Split(String("/etc/os-release"), "\n") {
		if strings.HasPrefix(elem, "NAME=") {
			name = strings.Trim(elem[5:], `"`)
		}
		if strings.HasPrefix(elem, "VERSION=") {
			version = strings.Trim(elem[8:], `"`)
		}
	}
	return strings.Join([]string{name, version}, " ")
}
