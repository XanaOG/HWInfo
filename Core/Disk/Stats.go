package Disk

import (
	"fmt"
	"log"

	"github.com/shirou/gopsutil/disk"
)

func GetStats() {
	Partitions, err := disk.Partitions(true)
	if err != nil {
		log.Fatalf("Error retrieving disk partitions: %v", err)
	}

	fmt.Println("Disk Info:")
	for _, partition := range Partitions {
		diskInfo, err := disk.Usage(partition.Mountpoint)
		if err != nil {
			log.Printf("Error retrieving disk info for %s: %v", partition.Mountpoint, err)
			continue
		}
		fmt.Printf("  Partition: %s\n  Total: %.2f GB\n  Used: %.2f GB\n  Free: %.2f GB\n\n",
			partition.Mountpoint, float64(diskInfo.Total)/1e9, float64(diskInfo.Used)/1e9, float64(diskInfo.Free)/1e9)
	}
}
