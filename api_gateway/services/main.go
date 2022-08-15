package services

import (
	"fmt"

	"github.com/jaloldinov/IMAN-Updated/api_gateway/config"
	"github.com/jaloldinov/IMAN-Updated/api_gateway/genproto/first_service"
	"github.com/jaloldinov/IMAN-Updated/api_gateway/genproto/second_service"

	"google.golang.org/grpc"
)

type ServiceManager interface {
	FirstService() first_service.FirstServiceClient
	SecondService() second_service.SecondServiceClient
}

type grpcClients struct {
	firstService  first_service.FirstServiceClient
	secondService second_service.SecondServiceClient
}

func NewGrpcClients(conf *config.Config) (ServiceManager, error) {

	connFirstService, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.FirstServiceHost, conf.FirstServicePort),
		grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	connSecondService, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.SecondServiceHost, conf.SecondServicePort),
		grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &grpcClients{
		firstService:  first_service.NewFirstServiceClient(connFirstService),
		secondService: second_service.NewSecondServiceClient(connSecondService),
	}, nil
}

func (g *grpcClients) FirstService() first_service.FirstServiceClient {
	return g.firstService
}

func (g *grpcClients) SecondService() second_service.SecondServiceClient {
	return g.secondService
}
