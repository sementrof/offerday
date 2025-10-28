package deps

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sementrof/offerday/internal/config"
	"github.com/sementrof/offerday/internal/db"
	"github.com/sementrof/offerday/internal/logger"
	"go.uber.org/zap"
)

type DB struct {
	Users      db.UsersQuery
	Categories db.CategoriesQuery
	Locations  db.LocationsQuery
	Events     db.EventsQuery
}

type Dependencies struct {
	DB     DB
	Pool   *pgxpool.Pool
	Logger *zap.Logger
}

func ProvideDependencies(ctx context.Context, cfg config.AppConfig) (*Dependencies, error) {
	logger := logger.NewLogger()

	pool, err := db.Connection(cfg, logger)
	if err != nil {
		logger.Fatal("Failed to init db", zap.Error(err))
		return nil, err
	}

	deps := &Dependencies{
		DB: DB{
			Users:      db.NewUsersQuery(pool, logger),
			Categories: db.NewCategoriesQuery(pool, logger),
			Locations:  db.NewLocationsQuery(pool, logger),
			Events:     db.NewEventsQuery(pool, logger),
		},
		Pool:   pool,
		Logger: logger,
	}

	if err := pool.Ping(ctx); err != nil {
		logger.Fatal("Failed to ping database", zap.Error(err))
		pool.Close()
		return nil, err
	}

	logger.Info("Dependencies initialized successfully")
	return deps, nil
}

func (d *Dependencies) Cleanup() {
	d.Logger.Info("Cleaning up dependencies")
	d.Logger.Sync()
	d.Pool.Close()
}
