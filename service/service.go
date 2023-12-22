package service

import (
	"fmt"
	"github.com/guneyin/sbda-product-category-service/config"
	sdk "github.com/guneyin/sbda-sdk"
	pb "github.com/guneyin/sbda-sdk/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"net"
	"os"
)

type Service struct {
	cfg     *config.Config
	log     *sdk.Logger
	handler *grpcHandler
	ds      *sdk.DiscoveryService
}

func NewService() (*Service, error) {
	cfg, err := config.GetConfig()
	if err != nil {
		return nil, err
	}

	hnd, err := newGrpcHandler(cfg)
	if err != nil {
		return nil, err
	}

	ds, err := sdk.NewDiscoveryService(cfg.DiscoverySvcAddr)
	if err != nil {
		return nil, err
	}

	return &Service{
		cfg:     cfg,
		log:     sdk.NewLogger(),
		handler: hnd,
		ds:      ds,
	}, nil
}

var _ sdk.IService = (*Service)(nil)

func (as Service) Register() error {
	return as.ds.RegisterService(as)
}

func (as Service) UnRegister() error {
	as.log.Warn("no-op Service/UnRegister")

	return nil
}

func (as Service) ServiceInfo() *sdk.ServiceInfo {
	hostName, _ := os.Hostname()

	si := &sdk.ServiceInfo{
		ID:       sdk.ProductCategoryService.String(),
		Name:     sdk.ProductCategoryService.String(),
		IP:       hostName,
		Port:     as.cfg.RpcPort,
		Protocol: sdk.ServiceProtocolGrpc,
	}

	return si
}

func (as Service) Serve() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", as.cfg.RpcPort))
	if err != nil {
		return err
	}

	as.log.Info(fmt.Sprintf("product category service running on %s", lis.Addr().String()))

	grpcServer := grpc.NewServer()

	pb.RegisterCategoryServiceServer(grpcServer, as.handler)
	grpc_health_v1.RegisterHealthServer(grpcServer, as.handler)

	if err = grpcServer.Serve(lis); err != nil {
		return err
	}

	return nil
}
