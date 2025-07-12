package config

type NetworkConfig struct {
	StartBlock uint64
	EndBlock   uint64
}

type Config struct {
	NetworkConfig NetworkConfig
	isDebug       bool
}

func ParseConfig() (*Config, error) {
	return &Config{}, nil
}

func (c *Config) IsDebug() bool {
	return c.isDebug
}
