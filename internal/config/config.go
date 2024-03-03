package config

import (
	env "github.com/caarlos0/env/v6"
)

type Config struct {
	Postgres Postgres
	Redis    Redis
}

type Log struct {
	Level string `env:"LOG_LEVEL"`
	Mod   string `env:"LOG_MODE"`
}

type Postgres struct {
	Host string `env:"POSTGRES_STORE"`
	// ...
}

type Redis struct {
	Host string `env:"REDIS_HOST"`
	// ...
}

// NewFromEnv creates reads application configuration from the file.
func NewFromEnv() (*Config, error) {
	var config Config
	if err := env.Parse(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
