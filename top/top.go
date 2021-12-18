package top

import (
	"bytes"
	"fmt"
	"os"
	"r/scanner"
	"sort"
)

func Call(limit int) []byte {
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
	page := &bytes.Buffer{}
	for i, proc := range processes {
		if i > limit {
			break
		}
		// Num wc Mem CPU% Name
		page.WriteString(fmt.Sprintf("%3d) %2d %7s %6s  %s\n", i+1,
			proc.Count,
			sizeFormat(float64(proc.MemoryBytes)),
			fmt.Sprintf("%.1f", proc.CPUPercent),
			proc.Name))
	}
	return page.Bytes()
}

func combineSimilarItem(processes *[]scanner.Process) {
	sets := make(map[string]scanner.Process)
	for _, elem := range *processes {
		if elem.Process.Pid == PIDSelf {
			continue
		}
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
		*processes = append(*processes, proc)
	}
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

const (
	_KB = 1024
	_MB = 1024 * 1024
	_GB = 1024 * 1024 * 1024
)

var PIDSelf = int32(os.Getpid())
