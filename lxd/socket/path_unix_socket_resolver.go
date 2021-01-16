package socket

import "gitlab.com/masakura/docker-machine-driver-lxd/lxd/file"

type PathUnixSocketResolver struct {
	fileResolver file.FileResolver
	path         string
}

func (r PathUnixSocketResolver) Resolve() string {
	if r.fileResolver.IsExist(r.path) {
		return r.path
	}
	return ""
}

func NewPathUnixSocketResolver(fileResolver file.FileResolver, path string) PathUnixSocketResolver {
	return PathUnixSocketResolver{
		fileResolver: fileResolver,
		path:         path,
	}
}
