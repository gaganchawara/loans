package config

import (
	"path/filepath"
	"runtime"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	// Define a mock configuration struct
	var cfg struct {
		SomeValue string `toml:"some_value"`
	}

	// Define a base path, environment, and an expected value
	filePath := getTestFilePath()
	env := "test"
	expectedValue := "test_value"

	// Load the configuration
	err := LoadConfig(filePath, env, &cfg)
	if err != nil {
		t.Fatalf("Error loading configuration: %v", err)
	}

	// Check if the loaded configuration matches the expected value
	if cfg.SomeValue != expectedValue {
		t.Errorf("Expected: %s, Got: %s", expectedValue, cfg.SomeValue)
	}
}

func getTestFilePath() string {
	_, b, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(b), "testdata")
}
