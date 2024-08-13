package repository

import "github.com/v1adhope/music-artist-service/pkg/postgresql"

type Repos struct {
	Artist *ArtistRepo
}

func New(d *postgresql.Postgres) *Repos {
	return &Repos{
		Artist: NewArtist(d),
	}
}
