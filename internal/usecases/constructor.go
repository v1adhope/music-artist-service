package usecases

import (
	"github.com/v1adhope/music-artist-service/internal/usecases/infrastructure/repositories"
	"github.com/v1adhope/music-artist-service/internal/usecases/infrastructure/validation"
)

type Usecases struct {
	Artist *ArtistUsecase
}

func New(r *repositories.Repos, v *validation.Validator) *Usecases {
	return &Usecases{
		Artist: NewArtist(r.Artist, v),
	}
}
