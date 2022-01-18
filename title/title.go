package title

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func Get() {
	for i, cpu := range readCpus() {
		fmt.Printf(" CPU%d %s\n", i+1, cpu)
	}
	//for _, elem := range readGpuAndUSB() {fmt.Printf("%s\n", elem)}
	OSName, OSVersion := release()
	fmt.Printf("\n Num Count  Memory                             Name    CPU%%  / %s %s\n"+
		"--------------------------------------------------------------------------------------\n",
		OSName, OSVersion)
}

func readCpus() []string {
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
		for _, elem := range info {
			if strings.HasPrefix(elem, "Model") {
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

func readGpuAndUSB() []string {
	lsPCI, err := exec.Command("lspci").Output()
	if err != nil {
		return nil
	}

	// filter
	gpus := make([]string, 0)
	usbs := make([]string, 0)
	for _, line := range strings.Split(string(lsPCI), "\n") {
		if len(line) < 8 {
			continue
		}
		elem := line[8:]
		switch {
		case strings.HasPrefix(elem, "VGA"):
			for index, char := range elem {
				if char == ':' {
					elem = elem[index+1:]
					break
				}
			}
			gpus = append(gpus, strings.TrimSpace(elem))
		case strings.HasPrefix(elem, "USB"):
			for index, char := range elem {
				if char == ':' {
					elem = elem[index+1:]
					break
				}
			}
			usbs = append(usbs, strings.TrimSpace(elem))
		}
	}

	// return
	slice := make([]string, 0, len(gpus)+len(usbs))
	for i, elem := range gpus {
		slice = append(slice, fmt.Sprintf("GPU%d %s", i+1, elem))
	}
	for i, elem := range usbs {
		slice = append(slice, fmt.Sprintf("USB%d %s", i+1, elem))
	}
	return slice
}

func release() (name, version string) {
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
	return
}

func cat(fp string) string {
	data, _ := os.ReadFile(fp)
	return string(data)
}
