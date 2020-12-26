package main

import (
	"github.com/docker/machine/libmachine/drivers/plugin"
	"gitlab.com/masakura/docker-machine-driver-lxd/drivers/lxd"
)

func main() {
	plugin.RegisterDriver(lxd.NewDriver())
}
