package storage

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/massivemadness/articles-server/internal/config"
)

type Storage struct {
	*pgx.Conn
}

func New(cfg *config.Config) (*Storage, error) {
	url := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.DbName,
	)
	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		return nil, err
	}
	defer conn.Close(context.Background())

	return &Storage{conn}, nil
}
