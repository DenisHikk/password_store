package db

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ConfigDB struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string

	MaxConns          int32
	MinConns          int32
	MaxConnLifetime   time.Duration
	MaxConnIdleTime   time.Duration
	HealthCheckPeriod time.Duration
	ConnectionTimeout time.Duration
}

func (c ConfigDB) DSN() string {
	ssl := c.SSLMode
	if ssl == "" {
		ssl = "disable"
	}
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		c.User, c.Password, c.Host, c.Port, c.DBName, c.SSLMode,
	)
}

func NewPool(ctx context.Context, config ConfigDB) (*pgxpool.Pool, error) {
	cfg, err := pgxpool.ParseConfig(config.DSN())
	if err != nil {
		return nil, err
	}

	if config.MaxConns > 0 {
		cfg.MaxConns = config.MaxConns
	} else {
		cfg.MaxConns = 10
	}
	cfg.MinConns = config.MinConns
	if config.MaxConnLifetime == 0 {
		cfg.MaxConnLifetime = time.Hour
	} else {
		cfg.MaxConnLifetime = config.MaxConnLifetime
	}
	if config.MaxConnIdleTime == 0 {
		cfg.MaxConnIdleTime = 30 * time.Minute
	} else {
		cfg.MaxConnIdleTime = config.MaxConnIdleTime
	}
	if config.HealthCheckPeriod == 0 {
		cfg.HealthCheckPeriod = time.Minute
	} else {
		cfg.HealthCheckPeriod = config.HealthCheckPeriod
	}

	if config.ConnectionTimeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, config.ConnectionTimeout)
		defer cancel()
	}

	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, err
	}
	return pool, nil
}
