package mock_lxd

import "github.com/docker/machine/libmachine/drivers"

type FakeDriverOptions struct {
	flags map[string]interface{}
}

func (o *FakeDriverOptions) String(key string) string {
	if value, ok := o.flags[key]; ok {
		return value.(string)
	}
	return ""
}

func (o *FakeDriverOptions) StringSlice(string) []string {
	panic("implement me")
}

func (o *FakeDriverOptions) Int(string) int {
	panic("implement me")
}

func (o *FakeDriverOptions) Bool(string) bool {
	panic("implement me")
}

func NewFakeOptions(flags map[string]interface{}) drivers.DriverOptions {
	return &FakeDriverOptions{flags: flags}
}
