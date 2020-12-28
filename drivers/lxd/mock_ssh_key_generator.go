package lxd

import (
	"github.com/docker/machine/drivers/virtualbox"
	"github.com/stretchr/testify/assert"
	"testing"
)

type MockSSHKeyGenerator struct {
	t         *testing.T
	validPath string
}

func (m *MockSSHKeyGenerator) Generate(path string) error {
	assert.Equal(m.t, m.validPath, path)
	return nil
}

func NewMockSSHKeyGenerator(t *testing.T, validPath string) virtualbox.SSHKeyGenerator {
	return &MockSSHKeyGenerator{
		t:         t,
		validPath: validPath,
	}
}
