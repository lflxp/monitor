package main

import (
	"encoding/json"
	"fmt"

	"github.com/lflxp/monitor/setup"
	// "github.com/shirou/gopsutil/net"
)

func main() {
	// fmt.Println(net.Interfaces())
	jsons, err := json.Marshal(setup.NewCommon())
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(jsons))
}
