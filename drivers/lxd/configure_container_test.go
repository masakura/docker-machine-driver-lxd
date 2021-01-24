package lxd

import (
	"github.com/lxc/lxd/shared/api"
	"github.com/stretchr/testify/assert"
	"gitlab.com/masakura/docker-machine-driver-lxd/lxd/client"
	"testing"
)

func TestConfigure(t *testing.T) {
	settings := newEmptyContainerSettings()
	target := NewContainerConfigure(settings)

	target.Configure("[id_rsa.pub]", Options{ExternalNetwork: ""})

	assert.Equal(t, map[string]string{
		"security.nesting": "true",
		"user.user-data":   "#cloud-config\nssh_authorized_keys:\n  - [id_rsa.pub]",
	}, settings.Writable().Config)
}

func TestExternalNetwork(t *testing.T) {
	settings := newEmptyContainerSettings()
	target := NewContainerConfigure(settings)

	target.Configure("[id_rsa.pub]", Options{ExternalNetwork: "enp1s0"})

	assert.Equal(t, map[string]map[string]string{
		"eth0": {
			"name":    "eth0",
			"type":    "nic",
			"nictype": "macvlan",
			"parent":  "enp1s0",
		},
	}, settings.Writable().Devices)
}

func TestNoExternalNetwork(t *testing.T) {
	settings := newEmptyContainerSettings()
	target := NewContainerConfigure(settings)

	target.Configure("[id_rsa.pub]", Options{ExternalNetwork: ""})

	assert.Equal(t, map[string]map[string]string{}, settings.Writable().Devices)
}

func newEmptyContainerSettings() *client.ContainerSettings {
	return client.NewContainerSettings(&api.Container{ContainerPut: api.ContainerPut{
		Config:  map[string]string{},
		Devices: map[string]map[string]string{},
	}})
}
