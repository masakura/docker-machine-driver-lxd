package socket

type UnixSocketResolver interface {
	Resolve() string
}

func DefaultUnixSocketResolver() UnixSocketResolver {
	return NewUnixSocketResolver(NewFileResolver())
}

func NewUnixSocketResolver(fileResolver FileResolver) UnixSocketResolver {
	return NewCompositeUnixSocketResolver([]UnixSocketResolver{
		NewPathUnixSocketResolver(fileResolver, "/var/snap/lxd/common/lxd/unix.socket"),
		NewPathUnixSocketResolver(fileResolver, "/var/lib/lxd/unix.socket"),
	})
}
