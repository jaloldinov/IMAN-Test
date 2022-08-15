package grpc

import (
	"github.com/jaloldinov/IMAN-Updated/first_service/config"
	"github.com/jaloldinov/IMAN-Updated/first_service/genproto/first_service"
	"github.com/jaloldinov/IMAN-Updated/first_service/grpc/service"
	"github.com/jaloldinov/IMAN-Updated/first_service/pkg/logger"
	"github.com/jaloldinov/IMAN-Updated/first_service/storage"
	"google.golang.org/grpc"
)

func SetUpServer(cfg config.Config, log logger.LoggerI, strg storage.StorageI) (grpcServer *grpc.Server) {
	grpcServer = grpc.NewServer()
	first_service.RegisterFirstServiceServer(grpcServer, service.NewDataService(cfg, log, strg))

	return
}
