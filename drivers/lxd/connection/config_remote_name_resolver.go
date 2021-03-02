package connection

import (
	lxdConfig "github.com/lxc/lxd/lxc/config"
)

type ConfigRemoteNameResolver struct {
	config *lxdConfig.Config
}

func (r *ConfigRemoteNameResolver) Resolve() (string, error) {
	return r.config.DefaultRemote, nil
}

func NewConfigRemoteNameResolver(config *lxdConfig.Config) RemoteNameResolver {
	return &ConfigRemoteNameResolver{
		config: config,
	}
}
