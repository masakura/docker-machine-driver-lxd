package lxd

import (
	"github.com/docker/machine/libmachine/drivers"
	"github.com/docker/machine/libmachine/mcnflag"
	"github.com/docker/machine/libmachine/state"
)

type Driver struct {
	*drivers.BaseDriver
}

func (d Driver) Create() error {
	panic("implement me")
}

func (d Driver) GetCreateFlags() []mcnflag.Flag {
	panic("implement me")
}

func (d Driver) GetSSHHostname() (string, error) {
	panic("implement me")
}

func (d Driver) GetURL() (string, error) {
	panic("implement me")
}

func (d Driver) GetState() (state.State, error) {
	panic("implement me")
}

func (d Driver) Kill() error {
	panic("implement me")
}

func (d Driver) Remove() error {
	panic("implement me")
}

func (d Driver) Restart() error {
	panic("implement me")
}

func (d Driver) SetConfigFromFlags(opts drivers.DriverOptions) error {
	panic("implement me")
}

func (d Driver) Start() error {
	panic("implement me")
}

func (d Driver) Stop() error {
	panic("implement me")
}

func NewDriver(hostName string, storePath string) *Driver {
	return &Driver{}
}
