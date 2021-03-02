package connection

import (
	options2 "gitlab.com/masakura/docker-machine-driver-lxd/drivers/lxd/options"
)

type DriverOptionsRemoteNameResolver struct {
	options options2.Options
}

func (r *DriverOptionsRemoteNameResolver) Resolve() (string, error) {
	return r.options.Remote, nil
}

func NewDriverOptionsProvider(options options2.Options) RemoteNameResolver {
	return &DriverOptionsRemoteNameResolver{options: options}
}
