package socket

import "gitlab.com/masakura/docker-machine-driver-lxd/lxd/file"

type UnixSocketResolver interface {
	Resolve() string
}

func DefaultUnixSocketResolver() UnixSocketResolver {
	return NewUnixSocketResolver(file.NewFileResolver())
}

func NewUnixSocketResolver(fileResolver file.FileResolver) UnixSocketResolver {
	return NewCompositeUnixSocketResolver([]UnixSocketResolver{
		NewPathUnixSocketResolver(fileResolver, "/var/snap/lxd/common/lxd/unix.socket"),
		NewPathUnixSocketResolver(fileResolver, "/var/lib/lxd/unix.socket"),
	})
}
