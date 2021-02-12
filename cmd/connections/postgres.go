package connections

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

func GetPostgresDatabase(connectionString string) (pool *pgxpool.Pool, err error) {
	conf, err := pgxpool.ParseConfig(connectionString)
	if err != nil {
		return
	}

	pool, err = pgxpool.ConnectConfig(context.TODO(), conf)
	return
}
