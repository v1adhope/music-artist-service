package app

import (
	"context"
	"log"
	"time"

	"github.com/v1adhope/music-artist-service/internal/config"
)

// TODO: DI container
func Run(ctx context.Context) error {
	cfg, err := config.Build(ctx)
	if err != nil {
		return err
	}

	log.Println(cfg)

	time.Sleep(1 * time.Hour)

	return nil
}
