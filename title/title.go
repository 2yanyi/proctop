package title

import (
	"fmt"
	"github.com/utilgo/execve"
	"net"
	"os"
	"os/exec"
	"r/colors"
	"runtime"
	"strings"
)

func LanAddress() map[string]string {
	address := make(map[string]string, 10)
	nia, _ := net.InterfaceAddrs()
	ni, _ := net.Interfaces()
	for i, it := range ni {
		mac := it.HardwareAddr.String()
		if mac == "" {
			continue
		}
		if addr, has := nia[i].(*net.IPNet); has {
			address[mac] = addr.IP.String()
		}
	}
	return address
}

func showNameplate() {
	var lanAddress string
	for _, addr := range LanAddress() {
		lanAddress += " " + addr
	}
	address := fmt.Sprintf("(%s)%s ", execve.Args("", []string{"whoami"}), lanAddress)
	logo := strings.Join([]string{"\u001B[1;30;42m", " ProcTop ", "\u001B[0m"}, "")
	thread := colors.Green(fmt.Sprintf("%d*Thread", runtime.NumCPU()), colors.Italic)
	fmt.Printf("%s %s %s / %s %s\n", logo, uname(), thread, release(), address)
}

func Show() {
	header()
	fmt.Print(colors.White(
		"\n Num Count  Memory                             Name    CPU%                           ", colors.Underscore))
	fmt.Println(colors.White(" 永不宕机(never downtime)", colors.Italic))
}

func header() {
	showNameplate()
	for i, cpu := range readCPUs() {
		fmt.Println(colors.White(fmt.Sprintf(" - CPU%d %s", i+1, cpu), colors.Dark))
	}
}

func readCPUs() []string {
	cpus := make(map[string][]string)
	modelName := ""

	// filter
	info := strings.Split(cat("/proc/cpuinfo"), "\n")
	for _, elem := range info {
		if strings.HasPrefix(elem, "model name") {
			modelName = elem
		} else if strings.HasPrefix(elem, "physical id") {
			if elem == "" {
				continue
			}
			cpus[elem] = append(cpus[elem], modelName)
		}
	}

	// return
	slice := make([]string, 0, len(cpus))
	for i := range cpus {
		slice = append(slice, threadJoin(cpus[i][0], len(cpus[i])))
	}
	if len(slice) == 0 {

		// Compatible Raspberry Pi
		for _, elem := range info {
			if strings.HasPrefix(elem, "Model") {
				return []string{threadJoin(elem, runtime.NumCPU())}
			}
		}

		// Compatible other
		for _, elem := range info {
			if strings.HasPrefix(elem, "model name") {
				return []string{threadJoin(elem, runtime.NumCPU())}
			}
		}
	}
	return slice
}

func threadJoin(model string, threadCount int) string {
	elem := fmt.Sprintf("%s %d*Thread\n", model, threadCount)
	for index, char := range elem {
		if char == ':' {
			elem = elem[index+1:]
			break
		}
	}
	return strings.TrimSpace(elem)
}

func uname() (r string) {
	i := 0
	kernel, _ := exec.Command("uname", "-rs").Output()
	for _, char := range kernel {
		if char == '.' {
			i++
			if i == 2 {
				break
			}
		}
		r += string(char)
	}
	return
}

func release() string {
	var name, version string
	for _, elem := range strings.Split(cat("/etc/os-release"), "\n") {
		if strings.HasPrefix(elem, "NAME=") {
			name = elem[5:]
			if strings.HasPrefix(name, `"`) {
				name = strings.TrimPrefix(name, `"`)
			}
			if strings.HasSuffix(name, `"`) {
				name = strings.TrimSuffix(name, `"`)
			}
		}
		if strings.HasPrefix(elem, "VERSION=") {
			version = elem[8:]
			if strings.HasPrefix(version, `"`) {
				version = strings.TrimPrefix(version, `"`)
			}
			if strings.HasSuffix(version, `"`) {
				version = strings.TrimSuffix(version, `"`)
			}
		}
	}
	return strings.Join([]string{name, version}, " ")
}

func cat(fp string) string {
	data, _ := os.ReadFile(fp)
	return string(data)
}
