package options

import (
	"github.com/stretchr/testify/assert"
	"gitlab.com/masakura/docker-machine-driver-lxd/mock_lxd"
	"testing"
)

func TestNewOptions(t *testing.T) {
	actual := NewOptions(mock_lxd.NewFakeOptions(map[string]interface{}{
		"lxd-external-network": "enp1s0",
		"lxd-remote":           "remote1",
	}))

	assert.Equal(t, Options{
		ExternalNetwork: "enp1s0",
		Remote:          "remote1",
	}, actual)
}
