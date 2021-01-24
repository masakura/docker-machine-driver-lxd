package client

type ContainerDevices struct {
	devices map[string]map[string]string
}

func (d ContainerDevices) Set(key string, value map[string]string) {
	d.devices[key] = value
}

func NewContainerDevices(devices map[string]map[string]string) *ContainerDevices {
	return &ContainerDevices{
		devices: devices,
	}
}
