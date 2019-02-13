package common

import (
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
)

// CheckMemoryUsed returns the percentage of virtual memory used.
func CheckMemoryUsed() float64 {
	v, err := mem.VirtualMemory()
	if err != nil {
		return 0.0
	}

	return v.UsedPercent
}

const cpuCheckDuration = 500 * time.Millisecond

// CheckCPUUsed returns the percentage of cpu used.
func CheckCPUUsed() float64 {
	c, err := cpu.Percent(cpuCheckDuration, false)
	if err != nil || len(c) < 1 {
		return 0.0
	}

	return c[0]
}

// CheckDiskUsed returns the percentage of disk space used.
func CheckDiskUsed() float64 {
	d, err := disk.Usage("/")
	if err != nil {
		return 0.0
	}

	return d.UsedPercent
}

// CheckProcsRunning returns the number of processes running.
func CheckProcsRunning() int {
	l, err := load.Misc()
	if err != nil {
		return 0.0
	}

	return l.ProcsRunning
}
