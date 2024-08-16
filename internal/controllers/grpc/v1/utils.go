package v1

import (
	"github.com/v1adhope/music-artist-service/internal/entities"
	artistv1 "github.com/v1adhope/music-artist-service/pkg/api/proto/v1"
)

func convertArtistEntityToArtistType(artist entities.Artist) *artistv1.ArtistType {
	id := artist.GetId()
	status := artist.GetStatus().String()

	return &artistv1.ArtistType{
		Id:                &id,
		Name:              artist.GetName().String(),
		Description:       artist.GetDescription().String(),
		Website:           artist.GetWebsite().String(),
		MounthlyListeners: artist.GetMounthlyListeners(),
		Email:             artist.GetEmail().String(),
		Status:            &status,
	}
}

func convertArtistCreateReqToEntity(req *artistv1.CreateArtistReq) (entities.Artist, error) {
	artist := entities.Artist{}
	artist.SetMounthlyListeners(req.Data.GetMounthlyListeners())

	if err := artist.SetName(req.Data.GetName()); err != nil {
		return entities.Artist{}, err
	}

	if err := artist.SetDescription(req.Data.GetDescription()); err != nil {
		return entities.Artist{}, err
	}

	if err := artist.SetWebsite(req.Data.GetWebsite()); err != nil {
		return entities.Artist{}, err
	}

	if err := artist.SetEmail(req.Data.GetEmail()); err != nil {
		return entities.Artist{}, err
	}

	return artist, nil
}

func convertArtistReplaceReqToEntity(req *artistv1.ReplaceArtistReq) (entities.Artist, error) {
	artist := entities.Artist{}
	artist.SetId(req.Data.GetId())
	artist.SetMounthlyListeners(req.Data.GetMounthlyListeners())

	if err := artist.SetName(req.Data.GetName()); err != nil {
		return entities.Artist{}, err
	}

	if err := artist.SetDescription(req.Data.GetDescription()); err != nil {
		return entities.Artist{}, err
	}

	if err := artist.SetWebsite(req.Data.GetWebsite()); err != nil {
		return entities.Artist{}, err
	}

	if err := artist.SetEmail(req.Data.GetEmail()); err != nil {
		return entities.Artist{}, err
	}

	return artist, nil
}

func convertArtistIdTypeToEntity(ait *artistv1.ArtistIdType) entities.ArtistId {
	id := entities.ArtistId{}
	id.Set(ait.GetId())

	return id
}

func convertFromEntityToArtistIdType(id entities.ArtistId) *artistv1.ArtistIdType {
	return &artistv1.ArtistIdType{
		Id: id.Get(),
	}
}
