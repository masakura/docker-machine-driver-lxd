package main

import (
	"github.com/docker/machine/libmachine/drivers/plugin"
	"gitlab.com/masakura/docker-machine-driver-lxd/drivers/lxd"
	"gitlab.com/masakura/docker-machine-driver-lxd/lxd/socket"
)

func main() {
	err := socket.ConfigureEnvironmentVariables()
	if err != nil {
		panic(err)
	}

	plugin.RegisterDriver(lxd.NewDriver("", ""))
}
