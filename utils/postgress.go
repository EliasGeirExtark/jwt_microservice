package utils

import (
	"github.com/extark/jwt_microservice/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDB() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(Cfg.DBDSN), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	if err = db.AutoMigrate(
		&models.Account{},
	); err != nil {
		return nil, err
	}

	return db, nil
}
