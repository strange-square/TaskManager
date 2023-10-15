package grpc_api

import (
	"context"
	"log"

	"google.golang.org/protobuf/types/known/timestamppb"
	pb "main.go/internal/proto"
	"main.go/internal/repository"
)

func (s *GRPCServer) CreateProject(ctx context.Context, in *pb.CreateProjectRequest) (*pb.CreateProjectResponse, error) {
	id, err := s.Projects.Add(ctx, &repository.Project{Name: in.Name})
	if err != nil {
		log.Printf("error while adding project to database, err: [%s]", err.Error())
		return nil, err
	}

	return &pb.CreateProjectResponse{
		Id: id,
	}, nil
}

func (s *GRPCServer) UpdateProject(ctx context.Context, in *pb.UpdateProjectRequest) (*pb.UpdateProjectResponse, error) {
	ok, err := s.Projects.Update(ctx, &repository.Project{ID: in.Project.Id,
		Name: in.Project.Name})

	if err != nil {
		log.Printf("error while updating project in database, err: [%s]", err.Error())
		return nil, err
	}

	return &pb.UpdateProjectResponse{
		Ok: ok,
	}, nil
}

func (s *GRPCServer) ListProject(ctx context.Context, _ *pb.Empty) (*pb.ListProjectResponse, error) {
	list, err := s.Projects.List(ctx)
	if err != nil {
		log.Printf("error while getting list of projects from database, err: [%s]", err.Error())
		return nil, err
	}

	result := make([]*pb.Project, 0, len(list))
	for _, m := range list {
		var updatedAt *timestamppb.Timestamp
		if m.UpdatedAt.Valid {
			updatedAt = timestamppb.New(m.UpdatedAt.Time)
		}
		result = append(result, &pb.Project{
			Id:        m.ID,
			Name:      m.Name,
			CreatedAt: timestamppb.New(m.CreatedAt),
			UpdatedAt: updatedAt,
		})
	}
	return &pb.ListProjectResponse{
		Projects: result,
	}, nil
}

func (s *GRPCServer) GetProject(ctx context.Context, in *pb.GetProjectRequest) (*pb.GetProjectResponse, error) {
	project, err := s.Projects.GetById(ctx, in.Id)
	if err != nil {
		log.Printf("error while getting project by id from database, err: [%s]", err.Error())
		return nil, err
	}

	var updatedAt *timestamppb.Timestamp
	if project.UpdatedAt.Valid {
		updatedAt = timestamppb.New(project.UpdatedAt.Time)
	}

	return &pb.GetProjectResponse{
		Project: &pb.Project{
			Id:        project.ID,
			Name:      project.Name,
			CreatedAt: timestamppb.New(project.CreatedAt),
			UpdatedAt: updatedAt,
		},
	}, nil
}

func (s *GRPCServer) DeleteProject(ctx context.Context, in *pb.DeleteProjectRequest) (*pb.DeleteProjectResponse, error) {
	ok, err := s.Projects.DeleteById(ctx, int64(in.Id))

	if err != nil {
		log.Printf("error while deletion project by id from database, err: [%s]", err.Error())
		return nil, err
	}

	return &pb.DeleteProjectResponse{Ok: ok}, nil
}
