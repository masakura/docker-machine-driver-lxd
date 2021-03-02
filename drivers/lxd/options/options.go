package options

import "github.com/docker/machine/libmachine/drivers"

type Options struct {
	ExternalNetwork string
	Remote          string
}

func NewOptions(opts drivers.DriverOptions) Options {
	return Options{
		ExternalNetwork: opts.String("lxd-external-network"),
		Remote:          opts.String("lxd-remote"),
	}
}
