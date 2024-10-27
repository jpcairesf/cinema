package config

import (
	"github.com/jpcairesf/cinema/actor/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initialize() (*gorm.DB, error) {
	dsn := "host=localhost user=admin password=admin dbname=cinema port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Errorf("Failed to connect to database: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&entity.Actor{}, &entity.Contact{})
	if err != nil {
		logger.Errorf("Failed to migrate table: %v", err)
		return nil, err
	}

	return db, nil
}
