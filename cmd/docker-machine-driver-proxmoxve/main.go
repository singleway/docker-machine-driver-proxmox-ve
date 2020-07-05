package main

import (
	"github.com/docker/machine/libmachine/drivers/plugin"
	proxmoxve "github.com/singleway/docker-machine-driver-proxmox-ve"
)

func main() {
	plugin.RegisterDriver(proxmoxve.NewDriver("default", ""))
}
