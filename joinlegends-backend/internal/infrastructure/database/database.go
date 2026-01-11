package database

import (
	"go-backend-stream/internal/utilities/config"

	"github.com/go-goe/goe"
	"github.com/go-goe/postgres"
)

func Connect() (*StreamDB, error) {
	db, err := goe.Open[StreamDB](postgres.Open(config.DATABASE_CONNECTION, postgres.NewConfig(postgres.Config{})))
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (db *StreamDB) Close() error {
	return goe.Close(db)
}

func (db *StreamDB) AutoMigrate() error {
	return goe.Migrate(db).AutoMigrate()
}
