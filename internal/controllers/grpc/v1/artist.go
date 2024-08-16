package v1

import (
	"context"

	"github.com/v1adhope/music-artist-service/internal/usecases"
	artistv1 "github.com/v1adhope/music-artist-service/pkg/api/proto/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type artsitstServer struct {
	artistv1.UnimplementedArtistServer
	artistUc usecases.ArtistUsecaser
}

func newArtistServer(artistUc usecases.ArtistUsecaser) *artsitstServer {
	return &artsitstServer{
		artistUc: artistUc,
	}
}

func (srv *artsitstServer) Get(ctx context.Context, req *artistv1.GetArtistReq) (*artistv1.GetArtistReply, error) {
	id := convertArtistIdTypeToEntity(req.GetData())

	artist, err := srv.artistUc.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return &artistv1.GetArtistReply{
		Data: convertArtistEntityToArtistType(artist),
	}, nil
}

func (srv *artsitstServer) GetAll(ctx context.Context, req *emptypb.Empty) (*artistv1.GetAllArtistReply, error) {
	artists, err := srv.artistUc.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	respArtists := make([]*artistv1.ArtistType, 0)

	for _, artist := range artists {
		respArtists = append(respArtists, convertArtistEntityToArtistType(artist))
	}

	return &artistv1.GetAllArtistReply{
		Data: respArtists,
	}, nil

}

func (srv *artsitstServer) Create(ctx context.Context, req *artistv1.CreateArtistReq) (*artistv1.CreateArtistReply, error) {
	artist, err := convertArtistCreateReqToEntity(req)
	if err != nil {
		return nil, err
	}

	id, err := srv.artistUc.Create(ctx, artist)
	if err != nil {
		return nil, err
	}

	return &artistv1.CreateArtistReply{
		Data: convertFromEntityToArtistIdType(id),
	}, nil
}

func (srv *artsitstServer) Replace(ctx context.Context, req *artistv1.ReplaceArtistReq) (*emptypb.Empty, error) {
	artist, err := convertArtistReplaceReqToEntity(req)
	if err != nil {
		return nil, err
	}

	if err := srv.artistUc.Replace(ctx, artist); err != nil {
		return nil, err
	}

	return nil, nil
}

func (srv *artsitstServer) Delete(ctx context.Context, req *artistv1.DeleteArtistReq) (*emptypb.Empty, error) {
	id := convertArtistIdTypeToEntity(req.GetData())

	if err := srv.artistUc.Delete(ctx, id); err != nil {
		return nil, err
	}

	return nil, nil
}
