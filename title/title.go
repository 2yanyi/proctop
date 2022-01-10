package title

import (
	"fmt"
	"os"
	"os/exec"
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
	for _, line := range strings.Split(cat("/proc/cpuinfo"), "\n") {
		if strings.HasPrefix(line, "model name") {
			modelName = line
		} else if strings.HasPrefix(line, "physical id") {
			if line == "" {
				continue
			}
			cpus[line] = append(cpus[line], modelName)
		}
	}

	// return
	slice := make([]string, 0, len(cpus))
	for i := range cpus {
		elem := fmt.Sprintf("%s %d*Thread\n", cpus[i][0], len(cpus[i]))
		for index, char := range elem {
			if char == ':' {
				elem = elem[index+1:]
				break
			}
		}
		slice = append(slice, strings.TrimSpace(elem))
	}
	return slice
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
