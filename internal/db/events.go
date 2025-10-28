package db

import (
	"context"
	"time"

	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type Events struct {
	Id          uuid.UUID
	Title       string
	Description string
	Date        time.Time
	OrganizerId uuid.UUID
	CategoryId  uuid.UUID
	LocationId  uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type eventsQuery struct {
	runner *pgxpool.Pool
	logger *zap.Logger
}

type EventsQuery interface {
	NewInsertEvents(ctx context.Context, u *Events) error
}

func NewEventsQuery(runner *pgxpool.Pool, logger *zap.Logger) EventsQuery {
	return &eventsQuery{
		runner: runner,
		logger: logger,
	}
}

func (q eventsQuery) NewInsertEvents(ctx context.Context, c *Events) error {
	query := `INSERT INTO events(title, description, date, organizerId, CategoryId, LocationId, createdAt, updatedAt)
		  VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := q.runner.Exec(ctx, query, c.Title, c.Description, c.Date, c.OrganizerId, c.CategoryId, c.LocationId, c.CreatedAt, c.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}
