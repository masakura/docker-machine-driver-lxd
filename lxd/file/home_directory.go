package file

import (
	"os"
	"path"
)

type HomeDirectory interface {
	GetPath(relative string) string
}

type RealHomeDirectory struct {
}

func (h RealHomeDirectory) GetPath(relative string) string {
	home := os.Getenv("HOME")
	return path.Join(home, relative)
}

func NewHomeDirectory() HomeDirectory {
	return RealHomeDirectory{}
}
