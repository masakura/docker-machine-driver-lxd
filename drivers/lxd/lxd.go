package lxd

import (
	"github.com/docker/machine/libmachine/drivers"
	"github.com/docker/machine/libmachine/log"
	"github.com/docker/machine/libmachine/mcnflag"
	"github.com/docker/machine/libmachine/state"
	"github.com/pkg/errors"
	"gitlab.com/masakura/docker-machine-driver-lxd/drivers/lxd/options"
)

type Driver struct {
	*drivers.BaseDriver
	Options options.Options
}

func (d *Driver) DriverName() string {
	return d.proxy().DriverName()
}

func (d *Driver) Create() error {
	return d.proxy().Create()
}

func (d *Driver) GetCreateFlags() []mcnflag.Flag {
	return []mcnflag.Flag{
		mcnflag.StringFlag{
			Name:   "lxd-external-network",
			EnvVar: "LXD_EXTERNAL_NETWORK",
			Usage:  "LXD host network exposed to the external (Not LXD guest)",
		},
		mcnflag.StringFlag{
			Name:   "lxd-remote",
			EnvVar: "LXD_REMOTE",
			Usage:  "LXD remote name",
		},
	}
}

func (d *Driver) GetIP() (string, error) {
	return d.GetSSHHostname()
}

func (d *Driver) GetSSHHostname() (string, error) {
	hostname, err := d.proxy().GetSSHHostname()
	if err != nil {
		return "", err
	}

	if hostname[0] != "1"[0] {
		return "", errors.New(hostname)
	}

	log.Debug("GetSSHHostname() => " + hostname)

	return hostname, nil
}

func (d *Driver) GetSSHUsername() string {
	return "ubuntu"
}

func (d *Driver) GetURL() (string, error) {
	url, err := d.proxy().GetURL()
	if err != nil {
		log.Error(err)
		return "", err
	}

	log.Debug("GetURL() => " + url)
	return url, nil
}

func (d *Driver) GetState() (state.State, error) {
	machineState, err := d.proxy().GetState()
	if err != nil {
		return state.None, err
	}
	log.Debug("GetState() => " + machineState.String())
	return machineState, err
}

func (d *Driver) Kill() error {
	return d.proxy().Kill()
}

func (d *Driver) Remove() error {
	return d.proxy().Remove()
}

func (d *Driver) Restart() error {
	return d.proxy().Restart()
}

func (d *Driver) SetConfigFromFlags(opts drivers.DriverOptions) error {
	d.Options = options.NewOptions(opts)
	return nil
}

func (d *Driver) Start() error {
	return d.proxy().Start()
}

func (d *Driver) Stop() error {
	return d.proxy().Stop()
}

func (d *Driver) proxy() *DriverProxy {
	return NewDriverProxy(d)
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
