package database

import (
	"fmt"
	"time"

	"github.com/xkurozaru/pedometer-server/dependency/config"
	"github.com/xkurozaru/pedometer-server/domain/common"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDB(dbConfig config.DBConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Tokyo",
		dbConfig.Host, dbConfig.User, dbConfig.Password, dbConfig.Name, dbConfig.Port,
	)

	db, err := gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{
			NowFunc: func() time.Time {
				return common.DateTimeNow().Time()
			},
		},
	)
	if err != nil {
		return nil, err
	}
	db.Logger = db.Logger.LogMode(logger.Info)

	return db, nil
}
