package file

import "path"

type MockHomeDirectory struct {
	home string
}

func (h MockHomeDirectory) GetPath(relative string) string {
	return path.Join(h.home, relative)
}

func NewMockHomeDirectory(home string) MockHomeDirectory {
	return MockHomeDirectory{home: home}
}
