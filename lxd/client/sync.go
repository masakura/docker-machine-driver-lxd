package client

import (
	lxd "github.com/lxc/lxd/client"
	"github.com/lxc/lxd/shared/api"
)

type SyncClient struct {
	InstanceServer lxd.InstanceServer
}

func (c *SyncClient) GetConnection() (lxd.InstanceServer, error) {
	if c.InstanceServer == nil {
		connection, err := lxd.ConnectLXDUnix("/var/snap/lxd/common/lxd/unix.socket", nil)
		if err != nil {
			return nil, err
		}
		c.InstanceServer = connection
	}
	return c.InstanceServer, nil
}

func (c *SyncClient) CreateContainer(container api.ContainersPost) error {
	connection, err := c.GetConnection()
	if err != nil {
		return err
	}

	operation, err := connection.CreateContainer(container)
	if err != nil {
		return err
	}

	return operation.Wait()
}

func (c *SyncClient) UpdateContainerState(name string, state api.ContainerStatePut, ETag string) error {
	connection, err := c.GetConnection()
	if err != nil {
		return err
	}

	operation, err := connection.UpdateContainerState(name, state, ETag)
	if err != nil {
		return err
	}

	return operation.Wait()
}

func (c *SyncClient) GetState(name string) (state *api.ContainerState, ETag string, err error) {
	connection, err := c.GetConnection()
	if err != nil {
		return nil, "", err
	}

	return connection.GetContainerState(name)
}

func (c *SyncClient) GetContainer(name string) (*api.Container, string, error) {
	connection, err := c.GetConnection()
	if err != nil {
		return nil, "", err
	}

	return connection.GetContainer(name)
}

func (c *SyncClient) UpdateContainer(name string, container api.ContainerPut, ETag string) error {
	connection, err := c.GetConnection()
	if err != nil {
		return err
	}

	operation, err := connection.UpdateContainer(name, container, ETag)
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
