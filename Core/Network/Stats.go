package Network

import (
	"fmt"
	"log"
	"net"
)

func GetStats() {
	// Network Information
	Net, err := net.Interfaces()
	if err != nil {
		log.Fatalf("Error retrieving network info: %v", err)
	}
	fmt.Println("Network Interfaces:")
	for _, net := range Net {
		fmt.Printf("  Name: %s\n  MAC: %s\n  MTU: %d\n  Flags: %v\n\n", net.Name, net.HardwareAddr, net.MTU, net.Flags)
	}
}
