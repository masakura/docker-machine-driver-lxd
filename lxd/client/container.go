package client

import (
	"github.com/lxc/lxd/shared/api"
)

type LxdContainer struct {
	Client *SyncClient
	Name   string
}

func (c *LxdContainer) Start() error {
	client := c.Client

	return client.UpdateContainerState(c.Name, api.ContainerStatePut{
		Action:  "start",
		Timeout: -1,
	}, "")
}

func (c *LxdContainer) Stop() error {
	client := c.Client

	return client.UpdateContainerState(c.Name, api.ContainerStatePut{
		Action:  "stop",
		Timeout: -1,
	}, "")
}

func (c *LxdContainer) Restart() error {
	client := c.Client

	return client.UpdateContainerState(c.Name, api.ContainerStatePut{
		Action:  "restart",
		Timeout: -1,
	}, "")
}

func (c *LxdContainer) Delete() error {
	client := c.Client

	return client.DeleteContainer(c.Name)
}

func (c *LxdContainer) GetState() (*api.ContainerState, string, error) {
	client := c.Client

	return client.GetState(c.Name)
}

func (c *LxdContainer) GetSettings() (*ContainerSettings, string, error) {
	container, s, err := c.Client.GetContainer(c.Name)
	if err != nil {
		return nil, "", nil
	}

	return NewContainerSettings(container), s, nil
}

func (c *LxdContainer) Update(container api.ContainerPut, ETag string) error {
	return c.Client.UpdateContainer(c.Name, container, ETag)
}

func NewLxdContainer(client *SyncClient, name string) *LxdContainer {
	return &LxdContainer{
		Client: client,
		Name:   name,
	}
}
