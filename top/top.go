package top

import (
	"bytes"
	"fmt"
	"os"
	"r/scanner"
	"sort"
	"strings"
)

func Call(limit int) string {
	limit = limit - 1

	// 读取进程列表
	processes := scanner.Processes()

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
	for _, elem := range *processes {
		if elem.Process.Pid == PIDSelf {
			continue
		}
		rename(&elem.Name)
		if proc, has := sets[elem.Name]; has {
			proc.Count++
			proc.CPUPercent += elem.CPUPercent
			proc.MemoryBytes += elem.MemoryBytes
			sets[elem.Name] = proc
		} else {
			sets[elem.Name] = elem
		}
	}
	*processes = nil
	for _, proc := range sets {
		if proc.CPUPercent < 0.1 {
			if strings.HasSuffix(proc.Name, ".sh") ||
				proc.Name == "gnome-shell" ||
				proc.Name == "snapd" ||
				proc.Name == "PM2" ||
				proc.Name == "dockerd" ||
				proc.Name == "sshd" {
				proc.CPUPercent = 0.1
			}
		}
		*processes = append(*processes, proc)
	}
}

func rename(name *string) {
	switch {
	case strings.HasPrefix(*name, "systemd"):
		*name = "systemd"
	case strings.HasPrefix(*name, "ibus"):
		*name = "ibus"
	case strings.HasPrefix(*name, "dbus"):
		*name = "dbus"
	case strings.HasPrefix(*name, "cups"):
		*name = "cups"
	case strings.HasPrefix(*name, "xdg"):
		*name = "xdg"
	case strings.HasPrefix(*name, "evolution"):
		*name = "evolution"
	case strings.HasPrefix(*name, "pipewire"):
		*name = "pipewire"
	case strings.HasPrefix(*name, "gnome"):
		*name = "gnome-shell"
	case strings.HasPrefix(*name, "gvfs"):
		*name = "gvfs"
	case strings.HasPrefix(*name, "gsd"):
		*name = "gsd"
	case strings.HasPrefix(*name, "gdm"):
		*name = "gdm"
	case strings.HasPrefix(*name, "goa"):
		*name = "goa"
	case strings.HasPrefix(*name, "at-spi"):
		*name = "at-spi"
	case strings.HasPrefix(*name, "VBox"):
		*name = "VirtualBoxVM"
	case strings.HasPrefix(*name, "chrome"):
		*name = "chrome"
	case strings.HasPrefix(*name, "sysproxy-cmd"):
		*name = "lantern"
	case strings.HasPrefix(*name, "PM2"):
		*name = "PM2"

		// python3
	case strings.HasSuffix(*name, "python3"):
		*name = "python3"
	}
}

func fillScreen(processes []scanner.Process, limit int) (page bytes.Buffer) {
	for i, proc := range processes {
		if i > limit {
			break
		}
		cpu := fmt.Sprintf("%.1f", proc.CPUPercent)
		buf := fmt.Sprintf("%3d)  %2d  %7s  %32s  %6s  %s",
			i+1,                                   // Num
			proc.Count,                            // count
			sizeFormat(float64(proc.MemoryBytes)), // Memory
			nameFormat(proc.Name),                 // Name
			cpu,                                   // CPU
			descriptionMatch[strings.ToLower(proc.Name)],
		)
		if cpu == "0.0" {
			page.WriteString(strings.Join([]string{"\u001B[0;34;48m", buf, "\u001B[0m\n"}, ""))
		} else if cpu == "0.1" {
			page.WriteString(strings.Join([]string{"\u001B[0;36;48m", buf, "\u001B[0m\n"}, ""))
		} else if len(cpu) >= 5 {
			page.WriteString(strings.Join([]string{"\u001B[0;31;48m", buf, "\u001B[0m\n"}, ""))
		} else {
			page.WriteString(strings.Join([]string{"\u001B[0;33;48m", buf, "\u001B[0m\n"}, ""))
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

func nameFormat(s string) (_ string) {
	if len(s) > 32 {
		return s[:30] + ".."
	}
	return s
}

const (
	_KB = 1024
	_MB = 1024 * 1024
	_GB = 1024 * 1024 * 1024
)

var PIDSelf = int32(os.Getpid())
