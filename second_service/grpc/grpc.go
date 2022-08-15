package grpc

import (
	"github.com/jaloldinov/IMAN-Updated/post_service/config"
	"github.com/jaloldinov/IMAN-Updated/post_service/genproto/second_service"
	"github.com/jaloldinov/IMAN-Updated/post_service/grpc/service"
	"github.com/jaloldinov/IMAN-Updated/post_service/pkg/logger"
	"github.com/jaloldinov/IMAN-Updated/post_service/storage"
	"google.golang.org/grpc"
)

func SetUpServer(cfg config.Config, log logger.LoggerI, strg storage.StorageI) (grpcServer *grpc.Server) {
	grpcServer = grpc.NewServer()

	second_service.RegisterSecondServiceServer(grpcServer, service.NewSecondService(cfg, log, strg))
	return
}
