//go:generate mockgen -source ./repository.go -destination=./mocks/repository.go -package=mock_repository
package repository

import (
	"context"
	"errors"
)

var (
	ErrObjectNotFound = errors.New("object not found")
)

type TasksRepo interface {
	Add(ctx context.Context, task *Task) (int64, error)
	GetById(ctx context.Context, id int64) (*Task, error)
	List(ctx context.Context) ([]*Task, error)
	Update(ctx context.Context, task *Task) (bool, error)
	DeleteById(ctx context.Context, id int64) (bool, error)
}

type ProjectsRepo interface {
	Add(ctx context.Context, project *Project) (int64, error)
	GetById(ctx context.Context, id int64) (*Project, error)
	List(ctx context.Context) ([]*Project, error)
	Update(ctx context.Context, user *Project) (bool, error)
	DeleteById(ctx context.Context, id int64) (bool, error)
}
