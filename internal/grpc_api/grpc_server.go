package grpc_api

import (
	"context"

	pb "main.go/internal/proto"
	"main.go/internal/repository"
)

type GRPCServer struct {
	pb.UnimplementedTaskManagerServiceServer
	Tasks    repository.TasksRepo
	Projects repository.ProjectsRepo
}

func NewGRPCServer(ctx context.Context, tasks repository.TasksRepo, projects repository.ProjectsRepo) *GRPCServer {
	return &GRPCServer{
		Tasks:    tasks,
		Projects: projects,
	}
}
