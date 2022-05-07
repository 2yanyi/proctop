package internal

import (
	"github.com/jaypipes/ghw"
	"github.com/shirou/gopsutil/v3/host"
	"os/exec"
	"runtime"
	"strings"
)

type Environment struct {
	Vendor     string   `json:"vendor"`
	Name       string   `json:"name"`
	Perf       string   `json:"perf"`
	Processor  string   `json:"processor"`
	Graphics   string   `json:"graphics,omitempty"`
	Platform   string   `json:"platform"`
	Kernel     string   `json:"kernel"`
	Init       string   `json:"init,omitempty"`
	LanAddress []string `json:"lanAddress,omitempty"`
}

func SystemInfo() *Environment {
	it := &Environment{}
	it.Perf, it.Processor = cpuTitle()
	it.Graphics = strings.Join(graphics(), ", ")
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

func vendor() (_ string) {
	if product, _ := ghw.Product(); product != nil {
		return product.Vendor
	}
	return
}

func graphics() []string {
	drivers := make([]string, 0)
	info, err := ghw.GPU()
	if err != nil {
		Stderr(err.Error())
	}
	if info == nil {
		return nil
	}
	for _, driver := range info.GraphicsCards {
		if driver == nil {
			continue
		}
		if driver.DeviceInfo == nil {
			continue
		}
		if driver.DeviceInfo.Product == nil {
			continue
		}
		if driver.DeviceInfo.Product.Name == "SVGA II Adapter" {
			continue
		}
		if strings.Contains(driver.DeviceInfo.Product.Name, "Graphics") {
			continue
		}
		drivers = append(drivers, driver.DeviceInfo.Product.Name)
	}
	return drivers
}
