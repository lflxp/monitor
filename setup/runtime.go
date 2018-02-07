package setup

import (
	"runtime"
)

type Runtime struct {
	Goarch   string
	Goos     string
	Compiler string
}

func (this *Runtime) SetInfo() {
	this.Goarch = runtime.GOARCH
	this.Goos = runtime.GOOS
	this.Compiler = runtime.Compiler
}

func NewRuntime() *Runtime {
	data := &Runtime{}
	data.SetInfo()
	return data
}
