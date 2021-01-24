package lxd

import "gitlab.com/masakura/docker-machine-driver-lxd/lxd/client"

type ConfigureContainer struct {
	settings *client.ContainerSettings
}

func (c ConfigureContainer) Configure(sshPublicKey string) {
	c.enableSecurityNesting()
	c.addAuthorizedKeys(sshPublicKey)
}

func (c ConfigureContainer) enableSecurityNesting() {
	config := c.settings.Config()
	config.Set("security.nesting", "true")
}

func (c ConfigureContainer) addAuthorizedKeys(sshPublicKey string) {
	config := c.settings.Config()
	config.Set("user.user-data", "#cloud-config\nssh_authorized_keys:\n  - "+sshPublicKey)
}

func NewContainerConfigure(settings *client.ContainerSettings) ConfigureContainer {
	return ConfigureContainer{
		settings: settings,
	}
}
