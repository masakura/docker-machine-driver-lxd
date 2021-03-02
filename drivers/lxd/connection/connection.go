package connection

import (
	lxd "github.com/lxc/lxd/client"
	"gitlab.com/masakura/docker-machine-driver-lxd/drivers/lxd/options"
	"gitlab.com/masakura/docker-machine-driver-lxd/lxd/config"
)

func GetConnection(options options.Options) (lxd.InstanceServer, error) {
	c, err := config.NewConfig().Load()
	if err != nil {
		return nil, err
	}

	remote, err := NewDefaultRemoteNameResolver(options, c).Resolve()
	if err != nil {
		return nil, err
	}

	return c.GetInstanceServer(remote)
}
