package config

import (
	"cmp"
	"context"
	"fmt"
	"strconv"

	"github.com/joho/godotenv"

	v1 "github.com/v1adhope/music-artist-service/internal/controllers/grpc/v1"
	"github.com/v1adhope/music-artist-service/pkg/logger"
	"github.com/v1adhope/music-artist-service/pkg/postgresql"
)

type Config struct {
	Postgres   postgresql.Config
	Logger     logger.Config
	GrpcServer v1.Config
}

func Must(ctx context.Context) Config {
	if err := godotenv.Load(); err != nil {
		panic(fmt.Errorf("config: can't load envs from .env: %v", err))
	}

	envs, err := godotenv.Read()
	if err != nil {
		panic(fmt.Sprintf("config: can't read envs from .env: %v", err))
	}

	return Config{
		Postgres: postgresql.Config{
			ConnStr: cmp.Or(
				envs["APP_POSTGRES_CONN_STR"],
				"postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable",
			),
		},
		Logger: logger.Config{
			Level: cmp.Or(
				envs["APP_LOGGER_LEVEL"],
				"info",
			),
		},
		GrpcServer: v1.Config{
			Socket: cmp.Or(
				envs["APP_GRPC_SOCKET"],
				":50051",
			),
			WithTls: cmp.Or(
				mustBool(envs["APP_GRPC_WITH_TLS"]),
				false,
			),
			TlsFilePaths: v1.TlsFilePaths{
				Key: cmp.Or(
					envs["APP_GRPC_CA_KEY"],
					"./cert/ca_key.pem",
				),
				Cert: cmp.Or(
					envs["APP_GRPC_CA_CERT"],
					"./cert/ca_cert.pem",
				),
			},
		},
	}
}

func mustBool(target string) bool {
	value, err := strconv.ParseBool(target)
	if err != nil {
		panic(fmt.Sprintf("can't parse bool from %s: %v", target, err))
	}

	return value
}
