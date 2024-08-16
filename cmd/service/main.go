package main

import (
	"context"
	"log"

	"github.com/v1adhope/music-artist-service/internal/app"
	"github.com/v1adhope/music-artist-service/internal/config"
	"github.com/v1adhope/music-artist-service/pkg/logger"
)

func main() {
	ctx := context.Background()

	cfg := config.Must(ctx)

	logger := logger.New(cfg.Logger)

	if err := app.Run(ctx, cfg, logger); err != nil {
		log.Fatal(err)
	}
}
