package main

import (
	"github.com/bir/iken/config"
	"github.com/rs/zerolog/log"
)

type Config struct {
	LocalDebug  bool   `env:"DEBUG"`
	Port        int    `env:"PORT, 8080"`
	Host        string `env:"HOST, 0.0.0.0"`
	DatabaseURL string `env:"DATABASE_URL"`
}

func LoadConfig() *Config {
	c := Config{}
	if err := config.Load(&c); err != nil {
		log.Fatal().Err(err).Msg("error loading config")
	}

	if c.Port <= 0 || c.Port > 65535 {
		log.Fatal().Int("port", c.Port).Msg("invalid Port provided")
	}

	return &c
}
