package config

import (
	"io/ioutil"
	"os"
	"testing"
	"time"
)

func TestLoadConfig(t *testing.T) {
	configContent := `
engine:
  type: "in_memory"
network:
  address: "127.0.0.1:3223"
  max_connections: 100
  max_message_size: "4KB"
  idle_timeout: 5m
logging:
  level: "info"
  output: "stdout"
`
	tmpfile, err := ioutil.TempFile("", "config.*.yaml")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write([]byte(configContent)); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatalf("Failed to close temp file: %v", err)
	}

	config, err := LoadConfig(tmpfile.Name())
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	if config.Engine.Type != "in_memory" {
		t.Errorf("Expected engine type 'in_memory', got '%s'", config.Engine.Type)
	}

	if config.Network.Address != "127.0.0.1:3223" {
		t.Errorf("Expected address '127.0.0.1:3223', got '%s'", config.Network.Address)
	}

	if config.Network.MaxConnections != 100 {
		t.Errorf("Expected max connections 100, got %d", config.Network.MaxConnections)
	}

	if config.Network.MaxMessageSize != "4KB" {
		t.Errorf("Expected max message size '4KB', got '%s'", config.Network.MaxMessageSize)
	}

	if config.Network.IdleTimeout != 5*time.Minute {
		t.Errorf("Expected idle timeout 5m, got %v", config.Network.IdleTimeout)
	}

	if config.Logging.Level != "info" {
		t.Errorf("Expected logging level 'info', got '%s'", config.Logging.Level)
	}

	if config.Logging.Output != "stdout" {
		t.Errorf("Expected logging output 'stdout', got '%s'", config.Logging.Output)
	}
}

func TestLoadConfigDefaults(t *testing.T) {
	configContent := `{}` // Empty config

	tmpfile, err := ioutil.TempFile("", "config.*.yaml")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write([]byte(configContent)); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatalf("Failed to close temp file: %v", err)
	}

	config, err := LoadConfig(tmpfile.Name())
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	if config.Engine.Type != "in_memory" {
		t.Errorf("Expected default engine type 'in_memory', got '%s'", config.Engine.Type)
	}

	if config.Network.Address != "127.0.0.1:3223" {
		t.Errorf("Expected default address '127.0.0.1:3223', got '%s'", config.Network.Address)
	}

	if config.Network.MaxConnections != 10 {
		t.Errorf("Expected default max connections 100, got %d", config.Network.MaxConnections)
	}

	if config.Network.MaxMessageSize != "4KB" {
		t.Errorf("Expected default max message size '4KB', got '%s'", config.Network.MaxMessageSize)
	}

	if config.Network.IdleTimeout != 5*time.Second {
		t.Errorf("Expected default idle timeout 5m, got %v", config.Network.IdleTimeout)
	}

	if config.Logging.Level != "info" {
		t.Errorf("Expected default logging level 'info', got '%s'", config.Logging.Level)
	}

	if config.Logging.Output != "stdout" {
		t.Errorf("Expected default logging output 'stdout', got '%s'", config.Logging.Output)
	}
}
