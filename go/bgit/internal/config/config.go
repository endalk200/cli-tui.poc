package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

// Provider represents an AI provider configuration
type Provider struct {
	Name    string `mapstructure:"name"`
	EnvName string `mapstructure:"env_name"`
}

// Config holds all configuration for bgit
type Config struct {
	AIProvider Provider `mapstructure:"ai_provider"`
}

var (
	// Global config instance
	cfg *Config
)

// InitConfig initializes the configuration using Viper
func InitConfig(cfgFile string) error {
	if cfgFile != "" {
		// Use config file from the flag
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory
		home, err := os.UserHomeDir()
		if err != nil {
			return fmt.Errorf("failed to get home directory: %w", err)
		}

		// Search for config in home directory with name ".bgit" (without extension)
		viper.AddConfigPath(home)
		viper.AddConfigPath(".") // Also check current directory
		viper.SetConfigType("yaml")
		viper.SetConfigName(".bgit")
	}

	// Set default values
	viper.SetDefault("ai_provider.name", "OpenAI")
	viper.SetDefault("ai_provider.env_name", "OPENAI_API_KEY")

	// Enable environment variable support
	viper.AutomaticEnv()

	// Read in config file (if it exists)
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; create a default one
			if err := createDefaultConfig(); err != nil {
				return fmt.Errorf("failed to create default config: %w", err)
			}
		} else {
			return fmt.Errorf("failed to read config file: %w", err)
		}
	}

	// Unmarshal config into struct
	cfg = &Config{}
	if err := viper.Unmarshal(cfg); err != nil {
		return fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return nil
}

// createDefaultConfig creates a default configuration file
func createDefaultConfig() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	configPath := filepath.Join(home, ".bgit.yaml")

	// Set defaults
	viper.Set("ai_provider.name", "OpenAI")
	viper.Set("ai_provider.env_name", "OPENAI_API_KEY")

	// Write config file
	if err := viper.WriteConfigAs(configPath); err != nil {
		return err
	}

	fmt.Printf("Created default config file at: %s\n", configPath)

	// Read the newly created config
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}

// GetConfig returns the current configuration
func GetConfig() *Config {
	if cfg == nil {
		cfg = &Config{
			AIProvider: Provider{
				Name:    "OpenAI",
				EnvName: "OPENAI_API_KEY",
			},
		}
	}
	return cfg
}

// GetProvider returns the configured AI provider
func GetProvider() Provider {
	return GetConfig().AIProvider
}

// SetProvider updates the AI provider in the config
func SetProvider(name, envName string) error {
	viper.Set("ai_provider.name", name)
	viper.Set("ai_provider.env_name", envName)

	// Update in-memory config
	cfg.AIProvider = Provider{
		Name:    name,
		EnvName: envName,
	}

	return viper.WriteConfig()
}

// Available providers for reference
var AvailableProviders = []Provider{
	{
		Name:    "OpenAI",
		EnvName: "OPENAI_API_KEY",
	},
	{
		Name:    "OpenRouter",
		EnvName: "OPENROUTER_API_KEY",
	},
	{
		Name:    "Anthropic",
		EnvName: "ANTHROPIC_API_KEY",
	},
}
