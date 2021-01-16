package config

import (
	"github.com/stretchr/testify/assert"
	"gitlab.com/masakura/docker-machine-driver-lxd/lxd/file"
	"testing"
)

func TestResolveBySnap(t *testing.T) {
	f := file.NewMockFileResolver([]string{"/home/ubuntu/snap/lxd/current/.config/lxc/config.yml"})
	h := file.NewMockHomeDirectory("/home/ubuntu")

	target := NewConfigFileResolver(f, h)

	assert.Equal(t, "/home/ubuntu/snap/lxd/current/.config/lxc/config.yml", target.Resolve())
}

func TestResolveByLinuxPackage(t *testing.T) {
	f := file.NewMockFileResolver([]string{"/home/ubuntu/.config/lxc/config.yml"})
	h := file.NewMockHomeDirectory("/home/ubuntu")

	target := NewConfigFileResolver(f, h)

	assert.Equal(t, "/home/ubuntu/.config/lxc/config.yml", target.Resolve())
}
