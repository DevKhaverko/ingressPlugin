package ingress

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"ingressPlugin/proto/api"
)

type IngressPluginService interface {
	CreateOrChangeRoute(context.Context, string) (bool, error)
}

type serverAPI struct {
	api.UnimplementedIngressPluginServer
	service IngressPluginService
}

func RegisterServerAPI(gRPC *grpc.Server, ingressService IngressPluginService) {
	api.RegisterIngressPluginServer(gRPC, &serverAPI{service: ingressService})
}

func (s *serverAPI) CreateOrChangeRoute(ctx context.Context, req *api.AllocID) (*api.Response, error) {
	if req.ID == "" {
		return nil, status.Error(codes.InvalidArgument, "allocID must be set")
	}

	isChanged, err := s.service.CreateOrChangeRoute(ctx, req.GetID())
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to change or create")
	}
	return &api.Response{Result: isChanged}, nil
}
