package account

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type writer struct {
	db *pgxpool.Pool
}

// Delete implements Writer.
func (w *writer) Delete(ctx context.Context, data Account) error {
	query := `
		DELETE FROM accounts
		WHERE id = $1;
	`
	_, err := w.db.Exec(ctx, query, data.ID)
	return err
}

// Save implements Writer.
func (w *writer) Save(ctx context.Context, data Account) (res *Account, err error) {
	query := `
		INSERT INTO accounts (id, email, created_at)
			VALUES ($1, $2, NOW())
		ON CONFLICT (id)
		DO UPDATE
			SET id = $1, email = $2, updated_at = NOW();
	`
	err = w.db.QueryRow(ctx, query, data.ID, data.Email).Scan(&res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func NewWriter(db *pgxpool.Pool) Writer {
	return &writer{db}
}
