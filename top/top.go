package top

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"r/colors"
	"r/scanner"
	"r/top/homepage"
	"runtime"
	"sort"
	"strings"
)

func Call(limit int) string {

	// 读取进程列表
	processes := scanner.Processes()
	if processes == nil {
		return "nil"
	}

	// 合并同类项
	combineSimilarItem(&processes)

	// 排序
	sort.Slice(processes, func(i, j int) bool {
		if processes[i].CPUPercent > processes[j].CPUPercent {
			return true
		}
		return false
	})

	// 填充画面
	page := fillScreen(processes, limit)

	return page.String()
}

func combineSimilarItem(processes *[]scanner.Process) {
	sets := make(map[string]scanner.Process)
	for _, proc := range *processes {
		if proc.Process != nil {
			if proc.Process.Pid == ProcessId {
				continue
			}
		}
		if runtime.GOOS == "linux" {
			switch proc.Name {
			case "sudo", "su", "sh", "bash":
				continue
			}
			rename(&proc.Name, &proc.Commandline)
		}
		if elem, has := sets[proc.Name]; has {
			proc.Count += elem.Count
			proc.CPUPercent += elem.CPUPercent
			proc.MemoryBytes += elem.MemoryBytes
			sets[proc.Name] = proc
		} else {
			sets[proc.Name] = proc
		}
	}
	*processes = nil
	for _, proc := range sets {
		if proc.CPUPercent < 0.1 {
			if proc.Name == "sshd" {
				proc.CPUPercent = 0.1
			}
		}
		*processes = append(*processes, proc)
	}
}

func fillScreen(processes []scanner.Process, limit int) (page bytes.Buffer) {
	for i, proc := range processes {
		if i > limit {
			break
		}

		cpu := fmt.Sprintf("%.1f", proc.CPUPercent)
		buf := fmt.Sprintf("%3d)  %2d  %7s  %32s  %6s  %s",
			i,                                     // Num
			proc.Count,                            // count
			sizeFormat(float64(proc.MemoryBytes)), // Memory
			nameFormat(proc.Name),                 // Name
			cpu,                                   // CPU
			cpuFormat(strings.ToLower(proc.Name), &proc.CPUPercent),
		)
		if proc.Name == scanner.StatisticsTag {
			page.WriteString(colors.White(buf) + "\n")
		} else if cpu == "0.0" {
			page.WriteString(colors.Blue(buf) + "\n")
		} else if cpu == "0.1" {
			page.WriteString(colors.Cyan(buf) + "\n")
		} else if len(cpu) >= 5 {
			page.WriteString(colors.Red(buf) + "\n")
		} else {
			page.WriteString(colors.Yellow(buf) + "\n")
		}
	}
	return
}

func sizeFormat(bytes float64) (_ string) {
	if bytes >= _GB {
		return fmt.Sprintf("%.1fG", bytes/1024/1024/1024)
	} else if bytes >= _MB {
		return fmt.Sprintf("%.1fM", bytes/1024/1024)
	} else if bytes >= _KB {
		return fmt.Sprintf("%.1fK", bytes/1024)
	}
	return
}

func nameFormat(s string) string {
	if strings.HasPrefix(s, javaTag) {
		return "                            " + colors.Green(s)
	}
	if len(s) > 32 {
		s = s[:30] + ".."
	}
	return s
}

func cpuFormat(s string, cpu *float64) (_ string) {
	if s == scanner.StatisticsTag {
		return fmt.Sprintf("%.2f%%  %s", *cpu/cpuMax*100, cpuTemperature())
	}
	if homepage.Coreutils[s] {
		return "coreutils"
	}
	if homepage.UtilLinux[s] {
		return "util-linux"
	}
	return homepage.Components[s]
}

func cpuTemperature() (C string) {
	output, _ := exec.Command("sensors").Output()
	if len(output) == 0 {
		output, _ = exec.Command("vcgencmd", "measure_temp").Output()
		if len(output) != 0 {
			text := strings.TrimSpace(string(output))
			i := strings.Index(text, "=")
			C = text[i+1:]
			if C[0] != '-' {
				C = "+" + text[i+1:]
			}
		} else {
			return "(need to install lm_sensors)"
		}
	} else {
		for _, line := range strings.Split(string(output), "\n") {
			if strings.HasPrefix(line, "Core") {
				values := strings.Fields(line)
				if len(values) < 2 {
					continue
				}
				C = values[2]
				break
			}
		}
	}
	if len(C) <= 1 {
		return
	}
	N := C[1]
	switch {
	case N >= '5':
		return strings.Join([]string{"\u001B[1;31;47m ", C, " \u001B[0m"}, "")
	case N >= '7':
		return strings.Join([]string{"\u001B[1;37;41m ", C, " \u001B[0m"}, "")
	}
	return strings.Join([]string{"\u001B[1;34;47m ", C, " \u001B[0m"}, "")
}

const javaTag = "java:"

const (
	_KB = 1024
	_MB = 1024 * 1024
	_GB = 1024 * 1024 * 1024
)

var cpuMax = float64(runtime.NumCPU() * 100)
var ProcessId = int32(os.Getpid())
