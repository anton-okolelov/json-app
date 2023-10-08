package database

import (
	"context"
	"fmt"

	"github.com/anton-okolelov/json-app/internal/config"
	"github.com/jackc/pgx/v4/pgxpool"
)

// InitDB - initialize database
func InitDB(dbCfg config.DBConf) (*pgxpool.Pool, error) {
	cfg, _ := pgxpool.ParseConfig("")
	cfg.ConnConfig.Host = dbCfg.Host
	cfg.ConnConfig.Port = dbCfg.Port
	cfg.ConnConfig.User = dbCfg.User
	cfg.ConnConfig.Password = dbCfg.Password
	cfg.ConnConfig.Database = dbCfg.DBName
	cfg.ConnConfig.PreferSimpleProtocol = true
	cfg.MaxConns = 20

	dbPool, err := pgxpool.ConnectConfig(context.Background(), cfg)
	if err != nil {
		return nil, fmt.Errorf("cannot connect to db: %w", err)
	}

	return dbPool, nil
}
