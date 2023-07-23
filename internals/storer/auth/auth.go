package auth

import (
	"context"
	"time"
)

type Account struct {
	ID        string     `json:"id"`
	Email     string     `json:"email"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type Writer interface {
	Save(ctx context.Context, data Account) (res *Account, err error)
	Delete(ctx context.Context, data Account) error
}

type Reader interface {
	GetByID(ctx context.Context, id string) (res *Account, err error)
}
