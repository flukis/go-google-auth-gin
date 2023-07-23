package auth

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type reader struct {
	db *pgxpool.Pool
}

// GetByID implements Reader.
func (r *reader) GetByID(ctx context.Context, id string) (*Account, error) {
	var res Account
	query := `
		SELECT * FROM accounts WHERE id = $1;
	`
	err := r.db.QueryRow(ctx, query, id).Scan(&res.ID, &res.Email, &res.CreatedAt, &res.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func NewReader(db *pgxpool.Pool) Reader {
	return &reader{db}
}
