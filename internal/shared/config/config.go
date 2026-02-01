package config

import (
	"strings"

	"github.com/spf13/viper"
)

const (
	ENV_PRODUCTION  = "production"
	ENV_DEVELOPMENT = "development"
	ENV_LOCAL       = "local"
)

type Config struct {
	App struct {
		Name string
		Env  string
	}
	HTTP struct {
		Port int
	}
	DB struct {
		DSN      string
		Maxconns int
		Minconns int
	}
}

func LoadConfig() (*Config, error) {
	v := viper.New()

	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("./config")
	v.AddConfigPath(".")

	// Environment variable settings
	v.SetEnvPrefix("WALLET")
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Read configuration file
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	// Unmarshal into Config struct
	var Cfg Config
	if err := v.Unmarshal(&Cfg); err != nil {
		return nil, err
	}
	return &Cfg, nil
}
