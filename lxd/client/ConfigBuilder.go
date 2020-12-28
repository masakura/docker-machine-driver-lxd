package client

type ConfigBuilder struct {
	config map[string]string
}

func (b *ConfigBuilder) Set(key string, value string) *ConfigBuilder {
	b.config[key] = value
	return b
}

func NewConfigBuilder(config map[string]string) *ConfigBuilder {
	return &ConfigBuilder{config: config}
}
