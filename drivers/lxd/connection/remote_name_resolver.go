package connection

type RemoteNameResolver interface {
	Resolve() (string, error)
}
