package config

import (
	env "github.com/caarlos0/env/v6"
)

type Config struct {
	Postgres Postgres
	Redis    Redis
}

type Postgres struct {
	Host     string `env:"PG_HOST"`
	Port     int    `env:"PG_PORT"`
	User     string `env:"PG_USER"`
	Password string `env:"PG_PASSWORD"`
	Dbname   string `env:"PG_DB_NAME"`
}

type Redis struct {
	Addr     string `env:"REDIS_ADDR"`
	Password string `env:"REDIS_PASSWORD"`
	DB       int    `env:"REDIS_DB_NAME"`
}

func NewFromEnv() (*Config, error) {
	var config Config
	if err := env.Parse(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
