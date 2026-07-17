package db;

import (
"context"
"github.com/jackc/pgx/v5/pgxpool"
);

var DB *pgxpool.Pool;

func Connect() error {
  ctx := context.Background();

  pool, err := pgxpool.New(ctx, "postgres://andrii:password@localhost:5432/students");

  err = pool.Ping(ctx);

  if err != nil {
    return err;
  };

  DB = pool;

  return nil;
};