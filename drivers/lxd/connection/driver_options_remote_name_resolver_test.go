package connection

import (
	"github.com/stretchr/testify/assert"
	"gitlab.com/masakura/docker-machine-driver-lxd/drivers/lxd/options"
	"gitlab.com/masakura/docker-machine-driver-lxd/mock_lxd"
	"testing"
)

func TestGetRemoteNameFromOptions(t *testing.T) {
	o := newOptions(map[string]interface{}{
		"lxd-remote": "remote1",
	})
	target := NewDriverOptionsProvider(o)

	actual, _ := target.Resolve()

	assert.Equal(t, "remote1", actual)
}

func TestGetRemoteNameNoOptions(t *testing.T) {
	target := NewDriverOptionsProvider(newOptions(map[string]interface{}{}))

	actual, _ := target.Resolve()

	assert.Equal(t, "", actual)
}

func newOptions(flags map[string]interface{}) options.Options {
	return options.NewOptions(mock_lxd.NewFakeOptions(flags))
}
