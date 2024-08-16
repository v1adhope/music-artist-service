package testhelpers

import (
	"context"
	"testing"

	"github.com/stretchr/testify/mock"
	v1 "github.com/v1adhope/music-artist-service/internal/controllers/grpc/v1"
	"github.com/v1adhope/music-artist-service/internal/entities"
	"github.com/v1adhope/music-artist-service/internal/usecases"
	"github.com/v1adhope/music-artist-service/internal/usecases/infrastructure/repositories"
	"github.com/v1adhope/music-artist-service/internal/usecases/infrastructure/validation"
	reposmocks "github.com/v1adhope/music-artist-service/internal/usecases/mocks"
	artistv1 "github.com/v1adhope/music-artist-service/pkg/api/proto/v1"
	"github.com/v1adhope/music-artist-service/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	grpcSocket    = ":50111"
	loggerInfoLvl = "debug"
)

func MustGrpcArtistClient(t *testing.T) artistv1.ArtistClient {
	conn, err := grpc.NewClient(
		grpcSocket,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		conn.Close()
	})

	return artistv1.NewArtistClient(conn)
}

func MustGrpcServer(t *testing.T) *v1.Server {
	r := repositories.Repos{
		Artist: prepareArtistReposer(t),
	}

	v := validation.New()

	u := &usecases.Usecases{
		Artist: usecases.NewArtist(r.Artist, v),
	}

	l := logger.New(
		logger.Config{
			Level: loggerInfoLvl,
		},
	)

	s, err := v1.Build(
		v1.Config{
			Socket:  grpcSocket,
			WithTls: false,
		}, u, l)
	if err != nil {
		t.Fatal(err)
	}

	return s
}

func prepareArtistReposer(t *testing.T) *reposmocks.ArtistReposer {
	artistRepo := reposmocks.NewArtistReposer(t)

	artistRepo.On(
		"Create",
		mock.Anything,
		mock.Anything,
	).Return(
		func(ctx context.Context, artist entities.Artist) entities.ArtistId {
			id := entities.ArtistId{}
			id.Set("1ef58be4-58cf-6bf0-bff6-58a65fd20958")
			return id
		},
		func(ctx context.Context, artist entities.Artist) error {
			return nil
		},
	)

	artistRepo.On(
		"Delete",
		mock.Anything,
		mock.Anything,
	).Return(
		func(ctx context.Context, id entities.ArtistId) error {
			return nil
		},
	)

	artistRepo.On(
		"Get",
		mock.Anything,
		mock.Anything,
	).Return(
		func(ctx context.Context, id entities.ArtistId) entities.Artist {
			return GetExistingArtists()[0]
		},
		func(ctx context.Context, id entities.ArtistId) error {
			return nil
		},
	)

	artistRepo.On(
		"GetAll",
		mock.Anything,
	).Return(
		func(ctx context.Context) []entities.Artist {
			return GetExistingArtists()
		},
		func(ctx context.Context) error {
			return nil
		},
	)

	artistRepo.On(
		"Replace",
		mock.Anything,
		mock.Anything,
	).Return(nil)

	return artistRepo
}
