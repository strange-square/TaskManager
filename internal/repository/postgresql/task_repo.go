package postgresql

import (
	"context"
	"database/sql"

	"main.go/internal/db"
	"main.go/internal/repository"
)

type TasksRepo struct {
	db db.DBops
}

func NewTasks(db db.DBops) *TasksRepo {
	return &TasksRepo{db: db}
}

func (r *TasksRepo) Add(ctx context.Context, task *repository.Task) (int64, error) {
	var id int64
	err := r.db.ExecQueryRow(ctx,
		`INSERT INTO tasks(name, project_id) VALUES ($1, $2) RETURNING id`,
		task.Name, task.ProjectID).Scan(&id)
	return id, err
}

func (r *TasksRepo) GetById(ctx context.Context, id int64) (*repository.Task, error) {
	var task repository.Task
	err := r.db.Get(ctx, &task,
		"SELECT id,name,project_id,created_at,updated_at FROM tasks WHERE id=$1", id)
	if err == sql.ErrNoRows || err != nil && err.Error() == "scanning one: no rows in result set" {
		return nil, repository.ErrObjectNotFound
	}
	return &task, err
}

func (r *TasksRepo) List(ctx context.Context) ([]*repository.Task, error) {
	tasks := make([]*repository.Task, 0)
	err := r.db.Select(ctx, &tasks, "SELECT id,name,project_id,created_at,updated_at FROM tasks")
	return tasks, err
}

func (r *TasksRepo) Update(ctx context.Context, task *repository.Task) (bool, error) {
	result, err := r.db.Exec(ctx,
		"UPDATE tasks SET name = $1 WHERE id = $2", task.Name, task.ID)
	if result.RowsAffected() == 0 {
		return false, repository.ErrObjectNotFound
	}
	return true, err
}

func (r *TasksRepo) DeleteById(ctx context.Context, id int64) (bool, error) {
	result, err := r.db.Exec(ctx, "DELETE FROM tasks WHERE id=$1", id)
	if result.RowsAffected() == 0 {
		return false, repository.ErrObjectNotFound
	}
	return true, err
}
