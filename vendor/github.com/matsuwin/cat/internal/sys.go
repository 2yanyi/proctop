package internal

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

type Environment struct {
	Vendor    string `json:"vendor"`
	Name      string `json:"name"`
	Perf      string `json:"perf"`
	Processor string `json:"processor"`
	Graphics  string `json:"graphics,omitempty"`
	Platform  string `json:"platform"`
	Kernel    string `json:"kernel"`
	Init      string `json:"init,omitempty"`
}

func SystemInfo() *Environment {
	it := &Environment{}
	it.vendor().kernel().release().cpuTitle().storage()
	switch runtime.GOOS {

	case "windows":
		it.Platform = runtime.GOOS
		it.Kernel = "NT " + strings.Fields(it.Kernel)[0]
		it.Graphics = strings.Join(graphics(), ", ")

	case "linux":
		it.Kernel = "Linux " + strings.Split(it.Kernel, "-")[0]
		if fp, _ := exec.LookPath("systemctl"); fp != "" {
			it.Init = "systemd"
		} else if fp, _ = exec.LookPath("service"); fp != "" {
			it.Init = "upstart" // sysvinit
		} else {
			it.Init = "no init"
		}
		if "root" == os.Getenv("USER") {
			it.Graphics = strings.Join(graphics(), ", ")
		}
		if it.Platform == "" {
			it.android()
		}
	}

	it.Platform += "/" + runtime.GOARCH
	return it
}

func (it *Environment) cpuTitle() *Environment {
	it.Perf = fmt.Sprintf("%d & ", processorSpeed())
	stat, _ := cpu.Info()
	if len(stat) == 0 {
		return it
	}
	switch {
	case strings.HasPrefix(stat[0].ModelName, "AMD"):
		it.Processor = strings.TrimSpace(strings.Split(stat[0].ModelName, "with")[0])
	case strings.HasPrefix(stat[0].ModelName, "Intel"):
		it.Processor = strings.TrimSpace(strings.Split(stat[0].ModelName, "@")[0])
	default:
		it.Processor = stat[0].ModelName
		if it.Processor == "" {
			fp := "/proc/cpuinfo"
			for _, elem := range strings.Split(String(&fp), "\n") {
				if strings.HasPrefix(elem, "Hardware") {
					it.Processor = strings.TrimSpace(strings.Split(elem, ":")[1])
				}
				if strings.HasPrefix(elem, "Model") {
					it.Vendor = strings.TrimSpace(strings.Split(elem, ":")[1])
				}
			}
		}
	}
	it.Perf += fmt.Sprintf("Hertz=%.1fG.T%d", stat[0].Mhz/1000, runtime.NumCPU())
	info, _ := mem.VirtualMemory()
	if info != nil {
		it.Perf += fmt.Sprintf(" - Memory=%s", SizeFormat(float64(info.Total)))
	}
	return it
}
