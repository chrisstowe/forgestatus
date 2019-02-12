package common

import (
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
)

// MemoryUsed returns the percentage of virtual memory used.
func MemoryUsed() float64 {
	v, err := mem.VirtualMemory()
	if err != nil {
		return 0.0
	}

	return v.UsedPercent
}

const cpuCheckDuration = 500 * time.Millisecond

// CpuUsed returns the percentage of cpu used.
func CpuUsed() float64 {
	c, err := cpu.Percent(cpuCheckDuration, false)
	if err != nil || len(c) < 1 {
		return 0.0
	}

	return c[0]
}

// DiskUsed returns the percentage of disk space used.
func DiskUsed() float64 {
	d, err := disk.Usage("/")
	if err != nil {
		return 0.0
	}

	return d.UsedPercent
}

// ProcsRunning returns the number of processes running.
func ProcsRunning() int {
	l, err := load.Misc()
	if err != nil {
		return 0.0
	}

	return l.ProcsRunning
}
