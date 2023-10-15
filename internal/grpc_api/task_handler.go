package grpc_api

import (
	"context"
	"log"

	"google.golang.org/protobuf/types/known/timestamppb"
	"main.go/internal"
	pb "main.go/internal/proto"
	"main.go/internal/repository"
)

func (s *GRPCServer) CreateTask(ctx context.Context, in *pb.CreateTaskRequest) (*pb.CreateTaskResponse, error) {
	id, err := s.Tasks.Add(ctx, &repository.Task{Name: in.Name, ProjectID: in.ProjectId})
	if err != nil {
		log.Printf("error while adding task to database, err: [%s]", err.Error())
		return nil, err
	}

	internal.IncreaseRegCounter()

	return &pb.CreateTaskResponse{
		Id: id,
	}, nil
}

func (s *GRPCServer) UpdateTask(ctx context.Context, in *pb.UpdateTaskRequest) (*pb.UpdateTaskResponse, error) {
	ok, err := s.Tasks.Update(ctx, &repository.Task{ID: in.Task.Id,
		Name:      in.Task.Name,
		ProjectID: in.Task.ProjectId})

	if err != nil {
		log.Printf("error while updating task in database, err: [%s]", err.Error())
		return nil, err
	}

	return &pb.UpdateTaskResponse{
		Ok: ok,
	}, nil
}

func (s *GRPCServer) ListTask(ctx context.Context, _ *pb.Empty) (*pb.ListTaskResponse, error) {
	list, err := s.Tasks.List(ctx)
	if err != nil {
		log.Printf("error while getting list of tasks from database, err: [%s]", err.Error())
		return nil, err
	}

	result := make([]*pb.Task, 0, len(list))
	for _, m := range list {
		var updatedAt *timestamppb.Timestamp
		if m.UpdatedAt.Valid {
			updatedAt = timestamppb.New(m.UpdatedAt.Time)
		}
		result = append(result, &pb.Task{
			Id:        m.ID,
			Name:      m.Name,
			ProjectId: m.ProjectID,
			CreatedAt: timestamppb.New(m.CreatedAt),
			UpdatedAt: updatedAt,
		})
	}
	return &pb.ListTaskResponse{
		Tasks: result,
	}, nil
}

func (s *GRPCServer) GetTask(ctx context.Context, in *pb.GetTaskRequest) (*pb.GetTaskResponse, error) {
	task, err := s.Tasks.GetById(ctx, in.Id)
	if err != nil {
		log.Printf("error while getting task by id from database, err: [%s]", err.Error())
		return nil, err
	}

	var updatedAt *timestamppb.Timestamp
	if task.UpdatedAt.Valid {
		updatedAt = timestamppb.New(task.UpdatedAt.Time)
	}

	return &pb.GetTaskResponse{
		Task: &pb.Task{
			Id:        task.ID,
			Name:      task.Name,
			ProjectId: task.ProjectID,
			CreatedAt: timestamppb.New(task.CreatedAt),
			UpdatedAt: updatedAt,
		},
	}, nil
}

func (s *GRPCServer) DeleteTask(ctx context.Context, in *pb.DeleteTaskRequest) (*pb.DeleteTaskResponse, error) {
	ok, err := s.Tasks.DeleteById(ctx, int64(in.Id))

	if err != nil {
		log.Printf("error while deletion task by id from database, err: [%s]", err.Error())
		return nil, err
	}

	internal.IncreaseDeletedCounter()

	return &pb.DeleteTaskResponse{Ok: ok}, nil
}
