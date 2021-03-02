package config

import "github.com/lxc/lxd/lxc/config"

type Loader struct {
	resolver ConfigFileResolver
}

func (c *Loader) Load() (*config.Config, error) {
	file := c.resolver.Resolve()
	return config.LoadConfig(file)
}

func NewConfig() *Loader {
	return &Loader{
		resolver: NewConfigFileResolver(nil, nil),
	}
}
