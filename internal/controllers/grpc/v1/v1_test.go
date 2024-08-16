package v1_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	v1 "github.com/v1adhope/music-artist-service/internal/controllers/grpc/v1"
	"github.com/v1adhope/music-artist-service/internal/testhelpers"
	artistv1 "github.com/v1adhope/music-artist-service/pkg/api/proto/v1"
)

type artistSuite struct {
	suite.Suite
	ctx    context.Context
	server *v1.Server
	client artistv1.ArtistClient
}

func (suite *artistSuite) SetupSuite() {
	t := suite.T()

	suite.ctx = context.Background()

	suite.client = testhelpers.MustGrpcArtistClient(t)

	suite.server = testhelpers.MustGrpcServer(t)

	go suite.server.Run()
}

func (suite *artistSuite) TearDownSuite() {
	suite.server.GracefulStop()
}

func (suite *artistSuite) TestCreate() {
	t := suite.T()

	tcs := []struct {
		key      string
		input    *artistv1.CreateArtistReq
		expected *artistv1.ArtistIdType
	}{
		{
			key: "case 1",
			input: &artistv1.CreateArtistReq{
				Data: &artistv1.ArtistType{
					Name:              "Eminem",
					Description:       "One of the greatest rappers of his generation",
					Website:           "https://facebook.com/Eminem",
					MounthlyListeners: 13899500,
					Email:             "nf@example.com",
				},
			},
			expected: &artistv1.ArtistIdType{
				Id: "1ef58be4-58cf-6bf0-bff6-58a65fd20958",
			},
		},
	}

	for _, tc := range tcs {
		t.Run("", func(t *testing.T) {
			reply, err := suite.client.Create(suite.ctx, tc.input)
			assert.NoError(t, err, tc.key)
			assert.Equal(t, tc.expected, reply.GetData(), tc.key)
		})
	}
}

func (suite *artistSuite) TestGet() {
	t := suite.T()

	artist := testhelpers.GetExistingArtists()[0]
	case1Status := artist.GetStatus().String()
	case1Id := artist.GetId()

	tcs := []struct {
		key      string
		input    *artistv1.GetArtistReq
		expected *artistv1.ArtistType
	}{
		{
			key: "case 1",
			input: &artistv1.GetArtistReq{
				Data: &artistv1.ArtistIdType{
					Id: "1ef58be4-58cf-6bf0-bff6-58a65fd20958",
				},
			},
			expected: &artistv1.ArtistType{
				Id:                &case1Id,
				Name:              artist.GetName().String(),
				Description:       artist.GetDescription().String(),
				Website:           artist.GetWebsite().String(),
				MounthlyListeners: artist.GetMounthlyListeners(),
				Email:             artist.GetEmail().String(),
				Status:            &case1Status,
			},
		},
	}

	for _, tc := range tcs {
		t.Run("", func(t *testing.T) {
			reply, err := suite.client.Get(suite.ctx, tc.input)

			assert.NoError(t, err, tc.key)
			assert.Equal(t, tc.expected, reply.GetData(), tc.key)
		})
	}
}

func (suite *artistSuite) TestGetAll() {
	t := suite.T()

	artists := testhelpers.GetExistingArtists()
	case1artist1Status := artists[0].GetStatus().String()
	case1artist2Status := artists[1].GetStatus().String()
	case1artist3Status := artists[2].GetStatus().String()
	case1artist1Id := artists[0].GetId()
	case1artist2Id := artists[1].GetId()
	case1artist3Id := artists[2].GetId()

	tcs := []struct {
		key      string
		expected []*artistv1.ArtistType
	}{
		{
			key: "case 1",
			expected: []*artistv1.ArtistType{
				{
					Id:                &case1artist1Id,
					Name:              artists[0].GetName().String(),
					Description:       artists[0].GetDescription().String(),
					Website:           artists[0].GetWebsite().String(),
					MounthlyListeners: artists[0].GetMounthlyListeners(),
					Email:             artists[0].GetEmail().String(),
					Status:            &case1artist1Status,
				},
				{
					Id:                &case1artist2Id,
					Name:              artists[1].GetName().String(),
					Description:       artists[1].GetDescription().String(),
					Website:           artists[1].GetWebsite().String(),
					MounthlyListeners: artists[1].GetMounthlyListeners(),
					Email:             artists[1].GetEmail().String(),
					Status:            &case1artist2Status,
				},
				{
					Id:                &case1artist3Id,
					Name:              artists[2].GetName().String(),
					Description:       artists[2].GetDescription().String(),
					Website:           artists[2].GetWebsite().String(),
					MounthlyListeners: artists[2].GetMounthlyListeners(),
					Email:             artists[2].GetEmail().String(),
					Status:            &case1artist3Status,
				},
			},
		},
	}

	for _, tc := range tcs {
		t.Run("", func(t *testing.T) {
			reply, err := suite.client.GetAll(suite.ctx, nil)

			assert.NoError(t, err, tc.key)
			assert.Equal(t, tc.expected, reply.GetData(), tc.key)
		})
	}
}

func (suite *artistSuite) TestReplace() {
	t := suite.T()

	artist := testhelpers.GetExistingArtists()[0]
	artistStatus := artist.GetStatus().String()
	artistId := artist.GetId()

	tcs := []struct {
		key   string
		input *artistv1.ReplaceArtistReq
	}{
		{
			key: "case 1",
			input: &artistv1.ReplaceArtistReq{
				Data: &artistv1.ArtistType{
					Id:                &artistId,
					Name:              artist.GetName().String(),
					Description:       artist.GetDescription().String(),
					Website:           artist.GetWebsite().String(),
					MounthlyListeners: artist.GetMounthlyListeners(),
					Email:             artist.GetEmail().String(),
					Status:            &artistStatus,
				},
			},
		},
	}

	for _, tc := range tcs {
		t.Run("", func(t *testing.T) {
			reply, err := suite.client.Replace(suite.ctx, tc.input)

			assert.NoError(t, err, tc.key)
			assert.Empty(t, reply.String(), tc.key)
		})
	}
}

func (suite *artistSuite) TestDelete() {
	t := suite.T()

	tcs := []struct {
		key   string
		input *artistv1.DeleteArtistReq
	}{
		{
			key: "case 1",
			input: &artistv1.DeleteArtistReq{
				Data: &artistv1.ArtistIdType{
					Id: "1ef58be4-58cf-6bf0-bff6-58a65fd20958",
				},
			},
		},
	}

	for _, tc := range tcs {
		t.Run("", func(t *testing.T) {
			reply, err := suite.client.Delete(suite.ctx, tc.input)

			assert.NoError(t, err, tc.key)
			assert.Empty(t, reply.String(), tc.key)
		})
	}
}

func TestArtistSuite(t *testing.T) {
	suite.Run(t, new(artistSuite))
}
