package lxd

import (
	"github.com/docker/machine/libmachine/drivers"
	"github.com/docker/machine/libmachine/mcnflag"
	"github.com/docker/machine/libmachine/state"
	"github.com/pkg/errors"
	"log"
	"os"
)

var (
	outfile, _ = os.Create("log.txt")
	l          = log.New(outfile, "", 0)
)

type Driver struct {
	*drivers.BaseDriver
}

func (d *Driver) DriverName() string {
	return NewDriverProxy(d, nil, nil).DriverName()
}

func (d *Driver) Create() error {
	return NewDriverProxy(d, nil, nil).Create()
}

func (d *Driver) GetCreateFlags() []mcnflag.Flag {
	return []mcnflag.Flag{}
}

func (d *Driver) GetIP() (string, error) {
	return d.GetSSHHostname()
}

func (d *Driver) GetSSHHostname() (string, error) {
	hostname, err := NewDriverProxy(d, nil, nil).GetSSHHostname()
	if err != nil {
		return "", err
	}

	if hostname[0] != "1"[0] {
		return "", errors.New(hostname)
	}

	l.Println("GetSSHHostname() => " + hostname)

	return hostname, nil
}

func (d *Driver) GetSSHUsername() string {
	return "ubuntu"
}

func (d *Driver) GetURL() (string, error) {
	url, err := NewDriverProxy(d, nil, nil).GetURL()
	if err != nil {
		l.Panic(err)
		return "", err
	}

	l.Println("GetURL() => " + url)
	return url, nil
}

func (d *Driver) GetState() (state.State, error) {
	machineState, err := NewDriverProxy(d, nil, nil).GetState()
	if err != nil {
		return state.None, err
	}
	l.Println("GetState() => " + machineState.String())
	return machineState, err
}

func (d *Driver) Kill() error {
	l.Panic("Kill()")
	panic("implement me")
}

func (d *Driver) Remove() error {
	l.Panic("Remove()")
	panic("implement me")
}

func (d *Driver) Restart() error {
	l.Panic("Restart()")
	panic("implement me")
}

func (d *Driver) SetConfigFromFlags(opts drivers.DriverOptions) error {
	return nil
}

func (d *Driver) Start() error {
	l.Panic("Start()")
	panic("implement me")
}

func (d *Driver) Stop() error {
	l.Panic("Stop()")
	panic("implement me")
}

func NewDriver(hostName string, storePath string) *Driver {
	return &Driver{
		BaseDriver: newBaseDriver(hostName, storePath),
	}
}

func newBaseDriver(hostName string, storePath string) *drivers.BaseDriver {
	return &drivers.BaseDriver{
		MachineName: hostName,
		StorePath:   storePath,
		SSHUser:     "ubuntu",
		SSHPort:     22,
	}
}
