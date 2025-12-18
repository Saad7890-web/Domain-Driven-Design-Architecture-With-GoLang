package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/Saad7890-web/internal/domain/user"
)


type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository{
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, u *user.User) error {
	query := `
		INSERT INTO users (id, email, password, name, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	now := time.Now()

	_, err := r.db.ExecContext(
		ctx,
		query,
		u.ID,
		u.Email,
		u.Password,
		u.Name,
		now,
		now,
	)

	return err
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*user.User, error){
	query := `
		SELECT id, email, password, name, created_at, updated_at
		FROM users
		WHERE email = $1
	`

	var u user.User

	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&u.ID,
		&u.Email,
		&u.Password,
		&u.Name,
		&u.CreatedAt,
		&u.UpdatedAt,
	)


	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &u, nil
}