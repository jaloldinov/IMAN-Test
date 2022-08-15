package service

import (
	"context"

	"github.com/jaloldinov/IMAN-Updated/post_service/config"
	"github.com/jaloldinov/IMAN-Updated/post_service/genproto/second_service"
	"github.com/jaloldinov/IMAN-Updated/post_service/pkg/logger"
	"github.com/jaloldinov/IMAN-Updated/post_service/storage"
)

type secondService struct {
	cfg  config.Config
	log  logger.LoggerI
	strg storage.StorageI
	second_service.UnimplementedSecondServiceServer
}

func NewSecondService(cfg config.Config, log logger.LoggerI, strg storage.StorageI) *secondService {
	return &secondService{
		cfg:  cfg,
		log:  log,
		strg: strg,
	}
}

// GetAll returns all posts with given offset and limit
func (s *secondService) ListPosts(ctx context.Context, req *second_service.ListPostsRequest) (*second_service.ListPostsResponse, error) {
	resp, err := s.strg.Post().ListPosts(req.Limit, req.Offset, req.Search)
	if err != nil {
		s.log.Error("ListPosts", logger.Any("req", req), logger.Error(err))
	}

	return resp, err
}

// Get returns post by id
func (s *secondService) GetPost(ctx context.Context, req *second_service.GetPostRequest) (*second_service.Post, error) {
	resp, err := s.strg.Post().GetPost(int(req.PostId))
	if err != nil {
		s.log.Error("GetPost", logger.Any("req", req), logger.Error(err))
	}

	return resp, err
}

// UpdatePost updates post by id
func (s *secondService) UpdatePost(ctx context.Context, req *second_service.UpdatePostRequest) (*second_service.Result, error) {
	resp, err := s.strg.Post().UpdatePost(req)
	if err != nil {
		s.log.Error("UpdatePost", logger.Any("req", req), logger.Error(err))
	}

	return &second_service.Result{
		Message: resp.String(),
	}, nil
}

// DeletePost deletes post by given id
func (s *secondService) DeletePost(ctx context.Context, req *second_service.GetPostRequest) (*second_service.Result, error) {
	resp, err := s.strg.Post().DeletePost(int(req.PostId))
	if err != nil {
		s.log.Error("DeletePost", logger.Any("req", req), logger.Error(err))
	}

	return &second_service.Result{
		Message: resp.String(),
	}, nil
}
