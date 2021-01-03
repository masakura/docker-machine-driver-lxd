package socket

type MockFileResolver struct {
	paths []string
}

func (r MockFileResolver) IsExist(path string) bool {
	for _, s := range r.paths {
		if s == path {
			return true
		}
	}
	return false
}

func NewMockFileResolver(paths []string) FileResolver {
	return MockFileResolver{
		paths: paths,
	}
}
