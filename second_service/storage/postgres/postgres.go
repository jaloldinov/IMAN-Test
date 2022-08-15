package postgres

import (
	"github.com/jaloldinov/IMAN-Updated/post_service/storage"

	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jaloldinov/IMAN-Updated/post_service/config"
)

type Store struct {
	db   *pgxpool.Pool
	post storage.PostI
}

func NewPostgres(psqlConnString string, cfg config.Config) (storage.StorageI, error) {
	// First set up a connection pool
	config, err := pgxpool.ParseConfig(psqlConnString)
	if err != nil {
		return nil, err
	}

	config.AfterConnect = nil
	config.MaxConns = int32(cfg.PostgresMaxConnections)

	pool, err := pgxpool.ConnectConfig(context.Background(), config)

	return &Store{
		db: pool,
	}, err
}

func (s *Store) Post() storage.PostI {
	if s.post == nil {
		s.post = NewPostRepo(s.db)
	}
	return s.post
}
