package connection

import (
	"github.com/pkg/errors"
)

type CompositeRemoteNameResolver struct {
	providers []RemoteNameResolver
}

func (r *CompositeRemoteNameResolver) Resolve() (string, error) {
	for _, provider := range r.providers {
		connection, err := provider.Resolve()
		if err != nil {
			return "", err
		}

		if connection != "" {
			return connection, nil
		}
	}

	return "", errors.New("Cannot found lxd.InstanceServer.")
}

func NewCompositeRemoteNameResolver(resolvers []RemoteNameResolver) RemoteNameResolver {
	return &CompositeRemoteNameResolver{
		providers: resolvers,
	}
}
