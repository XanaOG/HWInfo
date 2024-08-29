package Host

import (
	"fmt"
	"log"

	"github.com/shirou/gopsutil/host"
)

func GetStats() {
	// Host Information
	Host, err := host.Info()
	if err != nil {
		log.Fatalf("Error retrieving host info: %v", err)
	}
	fmt.Printf("Host Info:\n  Hostname: %s\n  OS: %s\n  Platform: %s\n  Kernel: %s\n\n",
		Host.Hostname, Host.OS, Host.Platform, Host.KernelVersion)
}
