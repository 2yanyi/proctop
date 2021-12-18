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
		var cmdline string
		var ppid int32
		if ignore(proc, &ppid, &cmdline) {
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

func ignore(proc *process.Process, ppidCp *int32, cmdlineCp *string) (_ bool) {
	cmdline, _ := proc.Cmdline()
	if cmdline == "" {
		return true
	}
	times, _ := proc.Times()
	if times == nil {
		return true
	}
	if times.Total() == 0 {
		return true
	}
	ppid, _ := proc.Ppid()
	if ppid == 0 {
		return true
	}
	prov, _ := process.NewProcess(ppid)
	if prov != nil {
		name, _ := prov.Name()
		if name == "systemd" {
			return true
		}
	}
	*ppidCp = ppid
	*cmdlineCp = cmdline
	return
}
