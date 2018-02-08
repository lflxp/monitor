package main

import (
	"fmt"

	"github.com/lflxp/monitor/setup"
	// "github.com/shirou/gopsutil/net"
)

func main() {
	fmt.Println(setup.NewCommon())
	// fmt.Println(net.Interfaces())
}
