package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var poolConn *pgxpool.Pool

func InitDB() error {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		return fmt.Errorf("Error! DATABASE_URL is empty")
	}

	config, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		return fmt.Errorf("Error! unable to parse database URL: %s", err)
	}

	config.MaxConns = 10
	config.MinConns = 2
	config.MaxConnLifetime = time.Hour
	config.MaxConnIdleTime = 30 * time.Minute
	config.HealthCheckPeriod = time.Minute

	var pool *pgxpool.Pool
	pool, err = pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("Error create connect to database: %w", err)
	}

	err = pool.Ping(context.Background())
	if err != nil {
		log.Fatalf("Database not connect: %w", err)
	}

	poolConn = pool
	log.Println("Database connection established")
	return nil

}

func GetConnection() (*pgxpool.Conn, error) {
	if poolConn == nil {
		return nil, fmt.Errorf("Database pool is not init")
	}
	conn, err := poolConn.Acquire(context.Background())
	if err != nil {
		return nil, fmt.Errorf("Unable to acquire connect %w: ", err)
	}

	return conn, nil
}

func CloseConnection(conn *pgxpool.Conn) {
	if conn != nil {
		conn.Release()
	}
}
