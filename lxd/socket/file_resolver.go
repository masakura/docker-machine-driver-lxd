package socket

import (
	"github.com/docker/machine/libmachine/log"
	"os"
	"strconv"
)

type FileResolver interface {
	IsExist(path string) bool
}

type RealFileResolver struct {
}

func (r RealFileResolver) IsExist(path string) bool {
	_, err := os.Stat(path)
	exist := os.IsNotExist(err)
	log.Debug(path + " => " + strconv.FormatBool(exist))
	return !exist
}

func NewFileResolver() FileResolver {
	return RealFileResolver{}
}
