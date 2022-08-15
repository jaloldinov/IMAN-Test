package postgres

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/jaloldinov/IMAN-Updated/first_service/config"
	"github.com/jaloldinov/IMAN-Updated/first_service/storage"
)

type Store struct {
	db   *pgxpool.Pool
	data storage.DataI
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

func (s *Store) Data() storage.DataI {
	if s.data == nil {
		s.data = NewDataRepo(s.db)
	}
	return s.data
}
