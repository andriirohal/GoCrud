package db

import (
  "context"
  "fmt"
  "os"
  "github.com/jackc/pgx/v5/pgxpool"
);

var DB *pgxpool.Pool;

func Connect() error {
  ctx := context.Background();

  connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
    os.Getenv("PGUSER"),
    os.Getenv("PGPASSWORD"),
    os.Getenv("PGHOST"),
    os.Getenv("PGPORT"),
    os.Getenv("PGDATABASE"),
  );
  
  pool, err := pgxpool.New(ctx, connString);

  if err != nil {
    return err;
  };

  if err := pool.Ping(ctx); err != nil {
    return err;
  };

  DB = pool;

  return nil;
};