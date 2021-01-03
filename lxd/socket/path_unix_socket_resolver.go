package socket

type PathUnixSocketResolver struct {
	fileResolver FileResolver
	path         string
}

func (r PathUnixSocketResolver) Resolve() string {
	if r.fileResolver.IsExist(r.path) {
		return r.path
	}
	return ""
}

func NewPathUnixSocketResolver(fileResolver FileResolver, path string) PathUnixSocketResolver {
	return PathUnixSocketResolver{
		fileResolver: fileResolver,
		path:         path,
	}
}
