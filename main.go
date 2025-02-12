package main

import (
	"hardware-pulse/models"

	"github.com/gookit/color"
)

func main() {
	var cpu *models.CPU = &models.CPU{}
	var mem *models.Memory = &models.Memory{}

	usage, err := cpu.CalculateCPUUsage()
	memErr := mem.GetMemoryStatus()

	if err != nil || memErr != nil {
		color.Red.Println("CPU Error:", err)
		color.Red.Println("Memory Error:", memErr)
	} else {
		color.Green.Printf("CPU Usage: %.2f%%\n", usage)

		color.Cyan.Printf("Memory Usage: %d%%\n", mem.MemoryLoad)
		color.Yellow.Printf("Total Memory: %d MB\n", mem.TotalPhys/1024/1024)
		color.Yellow.Printf("Free Memory: %d MB\n", mem.AvailPhys/1024/1024)
		color.Magenta.Printf("Total Virtual Memory: %d MB\n", mem.TotalVirtual/1024/1024)
		color.Magenta.Printf("Free Virtual Memory: %d MB\n", mem.AvailVirtual/1024/1024)
	}
}
