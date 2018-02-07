package main

import (
	"fmt"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/docker"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/process"
)

func main() {
	fmt.Println(host.Info())
	fmt.Println(cpu.Info())
	fmt.Println(cpu.Times(false))
	fmt.Println(disk.Partitions(true))
	fmt.Println(disk.IOCounters())
	data, err := disk.Partitions(true)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, x := range data {
		fmt.Println(disk.Usage(x.Mountpoint))
	}
	fmt.Println(docker.GetDockerStat())
	fmt.Println(load.Avg())
	fmt.Println(load.Misc())
	fmt.Println(mem.SwapMemory())
	fmt.Println(mem.VirtualMemory())
	fmt.Println(net.IOCounters(true))
	fmt.Println(net.FilterCounters())
	fmt.Println(net.ProtoCounters([]string{"https", "tcp", "udp"}))
	fmt.Println(process.Pids())
	fmt.Println(process.Processes())
	// fmt.Println(host.Uptime())
	// fmt.Println(host.Users())
	// fmt.Println(host.BootTime())
	// fmt.Println(host.PlatformInformation())
	// fmt.Println(host.Virtualization())
	// fmt.Println(host.KernelVersion())
}
