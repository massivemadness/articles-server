package storage

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/massivemadness/articles-server/internal/config"
)

type Storage struct {
	*pgxpool.Pool
}

func New(cfg *config.Config) (*Storage, error) {
	url := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Name,
	)

	pool, err := pgxpool.New(context.Background(), url)
	if err != nil {
		return nil, err
	}

	err = pool.Ping(context.Background())
	if err != nil {
		return nil, err
	}

	return &Storage{pool}, nil
}
