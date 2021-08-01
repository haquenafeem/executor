package platform

import (
	"os"
	"runtime"
)

type Info struct {
	OS   string `json:"os"`
	Arch string `json:"arch"`
	Host string `json:"host"`
}

func Get() Info {
	hostname, _ := os.Hostname()
	return Info{
		OS:   runtime.GOOS,
		Arch: runtime.GOARCH,
		Host: hostname,
	}
}
