package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/v1adhope/music-artist-service/internal/config"
	v1 "github.com/v1adhope/music-artist-service/internal/controllers/grpc/v1"
	"github.com/v1adhope/music-artist-service/internal/usecases"
	"github.com/v1adhope/music-artist-service/internal/usecases/infrastructure/repositories"
	"github.com/v1adhope/music-artist-service/internal/usecases/infrastructure/validation"
	"github.com/v1adhope/music-artist-service/pkg/logger"
	"github.com/v1adhope/music-artist-service/pkg/postgresql"
)

func Run(ctx context.Context, cfg config.Config, logger logger.Logger) error {
	pg, err := postgresql.Build(ctx, cfg.Postgres)
	if err != nil {
		return err
	}

	repos := repositories.New(pg)

	validator := validation.New()

	uc := usecases.New(repos, validator)

	s, err := v1.Build(cfg.GrpcServer, uc, logger)
	if err != nil {
		return err
	}

	go s.Run()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	<-c

	s.Shutdown()

	return nil
}
