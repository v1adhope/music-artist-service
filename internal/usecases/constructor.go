package usecases

import (
	"github.com/v1adhope/music-artist-service/internal/usecases/infrastructure/repository"
	"github.com/v1adhope/music-artist-service/internal/usecases/infrastructure/validation"
)

type Usecases struct {
	Artist *ArtistUsecase
}

func New(r repository.Repos, v *validation.Validator) *Usecases {
	return &Usecases{
		Artist: NewArtist(r.Artist, v),
	}
}
