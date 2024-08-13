package main

import (
	"context"
	"log"

	"github.com/v1adhope/music-artist-service/internal/app"
)

func main() {
	ctx := context.Background()

	if err := app.Run(ctx); err != nil {
		log.Fatal(err)
	}
}
