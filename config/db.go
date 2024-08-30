package config

import (
	"api-backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=francesco@dedpartners dbname=colloqui port=5432 sslmode=disable TimeZone=Europe/Rome"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	DB = db
	return db, err
}

func GetBlankDB(db *gorm.DB) error {
	db.Migrator().DropTable(&models.Users{})
	db.Migrator().DropTable(&models.ProjectModel{})
	db.Migrator().DropTable(&models.TaskModel{})
	return nil
}
