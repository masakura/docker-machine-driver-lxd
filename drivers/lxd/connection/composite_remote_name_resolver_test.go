package connection

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetFirstConnection(t *testing.T) {
	target := NewCompositeRemoteNameResolver([]RemoteNameResolver{newMockRemoteNameResolver("remote1"), newMockRemoteNameResolver("remote2")})

	actual, _ := target.Resolve()

	assert.Equal(t, "remote1", actual)
}

func TestSkipNil(t *testing.T) {
	target := NewCompositeRemoteNameResolver([]RemoteNameResolver{newMockRemoteNameResolver(""), newMockRemoteNameResolver("remote2")})

	actual, _ := target.Resolve()

	assert.Equal(t, "remote2", actual)
}

type mockProvider struct {
	remote string
}

func (p mockProvider) Resolve() (string, error) {
	return p.remote, nil
}

func newMockRemoteNameResolver(remote string) mockProvider {
	return mockProvider{
		remote: remote,
	}
}
