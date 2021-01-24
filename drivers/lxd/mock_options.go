package lxd

import "github.com/docker/machine/libmachine/drivers"

type MockDriverOptions struct {
	options map[string]string
}

func (o MockDriverOptions) String(key string) string {
	return o.options[key]
}

func (o MockDriverOptions) StringSlice(string) []string {
	panic("not supported")
}

func (o MockDriverOptions) Int(string) int {
	panic("not supported")
}

func (o MockDriverOptions) Bool(string) bool {
	panic("not supported")
}

func NewMockDriverOptions(options map[string]string) drivers.DriverOptions {
	return MockDriverOptions{
		options: options,
	}
}
