package config

import (
	lxd "github.com/lxc/lxd/client"
	"github.com/lxc/lxd/lxc/config"
)

func GetDefaultInstanceServer() (lxd.InstanceServer, error) {
	configFile := NewConfigFileResolver(nil, nil).Resolve()

	loaded, err := config.LoadConfig(configFile)
	if err != nil {
		return nil, err
	}

	return loaded.GetInstanceServer(loaded.DefaultRemote)
}
