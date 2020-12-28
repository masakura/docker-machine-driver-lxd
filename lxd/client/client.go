package client

import (
	lxd "github.com/lxc/lxd/client"
	"github.com/lxc/lxd/shared/api"
)

type LxdClient struct {
	Client *SyncClient
}

func (c *LxdClient) CreateContainer(name string, source api.ContainerSource) (*LxdContainer, error) {
	err := c.Client.CreateContainer(api.ContainersPost{
		Name:   name,
		Source: source,
	})
	if err != nil {
		return nil, err
	}

	return c.GetContainer(name), nil
}

func (c *LxdClient) GetContainer(name string) *LxdContainer {
	return NewLxdContainer(c.Client, name)
}

func NewLxdClientWith(instanceServer lxd.InstanceServer) *LxdClient {
	client := NewSyncClientWith(instanceServer)
	return &LxdClient{Client: client}
}
