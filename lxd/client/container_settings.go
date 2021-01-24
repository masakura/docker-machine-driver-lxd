package client

import "github.com/lxc/lxd/shared/api"

type ContainerSettings struct {
	container *api.Container
}

func (b *ContainerSettings) Config() *ContainerConfig {
	return NewConfigBuilder(b.container.Config)
}

func (b *ContainerSettings) Devices() *ContainerDevices {
	return NewContainerDevices(b.container.Devices)
}

func (b *ContainerSettings) Writable() api.ContainerPut {
	return b.container.Writable()
}

func NewContainerSettings(container *api.Container) *ContainerSettings {
	return &ContainerSettings{
		container: container,
	}
}
