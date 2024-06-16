package database

import (
	"fmt"

	"github.com/xkurozaru/pedometer-server/dependency/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(dbConfig config.DBConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Tokyo",
		dbConfig.Host, dbConfig.User, dbConfig.Password, dbConfig.Name, dbConfig.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
