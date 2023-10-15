package repository

import (
	"database/sql"
	"time"
)

type Task struct {
	ID        int64        `db:"id"`
	Name      string       `db:"name"`
	ProjectID int64        `db:"project_id"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}

type Project struct {
	ID        int64        `db:"id"`
	Name      string       `db:"name"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}
