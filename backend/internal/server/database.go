package server

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func OpenDatabase(ctx context.Context) (*pgxpool.Pool, error) {
	user, password, database := os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB")
	if user == "" || password == "" || database == "" {
		return nil, errors.New("DBの必須設定がありません")
	}
	config, err := pgxpool.ParseConfig("")
	if err != nil {
		return nil, errors.New("DB設定を解釈できません")
	}
	config.ConnConfig.User = user
	config.ConnConfig.Password = password
	config.ConnConfig.Database = database
	config.ConnConfig.ConnectTimeout = 2 * time.Second
	config.MaxConns = 4
	return pgxpool.NewWithConfig(ctx, config)
}
