package client

import (
	lxd "github.com/lxc/lxd/client"
	"github.com/lxc/lxd/shared/api"
)

type SyncClient struct {
	InstanceServer lxd.InstanceServer
}

func (c *SyncClient) CreateContainer(container api.ContainersPost) error {
	connection := c.InstanceServer

	operation, err := connection.CreateContainer(container)
	if err != nil {
		return err
	}

	return operation.Wait()
}

func (c *SyncClient) UpdateContainerState(name string, state api.ContainerStatePut, ETag string) error {
	connection := c.InstanceServer

	operation, err := connection.UpdateContainerState(name, state, ETag)
	if err != nil {
		return err
	}

	return operation.Wait()
}

func (c *SyncClient) GetState(name string) (state *api.ContainerState, ETag string, err error) {
	connection := c.InstanceServer

	return connection.GetContainerState(name)
}

func (c *SyncClient) GetContainer(name string) (*api.Container, string, error) {
	connection := c.InstanceServer

	return connection.GetContainer(name)
}

func (c *SyncClient) UpdateContainer(name string, container api.ContainerPut, ETag string) error {
	connection := c.InstanceServer

	operation, err := connection.UpdateContainer(name, container, ETag)
	if err != nil {
		return err
	}

	return operation.Wait()
}

func (c *SyncClient) DeleteContainer(name string) error {
	connection := c.InstanceServer

	operation, err := connection.DeleteContainer(name)
	if err != nil {
		return err
	}

	return operation.Wait()
}

func NewSyncClientWith(connection lxd.InstanceServer) *SyncClient {
	return &SyncClient{
		InstanceServer: connection,
	}
}
