package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type WhatsAppConfig struct {
	Session string `yaml:"session"`
}

type TelegramConfig struct {
	Token string `yaml:"token"`
}

type DatabaseConfig struct {
	SID      string `yaml:"sid"`
	Username string `yaml:"user"`
	Password string `yaml:"password"`
}

type Config struct {
	WhatsApp WhatsAppConfig `yaml:"whatsapp"`
	Telegram TelegramConfig `yaml:"telegram"`
	Database DatabaseConfig `yaml:"database"`
}

// LoadConfig loads the configuration from a YAML file
func LoadConfig() (Config, error) {
	var cfg Config

	// Open the config file
	file, err := os.Open("config.yaml")
	if err != nil {
		return cfg, err
	}
	defer file.Close()

	// Decode the YAML data into the Config struct
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return cfg, err
	}

	return cfg, nil
}
