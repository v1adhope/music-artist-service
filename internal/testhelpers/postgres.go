package testhelpers

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"github.com/v1adhope/music-artist-service/internal/entities"
	"github.com/v1adhope/music-artist-service/pkg/postgresql"
)

type PostgresContainer struct {
	*postgres.PostgresContainer
	ConnStr string
}

const (
	pgImage        = "docker.io/postgres:16.4"
	pgDataBaseName = "music_artist_service"
	pgUsername     = "rat"
	pgPassword     = "secret"
)

func BuildPostgresContainer(ctx context.Context) (*PostgresContainer, error) {
	pgC, err := postgres.Run(ctx,
		pgImage,
		postgres.WithDatabase(pgDataBaseName),
		postgres.WithUsername(pgUsername),
		postgres.WithPassword(pgPassword),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").WithOccurrence(2).WithStartupTimeout(5*time.Second),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to start container: %w", err)
	}

	connStr, err := pgC.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		return nil, fmt.Errorf("can't get connStr: %w", err)
	}

	return &PostgresContainer{
		PostgresContainer: pgC,
		ConnStr:           connStr,
	}, nil
}

func Migrate(relPath, connStr string) error {
	m, err := migrate.New(relPath, connStr)
	if err != nil {
		return fmt.Errorf("can't get connection to migrate : %w", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("can't migrate: %w", err)
	}

	return nil
}

func Seed(ctx context.Context, d *postgresql.Postgres) error {
	artists := GetExistingArtists()

	sql, args, err := d.Builder.Insert("artists").
		Columns(
			"artist_id",
			"name",
			"description",
			"website",
			"mounthly_listeners",
			"email",
		).Values(
		artists[0].GetId(),
		artists[0].GetName().String(),
		artists[0].GetDescription().String(),
		artists[0].GetWebsite().String(),
		artists[0].GetMounthlyListeners(),
		artists[0].GetEmail().String(),
	).Values(
		artists[1].GetId(),
		artists[1].GetName().String(),
		artists[1].GetDescription().String(),
		artists[1].GetWebsite().String(),
		artists[1].GetMounthlyListeners(),
		artists[1].GetEmail().String(),
	).Values(
		artists[2].GetId(),
		artists[2].GetName().String(),
		artists[2].GetDescription().String(),
		artists[2].GetWebsite().String(),
		artists[2].GetMounthlyListeners(),
		artists[2].GetEmail().String(),
	).ToSql()
	if err != nil {
		return fmt.Errorf("can't build seed query: %w", err)
	}

	if _, err := d.Pool.Exec(ctx, sql, args...); err != nil {
		return fmt.Errorf("can't insert seed data: %w", err)
	}

	return nil
}

func GetExistingArtists() []entities.Artist {
	artists := make([]entities.Artist, 0, 3)

	artistNF := entities.Artist{}
	artistNF.SetId("1ef58be4-58cf-6bf0-bff6-58a65fd20958")
	artistNF.SetName("NF")
	artistNF.SetDescription("Raps with raw grit and emotional authenticity")
	artistNF.SetWebsite("https://facebook.com/nfrealmusci")
	artistNF.SetMounthlyListeners(13899500)
	artistNF.SetEmaiil("nf@example.com")
	artistNF.SetStatus(13899500)

	artists = append(artists, artistNF)

	artistEminem := entities.Artist{}
	artistEminem.SetId("1ef58be4-58d9-6de0-a863-f5d46f847dad")
	artistEminem.SetName("Eminem")
	artistEminem.SetDescription("One of the greatest rappers of his generation")
	artistEminem.SetWebsite("https://facebook.com/Eminem")
	artistEminem.SetMounthlyListeners(83563706)
	artistEminem.SetEmaiil("eminem@fhcustomercare.com")
	artistEminem.SetStatus(83563706)

	artists = append(artists, artistEminem)

	artistAltJ := entities.Artist{}
	artistAltJ.SetId("1ef58be4-58da-60a0-84fa-f187bb3f5677")
	artistAltJ.SetName("alt-J")
	artistAltJ.SetDescription("The dream was recorder from August 2020 until June 2021")
	artistAltJ.SetWebsite("https://facebook.com/altJ.band")
	artistAltJ.SetMounthlyListeners(9518740)
	artistAltJ.SetEmaiil("altj@example.com")
	artistAltJ.SetStatus(9518740)

	artists = append(artists, artistAltJ)

	return artists
}

func GetNotExistingArtist() entities.Artist {
	artist := entities.Artist{}
	artist.SetName("Green Day")
	artist.SetDescription("God's favorite band")
	artist.SetWebsite("https://facebook.com/greenday")
	artist.SetMounthlyListeners(33332230)
	artist.SetEmaiil("info@crushmusic.com")
	artist.SetStatus(33332230)

	return artist
}
