package scanner

import "github.com/shirou/gopsutil/v3/process"

type Process struct {
	Process     *process.Process
	Name        string
	CPUPercent  float64
	MemoryBytes uint64
	Count       int
}

func Processes() []Process {
	elements := make([]Process, 0, 100)
	processes, processErr := process.Processes()
	if processErr != nil {
		panic(processErr)
	}
	for _, proc := range processes {
		if ignore(proc) {
			continue
		}
		name, _ := proc.Name()
		CPUPercent, _ := proc.CPUPercent()
		MemoryInfo, _ := proc.MemoryInfo()

		elements = append(elements, Process{
			Process:    proc,
			Name:       name,
			CPUPercent: CPUPercent,
			//MemoryBytes: MemoryInfo.RSS + MemoryInfo.VMS + MemoryInfo.HWM + MemoryInfo.Data + MemoryInfo.Stack + MemoryInfo.Locked + MemoryInfo.Swap,
			MemoryBytes: MemoryInfo.RSS,
			Count:       1,
		})
	}
	return elements
}

func ignore(proc *process.Process) (_ bool) {
	cmdline, _ := proc.Cmdline()
	if cmdline == "" {
		return true
	}
	return
}
