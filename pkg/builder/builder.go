package builder

type Config struct {
	ValueA string
	ValueB string
}

type ConfigBuilder struct {
	config *Config
}

func (c *ConfigBuilder) ValueA(value string) *ConfigBuilder {
	c.config.ValueA = value
	return c
}

func (c *ConfigBuilder) ValueB(value string) *ConfigBuilder {
	c.config.ValueB = value
	return c
}

func (c *ConfigBuilder) Build() *Config {
	return c.config
}

func NewConfigBuilder() *ConfigBuilder {
	return &ConfigBuilder{config: &Config{}}
}
