package storage

import (
	"database/sql"
	"errors"
	"github.com/massivemadness/articles-server/internal/config"
)

type Storage struct {
	*sql.DB
}

func New(cfg *config.Config) (*Storage, error) {
	// TODO connect to db
	return nil, errors.New("test")
}
