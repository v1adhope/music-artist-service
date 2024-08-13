// INFO: how to avoid empty values (default)
// INFO: how to handle errors after seters
package repository

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/v1adhope/music-artist-service/internal/entities"
	"github.com/v1adhope/music-artist-service/pkg/postgresql"
)

type ArtistRepo struct {
	*postgresql.Postgres
}

func NewArtist(d *postgresql.Postgres) *ArtistRepo {
	return &ArtistRepo{d}
}

type artistDto struct {
	name              string
	description       string
	website           string
	mounthlyListeners uint64
	email             string
	status            string
}

func (r *ArtistRepo) Get(ctx context.Context, id entities.ArtistId) (entities.Artist, error) {
	sql, args, err := r.Builder.Select(
		"name",
		"description",
		"website",
		"mounthly_listeners",
		"email",
	).From("artists").
		Where(squirrel.Eq{
			"artist_id": id.Get(),
		}).ToSql()
	if err != nil {
		return entities.Artist{}, fmt.Errorf("repository: artist: get: can't build query: %w", err)
	}

	dto := artistDto{}

	if err := r.Pool.QueryRow(ctx, sql, args...).Scan(&dto.name, &dto.description, &dto.website, &dto.mounthlyListeners, &dto.email); err != nil {
		return entities.Artist{}, fmt.Errorf("repository: artist: get: can't map data: %w", err)
	}

	artist := entities.Artist{}
	artist.SetName(dto.name)
	artist.SetDescription(dto.description)
	artist.SetWebsite(dto.website)
	artist.SetMounthlyListeners(dto.mounthlyListeners)
	artist.SetEmaiil(dto.email)
	artist.SetStatus(dto.mounthlyListeners)

	return artist, nil
}

func (r *ArtistRepo) GetAll(ctx context.Context) ([]entities.Artist, error) {
	sql, args, err := r.Builder.Select(
		"name",
		"description",
		"website",
		"mounthly_listeners",
		"email",
	).From("artists").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("repository: artist: getAll: can't build query: %w", err)
	}

	artists := make([]entities.Artist, 0)
	dto := artistDto{}

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("repository: artist: getAll: can't fetch data: %w", err)
	}

	_, err = pgx.ForEachRow(rows, []any{&dto.name, &dto.description, &dto.website, &dto.mounthlyListeners, &dto.email}, func() error {
		artist := entities.Artist{}
		artist.SetName(dto.name)
		artist.SetDescription(dto.description)
		artist.SetWebsite(dto.website)
		artist.SetMounthlyListeners(dto.mounthlyListeners)
		artist.SetEmaiil(dto.email)
		artist.SetStatus(dto.mounthlyListeners)

		artists = append(artists, artist)

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("repository: artist: getAll: can't map data: %w", err)
	}

	return artists, nil
}

type artisctCreateDto struct {
	id string
}

func (r *ArtistRepo) Create(ctx context.Context, artist entities.Artist) (entities.ArtistId, error) {
	sql, args, err := r.Builder.Insert("artists").
		Columns(
			"name",
			"description",
			"website",
			"mounthly_listeners",
			"email",
		).Values(
		artist.GetName().String(),
		artist.GetDescription().String(),
		artist.GetWebsite().String(),
		artist.GetMounthlyListeners(),
		artist.GetEmail().String(),
	).Suffix("returning \"artist_id\"").
		ToSql()
	if err != nil {
		return entities.ArtistId{}, fmt.Errorf("repository: artist: create: can't build query: %w", err)
	}

	dto := artisctCreateDto{}

	if err := r.Pool.QueryRow(ctx, sql, args...).Scan(&dto.id); err != nil {
		return entities.ArtistId{}, fmt.Errorf("repository: artist: create: can't map: %w", err)
	}

	id := entities.ArtistId{}
	id.Set(dto.id)

	return id, nil
}

func (r *ArtistRepo) Delete(ctx context.Context, id entities.ArtistId) error {
	sql, args, err := r.Builder.Delete("artists").
		Where(squirrel.Eq{
			"artist_id": id.Get(),
		}).ToSql()
	if err != nil {
		return fmt.Errorf("repository: artist: delete: can't build query: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("repository: artist: delete: can't exec query: %w", err)
	}

	return nil
}

func (r *ArtistRepo) Replace(ctx context.Context, artist entities.Artist) error {
	valuesByColumns := squirrel.Eq{
		"name":               artist.GetEmail().String(),
		"description":        artist.GetDescription().String(),
		"website":            artist.GetWebsite().String(),
		"mounthly_listeners": artist.GetMounthlyListeners(),
		"email":              artist.GetEmail().String(),
	}

	sql, args, err := r.Builder.Update("artists").
		SetMap(valuesByColumns).
		Where(squirrel.Eq{
			"artist_id": artist.GetId(),
		}).
		ToSql()
	if err != nil {
		return fmt.Errorf("repository: artist: replace: can't build query: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("repository: artist: replace: can't exec query: %w", err)
	}

	return nil
}
