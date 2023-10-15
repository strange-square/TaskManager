package postgresql

import (
	"context"
	"database/sql"

	"main.go/internal/db"
	"main.go/internal/repository"
)

type ProjectsRepo struct {
	db db.DBops
}

func NewProjects(db db.DBops) *ProjectsRepo {
	return &ProjectsRepo{db: db}
}

func (r *ProjectsRepo) Add(ctx context.Context, project *repository.Project) (int64, error) {
	var id int64
	err := r.db.ExecQueryRow(ctx,
		`INSERT INTO projects(name) VALUES ($1) RETURNING id`, project.Name).Scan(&id)
	return id, err
}

func (r *ProjectsRepo) GetById(ctx context.Context, id int64) (*repository.Project, error) {
	var project repository.Project
	err := r.db.Get(ctx,
		&project, "SELECT id,name,created_at,updated_at FROM projects WHERE id=$1", id)

	if err == sql.ErrNoRows || err != nil && err.Error() == "scanning one: no rows in result set" {
		return nil, repository.ErrObjectNotFound
	}
	return &project, err
}

func (r *ProjectsRepo) List(ctx context.Context) ([]*repository.Project, error) {
	projects := make([]*repository.Project, 0)
	err := r.db.Select(ctx, &projects, "SELECT id,name,created_at,updated_at FROM projects")
	return projects, err
}

func (r *ProjectsRepo) Update(ctx context.Context, project *repository.Project) (bool, error) {
	result, err := r.db.Exec(ctx,
		"UPDATE projects SET name = $1 WHERE id = $2", project.Name, project.ID)
	if result.RowsAffected() == 0 {
		return false, repository.ErrObjectNotFound
	}
	return true, err
}

func (r *ProjectsRepo) DeleteById(ctx context.Context, id int64) (bool, error) {
	result, err := r.db.Exec(ctx, "DELETE FROM projects WHERE id=$1", id)
	if result.RowsAffected() == 0 {
		return false, repository.ErrObjectNotFound
	}
	return true, err
}
