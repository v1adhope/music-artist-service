package v1

import (
	"fmt"
	"log"
	"net"

	"github.com/v1adhope/music-artist-service/internal/usecases"
	artistv1 "github.com/v1adhope/music-artist-service/pkg/api/proto/v1"
	"github.com/v1adhope/music-artist-service/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

type (
	Config struct {
		Socket       string
		WithTls      bool
		TlsFilePaths TlsFilePaths
	}

	TlsFilePaths struct {
		Cert string
		Key  string
	}
)

type Server struct {
	*grpc.Server
	socket string
}

func Build(cfg Config, uc *usecases.Usecases, logger logger.Logger) (*Server, error) {
	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(errorHandler),
	}

	s := &grpc.Server{}

	if cfg.WithTls {
		creds, err := credentials.NewServerTLSFromFile(cfg.TlsFilePaths.Cert, cfg.TlsFilePaths.Key)
		if err != nil {
			return nil, fmt.Errorf("v1: controller: can't load creds: %w", err)
		}

		opts = append(opts, grpc.Creds(creds))

		s = grpc.NewServer(opts...)
	} else {
		s = grpc.NewServer(opts...)
	}

	artistSrv := newArtistServer(uc.Artist)

	reflection.Register(s)
	artistv1.RegisterArtistServer(s, artistSrv)

	return &Server{
		Server: s,
		socket: cfg.Socket,
	}, nil
}

func (s *Server) Run() {
	lis, err := net.Listen("tcp", s.socket)
	if err != nil {
		log.Fatalf("v1: controller: can't listen %s: %v", s.socket, err)
	}

	if err := s.Serve(lis); err != nil {
		log.Fatalf("v1: controller: can't start grpc server: %v", err)
	}
}

func (s *Server) Shutdown() {
	s.GracefulStop()
}
