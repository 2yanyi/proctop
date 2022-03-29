package title

import (
	"fmt"
	"github.com/utilgo/execve"
	"net"
	"os"
	"os/exec"
	"r/colors"
	"r/data/variable"
	"runtime"
	"strconv"
	"strings"
)

func LanAddress() map[string]string {
	address := make(map[string]string, 10)
	nia, _ := net.InterfaceAddrs()
	ni, _ := net.Interfaces()
	niN := len(ni) - 1
	for i := range nia {
		if niN < i {
			continue
		}
		if addr, has := nia[i].(*net.IPNet); has {
			address[ni[i].HardwareAddr.String()] = addr.IP.String()
		}
	}
	return address
}

func Show() {
	showNameplate()
	cpu := showCPUModel()
	fmt.Print(colors.White(
		"\n Num Count  Memory                             Name    CPU%                                     ", colors.Underscore))
	switch cpu {
	case "amd":
		fmt.Println(colors.White(" AMD YES!", colors.Italic))
	case "intel":
		fmt.Println(colors.White(" Intel NB!", colors.Italic))
	default:
		fmt.Println(colors.White(" Hello World", colors.Italic))
	}
}

func showNameplate() {
	var lanAddress string
	for _, addr := range LanAddress() {
		if strings.Contains(addr, ":") {
			continue
		}
		lanAddress += ", " + addr
	}
	address := fmt.Sprintf("(%s)%s ", execve.Args("", []string{"whoami"}), lanAddress)
	logo := strings.Join([]string{"\u001B[1;30;42m", " ProcTop ", "\u001B[0m"}, "")
	thread := colors.Fuchsia(fmt.Sprintf("%d*Thread", runtime.NumCPU()), colors.Italic)
	fmt.Printf("%s %s %s / %s %s\n", logo, uname(), thread, release(), address)
}

func showCPUModel() string {
	cpuTag := ""
	if variable.IsWin {
		cpuTag = cpuModelWindows()
	} else {
		cpuTag = cpuModelLinux()
	}
	fmt.Printf("[  CPU  ]  %s\n", colors.Green(cpuTag, colors.Italic))

	if strings.HasPrefix(cpuTag, "AMD") {
		return "amd"
	} else {
		return "intel"
	}
}

func cpuModelLinux() (_ string) {
	var name, ghz string
	for _, elem := range strings.Split(cat("/proc/cpuinfo"), "\n") {
		if strings.HasPrefix(elem, "model name") {
			name = splitValue(elem)
		}
		if strings.HasPrefix(elem, "cpu MHz") {
			num, _ := strconv.ParseFloat(splitValue(elem), 64)
			ghz = fmt.Sprintf("%.1fGHz", num/1000)
			break
		}
	}
	return fmt.Sprintf("%s @ %s", name, ghz)
}

func cpuModelWindows() (_ string) {
	var name, ghz string
	text := execve.Args("", []string{"wmic", "cpu", "list", "brief"})
	for _i, line := range strings.Split(text, "\n") {
		if _i == 0 {
			continue
		}
		values := strings.Split(line, "  ")
		vas := make([]string, 0, len(values))
		for i := len(values) - 1; i >= 0; i-- {
			if values[i] == "" {
				continue
			}
			vas = append(vas, values[i])
		}
		for i := range vas {
			if i == 1 {
				name = strings.TrimSpace(vas[i])
			}
			if i == 2 {
				num, _ := strconv.ParseFloat(vas[i], 64)
				ghz = fmt.Sprintf("%.1fGHz", num/1000)
			}
		}
	}
	return fmt.Sprintf("%s @ %s", name, ghz)
}

func splitValue(s string) string {
	vs := strings.Split(s, ":")
	if len(vs) < 2 {
		return "<null>"
	}
	return strings.TrimSpace(vs[1])
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
	if variable.IsWin {
		return "Windows NT"
	}
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
	if variable.IsWin {
		return "^_^"
	}
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
