package db

import (
	"context"
	"time"

	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type Locations struct {
	Id        uuid.UUID
	Name      string
	Addres    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type locationsQuery struct {
	runner *pgxpool.Pool
	logger *zap.Logger
}

type LocationsQuery interface {
	NewInsertLocations(ctx context.Context, u *Locations) error
}

func NewLocationsQuery(runner *pgxpool.Pool, logger *zap.Logger) LocationsQuery {
	return &locationsQuery{
		runner: runner,
		logger: logger,
	}
}

func (q locationsQuery) NewInsertLocations(ctx context.Context, c *Locations) error {
	query := `INSERT INTO locations(name, addres, createdAt, updatedAt)
		  VALUES ($1, $2, $3, $4)`
	_, err := q.runner.Exec(ctx, query, c.Name, c.Addres, c.CreatedAt, c.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}
