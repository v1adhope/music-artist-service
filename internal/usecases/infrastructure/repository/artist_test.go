// INFO: tests not isolated and weak (focus for testing through containers)
package repository_test

import (
	"context"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/v1adhope/music-artist-service/internal/entities"
	"github.com/v1adhope/music-artist-service/internal/testhelpers"
	"github.com/v1adhope/music-artist-service/internal/usecases/infrastructure/repository"
	"github.com/v1adhope/music-artist-service/internal/usecases/infrastructure/validation"
	"github.com/v1adhope/music-artist-service/pkg/postgresql"
)

type ArtistSuite struct {
	suite.Suite
	pgContainer *testhelpers.PostgresContainer
	repo        repository.ArtistReposer
	validator   validation.Validater
	ctx         context.Context
}

func (suite *ArtistSuite) SetupSuite() {
	suite.ctx = context.Background()

	pgC, err := testhelpers.BuildPostgresContainer(suite.ctx)
	if err != nil {
		log.Fatal(err)
	}

	suite.pgContainer = pgC

	driver, err := postgresql.Build(suite.ctx, postgresql.Config{
		ConnStr: suite.pgContainer.ConnStr,
	})
	if err != nil {
		log.Fatal(err)
	}

	if err := testhelpers.Migrate("file://../../../../db/migrations", suite.pgContainer.ConnStr); err != nil {
		log.Fatal(err)
	}

	if err := testhelpers.Seed(suite.ctx, driver); err != nil {
		log.Fatal(err)
	}

	suite.repo = repository.NewArtist(driver)

	suite.validator = validation.New()
}

func (suite *ArtistSuite) TearDownSuite() {
	if err := suite.pgContainer.Terminate(suite.ctx); err != nil {
		log.Fatalf("can't terminate postgres container: %s", err)
	}
}

func (suite *ArtistSuite) Test4Create() {
	t := suite.T()

	tcs := []struct {
		key   string
		input entities.Artist
	}{
		{
			key:   "case 1",
			input: testhelpers.GetNotExistingArtist(),
		},
	}

	for _, tc := range tcs {
		t.Run("", func(t *testing.T) {
			id, err := suite.repo.Create(suite.ctx, tc.input)
			ok := suite.validator.IsValidUuid(id.Get())

			assert.NoError(t, err, tc.key)
			assert.True(t, ok, tc.key)
		})
	}
}

func (suite *ArtistSuite) Test2Get() {
	t := suite.T()

	id := entities.ArtistId{}
	id.Set("1ef58be4-58cf-6bf0-bff6-58a65fd20958")

	tcs := []struct {
		key      string
		input    entities.ArtistId
		expected entities.Artist
	}{
		{
			key:      "case 1",
			input:    id,
			expected: testhelpers.GetExistingArtists()[0],
		},
	}

	for _, tc := range tcs {
		t.Run("", func(t *testing.T) {

			artist, err := suite.repo.Get(suite.ctx, id)

			assert.NoError(t, err, tc.key)
			assert.Equal(t, artist.GetName(), tc.expected.GetName(), tc.key)
			assert.Equal(t, artist.GetDescription(), tc.expected.GetDescription(), tc.key)
			assert.Equal(t, artist.GetWebsite(), tc.expected.GetWebsite(), tc.key)
			assert.Equal(t, artist.GetMounthlyListeners(), tc.expected.GetMounthlyListeners(), tc.key)
			assert.Equal(t, artist.GetEmail(), tc.expected.GetEmail(), tc.key)
			assert.Equal(t, artist.GetStatus(), tc.expected.GetStatus(), tc.key)
		})
	}
}

func (suite *ArtistSuite) Test1GetAll() {
	t := suite.T()

	tcs := []struct {
		key      string
		expected []entities.Artist
	}{
		{
			key:      "case 1",
			expected: testhelpers.GetExistingArtists(),
		},
	}

	for _, tc := range tcs {
		t.Run("", func(t *testing.T) {
			artists, err := suite.repo.GetAll(suite.ctx)

			assert.NoError(t, err, tc.key)
			for i, artist := range artists {
				assert.Equal(t, artist.GetName(), tc.expected[i].GetName(), tc.key)
				assert.Equal(t, artist.GetDescription(), tc.expected[i].GetDescription(), tc.key)
				assert.Equal(t, artist.GetWebsite(), tc.expected[i].GetWebsite(), tc.key)
				assert.Equal(t, artist.GetMounthlyListeners(), tc.expected[i].GetMounthlyListeners(), tc.key)
				assert.Equal(t, artist.GetEmail(), tc.expected[i].GetEmail(), tc.key)
				assert.Equal(t, artist.GetStatus(), tc.expected[i].GetStatus(), tc.key)
			}
		})
	}
}

func (suite *ArtistSuite) Test5Delete() {
	t := suite.T()

	id := entities.ArtistId{}
	id.Set("1ef58be4-58da-60a0-84fa-f187bb3f5677")

	tcs := []struct {
		key   string
		input entities.ArtistId
	}{
		{
			key:   "case 1",
			input: id,
		},
	}

	for _, tc := range tcs {
		t.Run("", func(t *testing.T) {
			err := suite.repo.Delete(suite.ctx, tc.input)

			assert.NoError(t, err, tc.key)
		})
	}
}

func (suite *ArtistSuite) Test3Replace() {
	t := suite.T()

	artistNF := testhelpers.GetExistingArtists()[0]

	artistNF.SetName("NF")
	artistNF.SetDescription("There is new desc")
	artistNF.SetWebsite("https://facebook.com/nfrealmusci")
	artistNF.SetMounthlyListeners(13899500)
	artistNF.SetEmaiil("nf@example.com")

	tcs := []struct {
		key   string
		input entities.Artist
	}{
		{
			key:   "case 1",
			input: artistNF,
		},
	}

	for _, tc := range tcs {
		t.Run("", func(t *testing.T) {
			err := suite.repo.Replace(suite.ctx, tc.input)

			assert.NoError(t, err, tc.key)
		})
	}
}

func TestArtistSuite(t *testing.T) {
	suite.Run(t, new(ArtistSuite))
}
