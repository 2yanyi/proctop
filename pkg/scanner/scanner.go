package scanner

import (
	"github.com/shirou/gopsutil/v3/process"
)

const StatisticsTag = "[statistics]"

type Process struct {
	Ppid          int32           // Ppid
	Commandline   string          // Commandline
	Name          string          // Name
	Status        map[string]byte // Status
	CPUPercent    float64         // CPUPercent
	MemoryBytes   uint64          // MemoryBytes
	NumThreads    int32           // NumThreads
	NumFDs        int32           // NumFDs
	FIOReadBytes  uint64          // FIOReadBytes
	FIOWriteBytes uint64          // FIOWriteBytes
	Count         int             // Count
}

func Processes() []*Process {
	slice := make([]*Process, 0, 500)
	processes, err := process.Processes()
	if err != nil {
		panic(err)
	}
	statistics := &Process{Name: StatisticsTag}
	for i := 0; i < len(processes); i++ {
		it := processInfo(processes[i])
		statistics.CPUPercent += it.CPUPercent
		statistics.MemoryBytes += it.MemoryBytes
		slice = append(slice, it)
	}
	slice = append(slice, statistics)
	return slice
}

func processInfo(proc *process.Process) *Process {
	it := &Process{}
	it.Ppid, _ = proc.PpidWithContext(nil)
	it.Commandline, _ = proc.CmdlineWithContext(nil)
	it.Name, _ = proc.NameWithContext(nil)
	it.CPUPercent, _ = proc.CPUPercentWithContext(nil)
	it.NumThreads, _ = proc.NumThreads()
	it.NumFDs, _ = proc.NumFDs()

	// MemoryBytes
	MemoryInfo, _ := proc.MemoryInfoWithContext(nil)
	if MemoryInfo != nil {
		it.MemoryBytes = MemoryInfo.RSS
	}

	// Status
	if it.Status == nil {
		it.Status = make(map[string]byte)
	}
	status, _ := proc.StatusWithContext(nil)
	if len(status) != 0 {
		for i := 0; i < len(status); i++ {
			it.Status[status[i]] = 0
		}
	}

	// FIO
	IOCounters, _ := proc.IOCounters()
	if IOCounters != nil {
		it.FIOReadBytes = IOCounters.ReadBytes - prevPage[it.Name].ReadBytes
		it.FIOWriteBytes = IOCounters.WriteBytes - prevPage[it.Name].WriteBytes
		prevPage[it.Name] = struct{ ReadBytes, WriteBytes uint64 }{IOCounters.ReadBytes, IOCounters.WriteBytes}
	}

	it.Count = 1
	return it
}

var prevPage = make(map[string]struct {
	ReadBytes, WriteBytes uint64
})
