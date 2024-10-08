package postgres

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func New() *pgxpool.Pool {
    dbUrl := os.Getenv("DATABASE_URL")
    config, err := pgxpool.ParseConfig(dbUrl)
    if err != nil {
        log.Fatalf("Unable to parse DATABASE_URL: %v\n", err)
    }

    config.MaxConns = 10
    config.HealthCheckPeriod = time.Minute

    pool, err := pgxpool.NewWithConfig(context.Background(), config)
    if err != nil {
        log.Fatalf("Unable to create connection pool: %v\n", err)
    }

    return pool
}
