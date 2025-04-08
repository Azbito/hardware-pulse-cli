package main

import (
	"hardware-pulse/display"
	"hardware-pulse/input"
	"hardware-pulse/models"
)

func main() {
	cpu := &models.CPU{}
	mem := &models.Memory{}

	input.Init()
	defer input.Restore()

	display.PrintHeader()
	display.PrintSystemInfo(cpu, mem)

	for {
		key := input.ReadKey()

		switch key {
		case ' ':
			display.Clean()
			display.PrintHeader()
			display.PrintSystemInfo(cpu, mem)
		case 'q', 'Q':
			return
		}
	}
}
