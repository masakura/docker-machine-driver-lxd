package lxd

import (
	"github.com/docker/machine/libmachine/state"
	"github.com/golang/mock/gomock"
	lxd "github.com/lxc/lxd/client"
	"github.com/lxc/lxd/shared/api"
	"github.com/stretchr/testify/assert"
	"gitlab.com/masakura/docker-machine-driver-lxd/mock_lxd"
	"gitlab.com/masakura/docker-machine-driver-lxd/mock_ssh"
	"gitlab.com/masakura/docker-machine-driver-lxd/ssh"
	"testing"
)

func TestNewDriverProxy(t *testing.T) {
	driver := CreateTestingDriverProxy("host1", nil, nil)

	assert.NotNil(t, driver)
}

func TestCreate(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockServer := mock_lxd.NewMockInstanceServer(controller)
	mockOperation := mock_lxd.NewMockOperation(controller)
	mockSSH := mock_ssh.NewMockSSHKeyProvider(controller)

	mockSSH.EXPECT().Generate("machines/host1/id_rsa").Return(nil)
	mockSSH.EXPECT().GetPublicKey("machines/host1/id_rsa").Return("ssh-rsa <key>", nil)

	mockServer.EXPECT().CreateContainer(api.ContainersPost{
		Name: "docker-machine-host1",
		Source: api.ContainerSource{
			Type:     "image",
			Mode:     "pull",
			Server:   "https://cloud-images.ubuntu.com/releases",
			Protocol: "simplestreams",
			Alias:    "20.04",
		},
	}).Return(mockOperation, nil)
	mockOperation.EXPECT().Wait()

	mockServer.EXPECT().GetContainer("docker-machine-host1").Return(&api.Container{
		ContainerPut: api.ContainerPut{
			Config: map[string]string{},
		},
	}, "tag1", nil)

	mockOperation = mock_lxd.NewMockOperation(controller)
	mockServer.EXPECT().UpdateContainer("docker-machine-host1", api.ContainerPut{
		Config: map[string]string{
			"security.nesting": "true",
			"user.user-data":   "#cloud-config\nssh_authorized_keys:\n  - ssh-rsa <key>",
		},
	}, "tag1").Return(mockOperation, nil)
	mockOperation.EXPECT().Wait()

	mockOperation = mock_lxd.NewMockOperation(controller)
	mockServer.EXPECT().UpdateContainerState("docker-machine-host1", api.ContainerStatePut{
		Action:  "start",
		Timeout: -1,
	}, "").Return(mockOperation, nil)
	mockOperation.EXPECT().Wait()

	driver := CreateTestingDriverProxy("host1", mockServer, mockSSH)

	err := driver.Create()

	assert.Nil(t, err)
}

func TestGetState(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockServer := mock_lxd.NewMockInstanceServer(controller)
	mockServer.EXPECT().GetContainerState("docker-machine-host1").Return(&api.ContainerState{
		StatusCode: api.Running,
	}, "", nil)

	driver := CreateTestingDriverProxy("host1", mockServer, nil)

	s, err := driver.GetState()

	assert.Equal(t, state.Running, s)
	assert.Nil(t, err)
}

func TestGetSSHHostname(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockServer := mock_lxd.NewMockInstanceServer(controller)
	mockServer.EXPECT().GetContainerState("docker-machine-host1").Return(&api.ContainerState{
		Network: map[string]api.ContainerStateNetwork{
			"eth0": {
				Addresses: []api.ContainerStateNetworkAddress{
					{Address: "192.168.234.12"},
				},
			},
		},
	}, "", nil)

	driver := CreateTestingDriverProxy("host1", mockServer, nil)

	host, err := driver.GetSSHHostname()

	assert.Equal(t, "192.168.234.12", host)
	assert.Nil(t, err)
}

func TestGetURL(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockServer := mock_lxd.NewMockInstanceServer(controller)
	mockServer.EXPECT().GetContainerState("docker-machine-host1").Return(&api.ContainerState{
		Network: map[string]api.ContainerStateNetwork{
			"eth0": {
				Addresses: []api.ContainerStateNetworkAddress{
					{Address: "192.168.234.12"},
				},
			},
		},
	}, "", nil)

	driver := CreateTestingDriverProxy("host1", mockServer, nil)

	url, err := driver.GetURL()

	assert.Equal(t, "tcp://192.168.234.12:2376", url)
	assert.Nil(t, err)
}

func TestGet(t *testing.T) {
	proxy := CreateTestingDriverProxy("host1", nil, nil)

	assert.Equal(t, "ubuntu", proxy.GetSSHUsername())
}

func TestDriverName(t *testing.T) {
	proxy := CreateTestingDriverProxy("host1", nil, nil)

	assert.Equal(t, "lxd", proxy.DriverName())
}

func CreateTestingDriverProxy(name string, connection lxd.InstanceServer, ssh ssh.SSHKeyProvider) *DriverProxy {
	return NewDriverProxy(&Driver{
		BaseDriver: newBaseDriver(name, ""),
	}, connection, ssh)
}
