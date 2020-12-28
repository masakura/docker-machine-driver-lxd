package client

import "github.com/lxc/lxd/shared/api"

type SettingsBuilder struct {
	container *api.Container
}

func (b *SettingsBuilder) Config() *ConfigBuilder {
	return NewConfigBuilder(b.container.Config)
}

func (b *SettingsBuilder) Writable() api.ContainerPut {
	return b.container.Writable()
}

func NewSettingsBuilder(container *api.Container) *SettingsBuilder {
	return &SettingsBuilder{
		container: container,
	}
}
