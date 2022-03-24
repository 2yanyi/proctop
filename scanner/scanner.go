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

func Processes() []*Process {
	slice := make([]*Process, 0, 500)
	processes, err := process.Processes()
	if err != nil {
		panic(err)
	}
	statistics := &Process{Name: StatisticsTag}
	for i := range processes {
		it := processInfo(processes[i])
		statistics.CPUPercent += it.CPUPercent
		statistics.MemoryBytes += it.MemoryBytes
		slice = append(slice, it)
	}
	slice = append(slice, statistics)
	return slice
}

func processInfo(proc *process.Process) *Process {
	it := &Process{Process: proc}
	it.Commandline, _ = proc.Cmdline()
	it.Name, _ = proc.Name()
	it.CPUPercent, _ = proc.CPUPercent()
	MemoryInfo, _ := proc.MemoryInfo()
	if MemoryInfo != nil {
		it.MemoryBytes = MemoryInfo.RSS
	}
	it.Count = 1
	return it
}
