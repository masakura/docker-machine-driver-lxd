package lxd

import (
	"gitlab.com/masakura/docker-machine-driver-lxd/drivers/lxd/options"
	"gitlab.com/masakura/docker-machine-driver-lxd/lxd/client"
)

type ConfigureContainer struct {
	settings *client.ContainerSettings
}

func (c ConfigureContainer) Configure(sshPublicKey string, options options.Options) {
	c.enableSecurityNesting()
	c.addAuthorizedKeys(sshPublicKey)
	c.addExternalNetwork(options.ExternalNetwork)
}

func (c ConfigureContainer) enableSecurityNesting() {
	config := c.settings.Config()
	config.Set("security.nesting", "true")
}

func (c ConfigureContainer) addAuthorizedKeys(sshPublicKey string) {
	config := c.settings.Config()
	config.Set("user.user-data", "#cloud-config\nssh_authorized_keys:\n  - "+sshPublicKey)
}

func (c ConfigureContainer) addExternalNetwork(network string) {
	if network == "" {
		return
	}

	devices := c.settings.Devices()
	devices.Set("eth0", map[string]string{
		"name":    "eth0",
		"type":    "nic",
		"nictype": "macvlan",
		"parent":  network,
	})
}

func NewContainerConfigure(settings *client.ContainerSettings) ConfigureContainer {
	return ConfigureContainer{
		settings: settings,
	}
}
