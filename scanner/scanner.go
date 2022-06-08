package scanner

import (
	"github.com/shirou/gopsutil/v3/process"
	"strings"
)

const StatisticsTag = "[statistics]"

type Process struct {
	Ppid          int32   // Ppid
	Commandline   string  // Commandline
	Name          string  // Name
	Status        string  // Status
	CPUPercent    float64 // CPUPercent
	MemoryBytes   uint64  // MemoryBytes
	NumThreads    int32   // NumThreads
	NumFDs        int32   // NumFDs
	FIOReadBytes  uint64  // FIOReadBytes
	FIOWriteBytes uint64  // FIOWriteBytes
	Count         int     // Count
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
	status, _ := proc.StatusWithContext(nil)
	if len(status) != 0 {
		it.Status = strings.Join(status, ",")
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

/*

// Running 正在运行，或者正在等待运行。
// Blocked 短期锁定，等待一个短暂的、不间断的操作（通常是 I/O）。
// Idle    空闲状态，不可中断睡眠的内核线程，任务睡眠超过约 20 秒。
// Lock    等待解锁，
// Sleep   睡眠状态，可中断状态睡眠，等待短暂的、可中断的操作。
// Stop    暂停状态，进程处于暂停或者跟踪状态。
// Wait    空闲中断，
// Zombie  僵尸进程，已终止但未被其父进程回收。

// Solaris states. See https://github.com/collectd/collectd/blob/1da3305c10c8ff9a63081284cf3d4bb0f6daffd8/src/processes.c#L2115
Daemon   = "daemon"
Detached = "detached"
System   = "system"
Orphan   = "orphan"

UnknownState = ""

*/
