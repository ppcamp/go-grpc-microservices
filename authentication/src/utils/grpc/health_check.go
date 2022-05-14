package grpc

import (
	"context"

	"google.golang.org/grpc/health/grpc_health_v1"
)

type MyGrpcService struct {
	grpc_health_v1.UnimplementedHealthServer
}

func (m *MyGrpcService) Check(_ context.Context, _ *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	return &grpc_health_v1.HealthCheckResponse{Status: grpc_health_v1.HealthCheckResponse_SERVING}, nil
}

func (m *MyGrpcService) Watch(_ *grpc_health_v1.HealthCheckRequest, _ grpc_health_v1.Health_WatchServer) error {
	panic("not implemented")
}
