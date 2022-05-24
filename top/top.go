package top

import (
	"bytes"
	"fmt"
	"github.com/matsuwin/cat"
	"r/colors"
	"r/data/variable"
	"r/scanner"
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
		if processes[i] == nil || processes[j] == nil {
			return false
		}
		return processes[i].CPUPercent > processes[j].CPUPercent
	})

	// 填充画面
	page := fillScreen(processes, limit)

	return page.String()
}

func combineSimilarItem(processes *[]*scanner.Process) {
	set := make(map[string]*scanner.Process)
	for _, proc := range *processes {
		if ignore(proc) {
			continue
		}
		if !variable.IsWin {
			switch proc.Name {
			case "sudo", "su", "sh", "bash":
				continue
			}
			if strings.HasSuffix(proc.Name, ".sh") && proc.CPUPercent < 0.1 {
				continue
			}
		}
		rename(&proc.Name, &proc.Commandline)
		// if strings.HasSuffix(proc.Name, " >") {continue}

		if elem, has := set[proc.Name]; has {
			proc.Count += elem.Count
			proc.CPUPercent += elem.CPUPercent
			proc.MemoryBytes += elem.MemoryBytes
			proc.NumThreads += elem.NumThreads
			proc.NumFDs += elem.NumFDs
			proc.FIOReadBytes += elem.FIOReadBytes
			proc.FIOWriteBytes += elem.FIOWriteBytes
			set[proc.Name] = proc
		} else {
			set[proc.Name] = proc
		}
	}
	*processes = nil
	for _, proc := range set {
		if proc.CPUPercent < 0.1 {
			if proc.Name == "sshd" {
				proc.CPUPercent = 0.1
			}
		}
		*processes = append(*processes, proc)
	}
}

func ignore(proc *scanner.Process) bool {
	if proc == nil {
		return true
	}
	return false
}

func fillScreen(processes []*scanner.Process, limit int) (page bytes.Buffer) {
	for i, proc := range processes {
		if proc == nil {
			continue
		}
		if i > limit {
			break
		}
		ioState := fmt.Sprintf("%7s/%s", cat.SizeFormat(float64(proc.FIOReadBytes)), cat.SizeFormat(float64(proc.FIOWriteBytes)))
		_ioState := strings.TrimSpace(ioState)
		if _ioState == "/" {
			ioState = ""
		} else {
			switch {
			case strings.HasPrefix(_ioState, "/"):
				ioState = "0" + _ioState
			case strings.HasSuffix(_ioState, "/"):
				ioState += "0"
			}
		}

		cpu := fmt.Sprintf("%.1f", proc.CPUPercent)
		buf := fmt.Sprintf("%3d)  %7d  %3d  %7s  %32s  %6s  %3d  %5d  %14s  %7s  %s",
			i,          // Num
			proc.Ppid,  // PPID
			proc.Count, // count
			cat.SizeFormat(float64(proc.MemoryBytes)), // Memory
			nameFormat(proc.Name),                     // Name
			cpu,                                       // CPU
			proc.NumThreads,
			proc.NumFDs,
			ioState,
			proc.Status, // status
			websiteFormat(strings.ToLower(proc.Name), &proc.CPUPercent),
		)
		if proc.Name == scanner.StatisticsTag {
			page.WriteString(colors.White(buf, colors.Zero) + "\n")
		} else if cpu == "0.0" {
			page.WriteString(colors.Blue(buf, colors.Zero) + "\n")
		} else if cpu == "0.1" {
			page.WriteString(colors.Cyan(buf, colors.Zero) + "\n")
		} else if len(cpu) >= 5 {
			page.WriteString(colors.Red(buf, colors.Zero) + "\n")
		} else {
			page.WriteString(colors.Yellow(buf, colors.Zero) + "\n")
		}
	}
	return
}

func nameFormat(s string) string {
	if strings.HasPrefix(s, javaTag) {
		return s
	}
	if len(s) > 32 {
		s = s[:30] + ".."
	}
	return s
}

func websiteFormat(s string, cpu *float64) (_ string) {
	if s == scanner.StatisticsTag {
		return fmt.Sprintf("%.2f%%  -- %s --", *cpu/cpuMax*100, loadAverage())
	}
	return Components[strings.TrimSuffix(s, " >")]
}

func loadAverage() (_ string) {
	text := cat.CommandArgs("", []string{"uptime"})
	average := strings.Split(text, "load average:")
	if len(average) != 2 {
		return
	}
	values := strings.Split(average[1], ",")
	if len(values) != 3 {
		return
	}
	return fmt.Sprintf("LOAD AVERAGE:%s/1m,%s/10m,%s/15m", values[0], values[1], values[2])
}

const javaTag = "J/"

var cpuMax = float64(runtime.NumCPU() * 100)
