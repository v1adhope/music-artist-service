package config

import (
	"context"
	"fmt"

	"github.com/joho/godotenv"

	"github.com/v1adhope/music-artist-service/pkg/postgresql"
)

type Config struct {
	Postgres postgresql.Config
}

func Build(ctx context.Context) (Config, error) {
	if err := godotenv.Load(); err != nil {
		return Config{}, fmt.Errorf("config: can't load envs from .env: %w", err)
	}

	envs, err := godotenv.Read()
	if err != nil {
		return Config{}, fmt.Errorf("config: can't read envs from .env: %w", err)
	}

	return Config{
		Postgres: postgresql.Config{
			ConnStr: envs["APP_POSTGRES_CONN_STR"],
		},
	}, nil
}
