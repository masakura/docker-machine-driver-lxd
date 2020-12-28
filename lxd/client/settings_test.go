package client

import (
	"github.com/lxc/lxd/shared/api"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSecurityNesting(t *testing.T) {
	builder := NewSettingsBuilder(&api.Container{ContainerPut: api.ContainerPut{Config: map[string]string{}}})

	builder.Config().Set("security.nesting", "true")

	assert.Equal(t, api.ContainerPut{
		Config: map[string]string{
			"security.nesting": "true",
		},
	}, builder.Writable())
}
