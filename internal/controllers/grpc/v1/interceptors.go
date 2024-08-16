package v1

import (
	"context"
	"errors"

	"github.com/v1adhope/music-artist-service/internal/entities"
	"github.com/v1adhope/music-artist-service/internal/objectvalues"
	"github.com/v1adhope/music-artist-service/internal/usecases"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func errorHandler(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	m, err := handler(ctx, req)
	if err != nil {
		switch {
		case errors.Is(err, usecases.ErrNotValidUuid),
			errors.Is(err, objectvalues.ErrNotValidEmail),
			errors.Is(err, objectvalues.ErrNotValidWebsite),
			errors.Is(err, objectvalues.ErrNotValidArtistName),
			errors.Is(err, objectvalues.ErrNotValidArtistDescription):
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}

		if errors.Is(err, entities.ErrNoContent) {
			return nil, status.Error(codes.NotFound, err.Error())
		}

		return nil, status.Error(codes.Internal, "Internal error")
	}

	return m, nil
}
