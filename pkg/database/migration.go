package database

import (
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
	"log/slog"
)

type Migration struct {
	cfg    Config
	db     *gorm.DB
	logger *slog.Logger
}

func NewMigration(db *gorm.DB, cfg Config, log *slog.Logger) *Migration {
	return &Migration{
		db:     db,
		cfg:    cfg,
		logger: log,
	}
}

func (m *Migration) CreateSchema() {
	m.db.Exec(fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s", m.cfg.DatabaseSchema))
}

func (m *Migration) Migrate() {
	var err error

	migration, err := m.getMigrationInstance("migrations/")
	if err != nil {
		m.logger.Error("error creating migration instance", slog.Any("error", err))
	}

	err = migration.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		m.logger.Error("error executing migration", slog.Any("error", err))
	}
}

func (m *Migration) getMigrationInstance(dir string) (*migrate.Migrate, error) {
	driver, err := m.getDriver()

	if err != nil {
		return nil, err
	}

	return migrate.NewWithDatabaseInstance(fmt.Sprintf("file://%s", dir), m.cfg.DatabaseDBName, driver)
}

func (m *Migration) getDriver() (database.Driver, error) {
	db, err := m.db.DB()

	if err != nil {
		return nil, err
	}

	return postgres.WithInstance(db, &postgres.Config{
		SchemaName: m.cfg.DatabaseSchema,
	})
}
