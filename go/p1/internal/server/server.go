package server

import (
	"context"
	"fmt"
	"log"

	api "example.com/at/api/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type grpcServer struct {
	api.UnimplementedActivity_LogServer
	Activities *Activities
}

func NewGRPCServer() (*grpc.Server, grpcServer) {
	var acc *Activities
	var err error
	if acc, err = NewActivities(); err != nil {
		log.Fatal(err)
	}
	gsrv := grpc.NewServer()
	srv := grpcServer{
		Activities: acc,
	}
	api.RegisterActivity_LogServer(gsrv, &srv)
	return gsrv, srv
}

func (s *grpcServer) Insert(ctx context.Context, activity *api.Activity) (*api.InsertResponse, error) {
	id, err := s.Activities.Insert(activity)
	if err != nil {
		return nil, fmt.Errorf("internal Error: %w", err)
	}
	res := api.InsertResponse{Id: int32(id)}
	return &res, nil
}

func (s *grpcServer) List(ctx context.Context, req *api.ListRequest) (*api.Activities, error) {
	activities, err := s.Activities.List(int(req.Offset))
	if err != nil {
		return nil, fmt.Errorf("internal Error: %w", err)
	}
	return &api.Activities{Activities: activities}, nil
}

func (s *grpcServer) Retrieve(ctx context.Context, req *api.RetrieveRequest) (*api.Activity, error) {
	resp, err := s.Activities.Retrieve(int(req.Id))
	if err == ErrIDNotFound {
		return nil, status.Error(codes.NotFound, "id was not found")
	}
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}
