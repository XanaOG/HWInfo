package Memory

import (
	"fmt"
	"log"

	"github.com/shirou/gopsutil/mem"
)

func GetStats() {
	// Memory Information
	MemoryInfo, err := mem.VirtualMemory()
	if err != nil {
		log.Fatalf("Error retrieving memory info: %v", err)
	}
	fmt.Printf("Memory Info:\n  Total: %.2f GB\n  Used: %.2f GB\n  Free: %.2f GB\n\n",
		float64(MemoryInfo.Total)/1e9, float64(MemoryInfo.Used)/1e9, float64(MemoryInfo.Free)/1e9)
}
