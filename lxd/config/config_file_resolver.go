package config

import (
	"gitlab.com/masakura/docker-machine-driver-lxd/lxd/file"
	"path"
)

type ConfigFileResolver interface {
	Resolve() string
}

type RealConfigFileResolver struct {
	file file.FileResolver
	home file.HomeDirectory
}

func (r RealConfigFileResolver) Resolve() string {
	bases := []string{"snap/lxd/current", ""}

	for _, base := range bases {
		configPath := r.getConfigFilePath(base)
		if r.file.IsExist(configPath) {
			return configPath
		}
	}

	return ""
}

func (r RealConfigFileResolver) getConfigFilePath(base string) string {
	absolute := r.home.GetPath(base)
	return path.Join(absolute, ".config/lxc/config.yml")
}

func NewConfigFileResolver(f file.FileResolver, h file.HomeDirectory) ConfigFileResolver {
	if f == nil {
		f = file.NewFileResolver()
	}
	if h == nil {
		h = file.NewHomeDirectory()
	}

	return RealConfigFileResolver{
		file: f,
		home: h,
	}
}
