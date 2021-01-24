package lxd

import "github.com/docker/machine/libmachine/drivers"

type Options struct {
	ExternalNetwork string
}

func NewOptions(opts drivers.DriverOptions) Options {
	return Options{
		ExternalNetwork: opts.String("lxd-external-network"),
	}
}
