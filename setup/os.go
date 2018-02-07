package setup

import (
	"os"
)

type Os struct {
	Hostname string
}

func (this *Os) GetHostname() {
	var err error
	this.Hostname, err = os.Hostname()
	if err != nil {
		this.Hostname = err.Error()
	}
}

func NewOs() *Os {
	data := &Os{}
	data.GetHostname()
	return data
}
