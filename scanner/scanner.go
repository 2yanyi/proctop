package scanner

import "github.com/shirou/gopsutil/v3/process"

const StatisticsTag = "[statistics]"

type Process struct {
	Process     *process.Process
	Commandline string
	Name        string
	CPUPercent  float64
	MemoryBytes uint64
	Count       int
}

func Processes() []Process {
	elements := make([]Process, 0, 500)
	processes, processErr := process.Processes()
	if processErr != nil {
		panic(processErr)
	}
	statistics := Process{Name: StatisticsTag}
	for _, proc := range processes {
		elem, has := ignore(proc)
		if has {
			continue
		}
		statistics.CPUPercent += elem.CPUPercent
		statistics.MemoryBytes += elem.MemoryBytes
		elements = append(elements, elem)
	}
	elements = append(elements, statistics)
	return elements
}

func ignore(proc *process.Process) (elem Process, _ bool) {
	elem.Commandline, _ = proc.Cmdline()
	if elem.Commandline == "" {
		return elem, true
	}
	elem.Process = proc
	elem.Name, _ = proc.Name()
	elem.CPUPercent, _ = proc.CPUPercent()
	MemoryInfo, _ := proc.MemoryInfo()
	if MemoryInfo != nil {
		elem.MemoryBytes = MemoryInfo.RSS
	}
	elem.Count = 1
	return
}
