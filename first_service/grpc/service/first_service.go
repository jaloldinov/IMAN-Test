package service

import (
	"context"

	"github.com/jaloldinov/IMAN-Updated/first_service/config"
	"github.com/jaloldinov/IMAN-Updated/first_service/genproto/first_service"
	"github.com/jaloldinov/IMAN-Updated/first_service/pkg/logger"
	"github.com/jaloldinov/IMAN-Updated/first_service/storage"
	"google.golang.org/protobuf/types/known/emptypb"
)

type firstService struct {
	cfg  config.Config
	log  logger.LoggerI
	strg storage.StorageI
	first_service.UnimplementedFirstServiceServer
}

func NewDataService(cfg config.Config, log logger.LoggerI, strg storage.StorageI) *firstService {
	return &firstService{
		cfg:  cfg,
		log:  log,
		strg: strg,
	}
}

func (s firstService) InsertPosts(ctx context.Context, req *emptypb.Empty) (*first_service.InsertPostsResponse, error) {
	resp, err := s.strg.Data().InsertPosts()
	if err != nil {
		s.log.Error("InsertPosts", logger.Error(err))
		return nil, err
	}

	return &first_service.InsertPostsResponse{
		Message: resp,
	}, nil
}
