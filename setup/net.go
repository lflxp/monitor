package setup

import (
	"fmt"
	"net"
)

type Net struct {
	Ip string
}

func (this *Net) Get() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println(ipnet.IP.String())
				fmt.Println(ipnet.Mask.String(), ipnet.Network())
			}
		}
		fmt.Println(a.String(), a.Network())
	}
}

func NewNet() *Net {
	data := &Net{}
	data.Get()
	return data
}
