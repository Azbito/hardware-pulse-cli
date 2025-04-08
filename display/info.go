package display

import (
	"fmt"
	"hardware-pulse/models"
	"hardware-pulse/utils"

	"github.com/gookit/color"
)

func Clean() {
		fmt.Print("\033[H\033[2J")
}

func PrintHeader() {
	asciiArt := `
         _       _    _                   _          _            _             _                   _           _
        / /\    / /\ / /\                /\ \       /\ \         / /\      _   / /\                /\ \        /\ \
       / / /   / / // /  \              /  \ \     /  \ \____   / / /    / /\ / /  \              /  \ \      /  \ \
      / /_/   / / // / /\ \            / /\ \ \   / /\ \_____\ / / /    / / // / /\ \            / /\ \ \    / /\ \ \
     / /\ \__/ / // / /\ \ \          / / /\ \_\ / / /\/___  // / /_   / / // / /\ \ \          / / /\ \_\  / / /\ \_\
    / /\ \___\/ // / /  \ \ \        / / /_/ / // / /   / / // /_//_/\/ / // / /  \ \ \        / / /_/ / / / /_/_ \/_/
   / / /\/___/ // / /___/ /\ \      / / /__\/ // / /   / / // _______/\/ // / /___/ /\ \      / / /__\/ / / /____/\
  / / /   / / // / /_____/ /\ \    / / /_____// / /   / / // /  \____\  // / /_____/ /\ \    / / /_____/ / /\____\/
 / / /   / / // /_________/\ \ \  / / /\ \ \  \ \ \__/ / //_/ /\ \ /\ \// /_________/\ \ \  / / /\ \ \  / / /______
/ / /   / / // / /_       __\ \_\/ / /  \ \ \  \ \___\/ / \_\//_/ /_/ // / /_       __\ \_\/ / /  \ \ \/ / /_______\
\/_/    \/_/ \_\___\     /____/_/\/_/    \_\/   \/_____/      \_\/\_\/ \_\___\     /____/_/\/_/    \_\/\/__________/

Press SPACE to run once again
	`

	color.Cyan.Println(asciiArt)
	color.Red.Println("Press Q to quit")
}


func PrintSystemInfo(cpu *models.CPU, mem *models.Memory) {
	usage, err := cpu.CalculateCPUUsage()
	usageTime := cpu.GetUptime()
	memErr := mem.GetMemoryStatus()

	if err != nil || memErr != nil {
		color.Red.Println("⚠️ Error while gathering system informations")
		if err != nil {
			color.Red.Println("CPU Error:", err)
		}
		if memErr != nil {
			color.Red.Println("Memory Error:", memErr)
		}
		return
	}

	color.Yellow.Printf("⌛ You've been using your computer for: %s\n", utils.FormatDuration(usageTime))
	color.Green.Printf("🖥️  CPU Usage: %.2f%%\n", usage)
	fmt.Println()

	color.Cyan.Println("🗄️  Physical memory:")
	color.Yellow.Printf(" ⋈ Total: %s\n", utils.FormatBytes(mem.TotalPhys))
	color.Yellow.Printf(" ⋈ Free: %s\n", utils.FormatBytes(mem.AvailPhys))
	color.Cyan.Printf(" ⋈ Usage:   %d%%\n", mem.MemoryLoad)

	fmt.Println()

	color.Magenta.Println("📦  Virtual memory:")
	color.Magenta.Printf("  ⋈ Total: %s\n", utils.FormatBytes(mem.TotalVirtual))
	color.Magenta.Printf("  ⋈ Free: %s\n", utils.FormatBytes(mem.AvailVirtual))
}
