package db

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	model "github.com/bayarindevteam/bayaringo/graph/model"
)

// CreateLink ...
func (d *DatabaseStorage) CreateUser(ctx context.Context, u *model.User) (string, error) {
	var lastID string

	query := sq.Insert("users").
		Columns("name", "username", "email", "password").
		Values(u.Name, u.Username, u.Password, u.Password).
		Suffix("RETURNING \"id\"").
		RunWith(d.db).
		PlaceholderFormat(sq.Dollar)

	err := query.QueryRowContext(ctx).Scan(&lastID)

	return lastID, err
}
