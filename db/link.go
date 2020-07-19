package db

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	// "github.com/lib/pq"
	model "github.com/bayarindevteam/bayaringo/graph/model"
)

// CreateLink ...
func (d *DatabaseStorage) CreateLink(ctx context.Context, l *model.Link) (string, error) {
	var lastID string

	query := sq.Insert("links").
		Columns("title", "address", "userid").
		Values(l.Title, l.Address, "999242e0-4855-44b6-86fb-41fd3b8bc526").
		Suffix("RETURNING \"id\"").
		RunWith(d.db).
		PlaceholderFormat(sq.Dollar)

	err := query.QueryRowContext(ctx).Scan(&lastID)

	return lastID, err
}
