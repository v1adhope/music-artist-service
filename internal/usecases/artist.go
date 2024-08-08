package usecases

import (
	"context"

	"github.com/v1adhope/music-artist-service/internal/entities"
)

var _ ArtistUsecaser = (*ArtistUsecase)(nil)

type ArtistUsecase struct {
	ArtistRepo ArtistReposer
	Validator  Validater
}

func (u *ArtistUsecase) Get(ctx context.Context, id entities.ArtistId) (entities.Artist, error) {
	err := u.Validator.ValidateUuid(ctx, id.String())
	if err != nil {
		return entities.Artist{}, err
	}

	artist, err := u.ArtistRepo.Get(ctx, id)
	if err != nil {
		return entities.Artist{}, err
	}

	return artist, err
}

func (u *ArtistUsecase) GetAll(ctx context.Context) ([]entities.Artist, error) {
	artists, err := u.ArtistRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return artists, nil
}

func (u *ArtistUsecase) Create(ctx context.Context, artist entities.Artist) (entities.ArtistId, error) {
	id, err := u.ArtistRepo.Create(ctx, artist)
	if err != nil {
		return entities.ArtistId{}, err
	}

	return id, nil
}

func (u *ArtistUsecase) Replace(ctx context.Context, artist entities.Artist) error {
	id := artist.GetID()

	err := u.Validator.ValidateUuid(ctx, id)
	if err != nil {
		return err
	}

	if err := u.ArtistRepo.Replace(ctx, artist); err != nil {
		return err
	}

	return nil
}

func (u *ArtistUsecase) Delete(ctx context.Context, id entities.ArtistId) error {
	err := u.Validator.ValidateUuid(ctx, id.String())
	if err != nil {
		return err
	}

	if err := u.ArtistRepo.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}
