package client

type ContainerConfig struct {
	config map[string]string
}

func (c *ContainerConfig) Set(key string, value string) *ContainerConfig {
	c.config[key] = value
	return c
}

func NewConfigBuilder(config map[string]string) *ContainerConfig {
	return &ContainerConfig{config: config}
}
