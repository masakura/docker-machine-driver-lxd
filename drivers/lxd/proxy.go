package lxd

import (
	"github.com/docker/machine/libmachine/log"
	"github.com/docker/machine/libmachine/state"
	lxd "github.com/lxc/lxd/client"
	"github.com/lxc/lxd/shared/api"
	"gitlab.com/masakura/docker-machine-driver-lxd/drivers/lxd/utils"
	"gitlab.com/masakura/docker-machine-driver-lxd/lxd/client"
	"gitlab.com/masakura/docker-machine-driver-lxd/lxd/socket"
	"gitlab.com/masakura/docker-machine-driver-lxd/ssh"
)

type DriverProxy struct {
	driver     *Driver
	lxdClient  *client.LxdClient
	ssh        ssh.SSHKeyProvider
	unixSocket socket.UnixSocketResolver
}

func (p *DriverProxy) Create() error {
	c := p.lxdClient
	d := p.driver
	s := p.GetSSHKeyProvider()

	if err := s.Generate(d.GetSSHKeyPath()); err != nil {
		return err
	}

	log.Info("Creating LXD container...")
	container, err := c.CreateContainer(p.getContainerName(), api.ContainerSource{
		Type:     "image",
		Mode:     "pull",
		Server:   "https://cloud-images.ubuntu.com/releases",
		Protocol: "simplestreams",
		Alias:    "20.04",
	})
	if err != nil {
		return err
	}

	log.Info("Configure LXD container...")
	settings, tag, err := container.Get()
	if err != nil {
		return err
	}

	settingsBuilder := client.NewSettingsBuilder(settings)
	settingsBuilder.Config().Set("security.nesting", "true")

	key, err := s.GetPublicKey(d.GetSSHKeyPath())
	if err != nil {
		return err
	}

	settingsBuilder.Config().Set("user.user-data", "#cloud-config\nssh_authorized_keys:\n  - "+key)

	container.Update(settingsBuilder.Writable(), tag)

	log.Info("Starting LXD container...")
	err = container.Start()
	if err != nil {
		return err
	}

	return nil
}

func (p *DriverProxy) GetState() (state.State, error) {
	containerState, _, err := p.getContainer().GetState()
	if err != nil {
		return state.None, err
	}

	return utils.GetDockerMachineState(containerState), nil
}

func (p *DriverProxy) GetSSHHostname() (string, error) {
	container := p.lxdClient.GetContainer("docker-machine-" + p.driver.MachineName)

	containerState, _, err := container.GetState()
	if err != nil {
		return "", err
	}

	addresses := containerState.Network["eth0"].Addresses
	if len(addresses) > 0 {
		return addresses[0].Address, nil
	}
	return "", nil
}

func (p *DriverProxy) GetSSHKeyProvider() ssh.SSHKeyProvider {
	if p.ssh != nil {
		return p.ssh
	}

	return ssh.NewSSHKeyProvider()
}

func (p *DriverProxy) GetSSHUsername() string {
	return p.driver.GetSSHUsername()
}

func (p *DriverProxy) GetURL() (string, error) {
	hostname, err := p.GetSSHHostname()
	if err != nil {
		return "", err
	}

	if hostname == "" {
		return "", nil
	}

	return "tcp://" + hostname + ":2376", nil
}

func (p *DriverProxy) DriverName() string {
	return "lxd"
}

func (p *DriverProxy) Start() error {
	return p.getContainer().Start()
}

func (p *DriverProxy) Stop() error {
	return p.getContainer().Stop()
}

func (p *DriverProxy) Restart() error {
	return p.getContainer().Restart()
}

func (p *DriverProxy) Remove() error {
	container := p.getContainer()

	// ignore stop error. (`Error: The container is already stopped`)
	_ = container.Stop()

	return container.Delete()
}

func (p *DriverProxy) Kill() error {
	return p.getContainer().Stop()
}

func (p *DriverProxy) getContainerName() string {
	return "docker-machine-" + p.driver.GetMachineName()
}

func (p *DriverProxy) getContainer() *client.LxdContainer {
	return p.lxdClient.GetContainer(p.getContainerName())
}

func NewDriverProxy(driver *Driver, connection lxd.InstanceServer, ssh ssh.SSHKeyProvider) *DriverProxy {
	return &DriverProxy{
		driver:    driver,
		lxdClient: client.NewLxdClientWith(connection),
		ssh:       ssh,
	}
}
