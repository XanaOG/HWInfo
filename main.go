package main

import (
	CPU "github.com/XanaOG/HWInfo/Core/CPU"
	Disk "github.com/XanaOG/HWInfo/Core/Disk"
	Host "github.com/XanaOG/HWInfo/Core/Host"
	Memory "github.com/XanaOG/HWInfo/Core/Memory"
	Network "github.com/XanaOG/HWInfo/Core/Network"
)

func main() {
	CPU.GetStats()
	Memory.GetStats()
	Disk.GetStats()
	Host.GetStats()
	Network.GetStats()
}
