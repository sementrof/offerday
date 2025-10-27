package db

import (
	"context"
	"time"

	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type Users struct {
	Id        uuid.UUID
	Name      string
	Password  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type usersQuery struct {
	runner *pgxpool.Pool
	logger *zap.Logger
}

type UsersQuery interface {
	NewInsertUser(ctx context.Context, u *Users) error
	// DelectetOrderNew(ctx context.Context, id int64) error
}

func NewUsersQuery(runner *pgxpool.Pool, logger *zap.Logger) UsersQuery {
	return &usersQuery{
		runner: runner,
		logger: logger,
	}
}

func (q usersQuery) NewInsertUser(ctx context.Context, u *Users) error {
	query := `INSERT INTO users(name, email, password, createdAt, updatedAt)
		  VALUES ($1, $2, $3, $4, $5)`
	_, err := q.runner.Exec(ctx, query, u.Name, u.Email, u.Password, u.CreatedAt, u.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (q usersQuery) DeleteUsers(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM users WHERE id = '$1'`
	_, err := q.runner.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}
