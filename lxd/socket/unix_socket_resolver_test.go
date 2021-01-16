package socket

import (
	"github.com/stretchr/testify/assert"
	"gitlab.com/masakura/docker-machine-driver-lxd/lxd/file"
	"testing"
)

func TestResolveBySnap(t *testing.T) {
	files := file.NewMockFileResolver([]string{
		"/var/snap/lxd/common/lxd/unix.socket",
	})
	target := NewUnixSocketResolver(files)

	assert.Equal(t, "/var/snap/lxd/common/lxd/unix.socket", target.Resolve())
}

func TestResolveByLinuxPackage(t *testing.T) {
	files := file.NewMockFileResolver([]string{
		"/var/lib/lxd/unix.socket",
	})
	target := NewUnixSocketResolver(files)

	assert.Equal(t, "/var/lib/lxd/unix.socket", target.Resolve())
}
