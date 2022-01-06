package top

import (
	"bytes"
	"fmt"
	"os"
	"r/flagset"
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

func renameJava(name, commandline *string) {

	// Java: JetBrains
	if strings.Contains(*commandline, "-Didea.vendor.name=JetBrains") {
		for i, value := range strings.Split(*commandline, "-Didea.platform.prefix=") {
			if i == 0 {
				continue
			}
			*name = ""
			for _, char := range value {
				if char == ' ' {
					break
				}
				*name += string(char)
			}
			*name = javaTag + strings.ToLower(*name)
		}
		return
	}

	// Java: hadoop
	if strings.Contains(*commandline, "-Dhadoop.log") {
		*name = javaTag + "hadoop"
		return
	}

	// Java: hbase
	if strings.Contains(*commandline, "-Dhbase.log") {
		*name = javaTag + "hbase"
		return
	}

	// Java: zookeeper
	if strings.Contains(*commandline, "-Dzookeeper.log") {
		*name = javaTag + "zookeeper"
		return
	}

	// Java: kafka
	if strings.Contains(*commandline, "-Dkafka.log") {
		*name = javaTag + "kafka"
		return
	}

	// java: jar
	jars := strings.Count(*commandline, ".jar")
	if jars == 1 {
		values := strings.Split(*commandline, ".jar")[0]
		value := values[strings.LastIndex(values, "/")+1:]
		*name = javaTag + value + ".jar"
		return
	} else if jars > 1 {
		*name = *commandline
		return
	}

	values := strings.Fields(strings.ReplaceAll(*commandline, "=", " "))
	tag := values[len(values)-2]
	value := ""
	if tag[0] == '-' {
		value = values[len(values)-3]
		if value[0] == '/' {
			value = values[len(values)-1]
		}
	} else {
		value = values[len(values)-1]
	}
	*name = javaTag + value
}

// 多个进程合并
func rename(name, commandline *string) {

	if *flagset.Java {
		if *name == "java" {
			renameJava(name, commandline)
			return
		}
	}

	switch {

	// System
	case strings.HasPrefix(*name, "upstart"):
		*name = "upstart"
	case strings.HasPrefix(*name, "indicator"):
		*name = "indicator"
	case strings.HasPrefix(*name, "systemd"), *name == "(sd-pam)":
		*name = "systemd"
	case strings.HasPrefix(*name, "dbus"):
		*name = "dbus"
	case strings.HasPrefix(*name, "ibus"):
		*name = "ibus"
	case strings.HasPrefix(*name, "cups"):
		*name = "cups"
	case strings.HasPrefix(*name, "xdg"):
		*name = "xdg"
	case strings.HasPrefix(*name, "fcitx"):
		*name = "fcitx"
	case strings.HasPrefix(*name, "evolution"):
		*name = "evolution"
	case strings.HasPrefix(*name, "pipewire"):
		*name = "pipewire"
	case strings.HasPrefix(*name, "unity"):
		*name = "unity-tools"

	// GNOME
	case strings.HasPrefix(*name, "tracker"):
		*name = "tracker"
	case strings.HasPrefix(*name, "gvfs"):
		*name = "gvfs"
	case strings.HasPrefix(*name, "gdm"):
		*name = "gdm"
	case strings.HasPrefix(*name, "gsd"):
		*name = "gsd"
	case strings.HasPrefix(*name, "goa"):
		*name = "goa"
	case strings.HasPrefix(*name, "at-spi"):
		*name = "at-spi"
	case strings.HasPrefix(*name, "gnome"):
		{
			switch *name {
			case "gnome-terminal", "gnome-terminal.real":
				*name = "terminal"
			case "gnome-disks":
				*name = "disks"
			default:
				*name = "gnome-shell"
			}
		}

	// Database
	case strings.HasPrefix(*name, "clickhouse"):
		*name = "clickhouse"
	case strings.HasPrefix(*name, "mongo"):
		*name = "mongodb"
	case strings.HasPrefix(*name, "mysql"):
		*name = "mysql"
	case strings.HasPrefix(*name, "redis"):
		*name = "redis"

	// Applications
	case strings.HasPrefix(*name, "chrome"):
		*name = "chrome"
	case strings.HasPrefix(*name, "sysproxy-cmd"):
		*name = "lantern"
	case strings.HasPrefix(*name, "docker"):
		*name = "docker"
	case strings.HasPrefix(*name, "PM2"):
		*name = "PM2"
	case strings.HasPrefix(*name, "VBox"), *name == "VirtualBoxVM":
		if *name != "VBoxClient" {
			*name = "VirtualBoxVM"
		}
	case strings.HasPrefix(*name, "virt"), *name == "libvirtd":
		*name = "virt-manager"
	}

	// path > name
	if strings.HasPrefix(*name, "/") {
		tmp := *name
		*name = tmp[strings.LastIndex(tmp, "/")+1:]
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
			homepage.WebsiteMatch(strings.ToLower(proc.Name), &proc.CPUPercent),
		)
		if proc.Name == scanner.StatisticsTag {
			page.WriteString(strings.Join([]string{"\u001B[0;37;48m", buf, "\u001B[0m\n"}, ""))
		} else if cpu == "0.0" {
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

func nameFormat(s string) string {
	if strings.HasPrefix(s, javaTag) {
		return "                            \u001B[0;32;48m" + s
	}
	if len(s) > 32 {
		s = s[:30] + ".."
	}
	return s
}

const javaTag = "java:"

const (
	_KB = 1024
	_MB = 1024 * 1024
	_GB = 1024 * 1024 * 1024
)

var ProcessId = int32(os.Getpid())
