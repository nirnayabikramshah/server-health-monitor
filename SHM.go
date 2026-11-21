// -- Making an Health Monitor -- //
package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/mem"
)

func main() {
	CPUpercent, err := cpu.Percent(time.Second, false)
	if err != nil {
		fmt.Println("Error getting CPU stats:", err)
	}
	fmt.Printf("CPU Usage: %.2f%%\n", CPUpercent[0])
	// -- For the CPU -- //
	v, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println("Error getting RAM stats:", err)
	}
	fmt.Printf("RAM Usage: %.2f%%(Used %v MB)\n ", v.UsedPercent, v.Used/1024/1024)
	// -- For the RAM -- //
	d, err := disk.Usage("C:")
	if err != nil {
		fmt.Println("Error getting  Disk stats:", err)
	}
	fmt.Printf("Disk Usage :%.2f%% ", d.UsedPercent)
	// -- For the Disk -- //
}
