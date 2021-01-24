package lxd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewOptions(t *testing.T) {
	actual := NewOptions(NewMockDriverOptions(map[string]string{
		"lxd-external-network": "enp1s0",
	}))

	assert.Equal(t, Options{
		ExternalNetwork: "enp1s0",
	}, actual)
}
