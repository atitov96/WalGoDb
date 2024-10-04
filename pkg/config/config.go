package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"time"
)

type Config struct {
	Engine  EngineConfig  `yaml:"engine"`
	Network NetworkConfig `yaml:"network"`
	Logging LoggingConfig `yaml:"logging"`
}

type EngineConfig struct {
	Type string `yaml:"type"`
}

type NetworkConfig struct {
	Address        string        `yaml:"address"`
	MaxConnections int           `yaml:"max_connections"`
	MaxMessageSize string        `yaml:"max_message_size"`
	IdleTimeout    time.Duration `yaml:"idle_timeout"`
}

type LoggingConfig struct {
	Level  string `yaml:"level"`
	Output string `yaml:"output"`
}

func LoadConfig(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, fmt.Errorf("error parsing config file: %w", err)
	}

	setDefault(&cfg)

	return &cfg, nil
}

func setDefault(config *Config) {
	if config.Engine.Type == "" {
		config.Engine.Type = "in_memory"
	}
	if config.Network.Address == "" {
		config.Network.Address = "127.0.0.1:3223"
	}
	if config.Network.MaxConnections == 0 {
		config.Network.MaxConnections = 10
	}
	if config.Network.MaxMessageSize == "" {
		config.Network.MaxMessageSize = "4KB"
	}
	if config.Network.IdleTimeout == 0 {
		config.Network.IdleTimeout = 5 * time.Second
	}
	if config.Logging.Level == "" {
		config.Logging.Level = "info"
	}
	if config.Logging.Output == "" {
		config.Logging.Output = "stdout"
	}
}
