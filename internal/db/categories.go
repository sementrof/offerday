package db

import (
	"context"
	"time"

	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type Categories struct {
	Id        uuid.UUID
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type categoriesQuery struct {
	runner *pgxpool.Pool
	logger *zap.Logger
}

type CategoriesQuery interface {
	NewInsertCategories(ctx context.Context, u *Categories) error
}

func NewCategoriesQuery(runner *pgxpool.Pool, logger *zap.Logger) CategoriesQuery {
	return &categoriesQuery{
		runner: runner,
		logger: logger,
	}
}

func (q categoriesQuery) NewInsertCategories(ctx context.Context, c *Categories) error {
	query := `INSERT INTO categories(name, createdAt, updatedAt)
		  VALUES ($1, $2, $3)`
	_, err := q.runner.Exec(ctx, query, c.Name, c.CreatedAt, c.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}
