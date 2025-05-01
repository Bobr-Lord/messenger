package repository

import (
	"fmt"
	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"gitlab.com/bobr-lord-messenger/user/internal/config"
	"time"
)

func NewPostgres(cfg *config.Config) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.PostgresUser, cfg.PostgresPass, cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresDB)
	logrus.Info(dsn)
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to postgres: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("unable to connect to postgres: %w", err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(10)

	err = migrateDB(cfg)
	if err != nil {
		fmt.Println("Successfully connected to postgres")

		return nil, err
	}
	return db, nil

}

func migrateDB(cfg *config.Config) error {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.PostgresUser, cfg.PostgresPass, cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresDB)
	m, err := migrate.New("file://migrations", dsn)
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil {
		if err != migrate.ErrNoChange {
			return fmt.Errorf("migrate up: %w", err)
		}
	}
	logrus.Info("migrate ok")
	return nil
}
