package service

import (
	"context"
	"github.com/guneyin/sbda-product-category-service/config"
	"github.com/guneyin/sbda-product-category-service/repo"
	pb "github.com/guneyin/sbda-sdk/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"
	"log"
)

type grpcHandler struct {
	config *config.Config
	repo   *repo.Repo
	pb.UnimplementedCategoryServiceServer
}

func newGrpcHandler(cfg *config.Config) (*grpcHandler, error) {
	if cfg == nil {
		log.Fatal("invalid config 4")
	}

	r, err := repo.NewRepo(cfg)
	if err != nil {
		return nil, err
	}

	return &grpcHandler{
		config: cfg,
		repo:   r,
	}, nil
}

func (h *grpcHandler) CreateCategory(ctx context.Context, in *pb.CreateCategoryRequest) (*pb.CreateCategoryResponse, error) {
	return h.repo.Category.Create(ctx, in)
}

func (h *grpcHandler) Check(ctx context.Context, in *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	return &grpc_health_v1.HealthCheckResponse{Status: grpc_health_v1.HealthCheckResponse_SERVING}, nil
}

func (h *grpcHandler) Watch(in *grpc_health_v1.HealthCheckRequest, _ grpc_health_v1.Health_WatchServer) error {
	return status.Error(codes.Unimplemented, "unimplemented")
}
