package repositories

import "github.com/v1adhope/music-artist-service/pkg/postgresql"

type Repos struct {
	Artist ArtistReposer
}

func New(d *postgresql.Postgres) *Repos {
	return &Repos{
		Artist: NewArtist(d),
	}
}
