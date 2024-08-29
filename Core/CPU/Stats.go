package CPU

import (
	"fmt"
	"log"

	"github.com/shirou/gopsutil/cpu"
)

func GetStats() {
	// CPU Information
	CPU, err := cpu.Info()
	if err != nil {
		log.Fatalf("Error retrieving CPU info: %v", err)
	}
	fmt.Println("CPU Info:")
	for _, cpu := range CPU {
		fmt.Printf("  Model: %s\n  Cores: %d\n  Speed: %.2f GHz\n\n", cpu.ModelName, cpu.Cores, cpu.Mhz/1000)
	}
}
