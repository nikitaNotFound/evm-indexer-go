package config

type NetworkConfig struct {
	StartBlock int64
	EndBlock   int64
	RpcUrl     string
}

type PGStorageConfig struct {
	ConnectionString string
}

type Config struct {
	NetworkConfig NetworkConfig
	isDebug       bool
	PGStorage     PGStorageConfig
}

func ParseConfig() (*Config, error) {
	return &Config{}, nil
}

func (c *Config) IsDebug() bool {
	return c.isDebug
}
