package repository

import (
	"context"

	"github.com/v1adhope/music-artist-service/internal/entities"
)

type ArtistReposer interface {
	Get(ctx context.Context, id entities.ArtistId) (entities.Artist, error)
	GetAll(ctx context.Context) ([]entities.Artist, error)
	Create(ctx context.Context, artist entities.Artist) (entities.ArtistId, error)
	Replace(ctx context.Context, artist entities.Artist) error
	Delete(ctx context.Context, id entities.ArtistId) error
}
