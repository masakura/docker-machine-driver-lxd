package connection

import (
	"github.com/lxc/lxd/lxc/config"
	"gitlab.com/masakura/docker-machine-driver-lxd/drivers/lxd/options"
)

type DefaultRemoteNameResolver struct {
	resolver RemoteNameResolver
}

func (r *DefaultRemoteNameResolver) Resolve() (string, error) {
	return r.resolver.Resolve()
}

func NewDefaultRemoteNameResolver(options options.Options, config *config.Config) RemoteNameResolver {
	return &DefaultRemoteNameResolver{
		resolver: NewCompositeRemoteNameResolver([]RemoteNameResolver{
			NewDriverOptionsProvider(options),
			NewConfigRemoteNameResolver(config),
		}),
	}
}
