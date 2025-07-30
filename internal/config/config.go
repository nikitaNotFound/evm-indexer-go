package config

import (
	"encoding/json"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/nikitaNotFound/evm-indexer-go/internal/networks"
)

type NetworkConfig struct {
	StartBlock int64            `json:"start_block" validate:"min=0"`
	EndBlock   int64            `json:"end_block" validate:"min=0"`
	RpcUrl     string           `json:"rpc_url" validate:"required,url"`
	Network    networks.Network `json:"network" validate:"required,oneof=eth arbitrum optimism base bnb"`
}

type PGStorageConfig struct {
	ConnectionString    string `json:"connection_string" validate:"required,url"`
	CreateDBIfNotExists bool   `json:"create_db_if_not_exists"`

	// Connection Pool Settings
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
	ConnMaxIdleTime time.Duration
}

type Config struct {
	NetworkConfig NetworkConfig   `json:"network_config" validate:"required"`
	PGStorage     PGStorageConfig `json:"pg_storage_config" validate:"required"`
	isDebug       bool
}

// ParseConfig reads and validates configuration from config.json file
func ParseConfig() (*Config, error) {
	data, err := os.ReadFile("config.json")
	if err != nil {
		return nil, err
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	// Set default values
	config.PGStorage.CreateDBIfNotExists = true

	validator := validator.New()
	if err := validator.Struct(config); err != nil {
		return nil, err
	}

	return &config, nil
}

func (c *Config) IsDebug() bool {
	return c.isDebug
}
