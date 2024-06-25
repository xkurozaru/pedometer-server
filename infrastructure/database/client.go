package database

import (
	"fmt"
	"sync"
	"time"

	"github.com/xkurozaru/pedometer-server/dependency/config"
	"github.com/xkurozaru/pedometer-server/domain/common"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db_instance *gorm.DB
var err_instance error
var once sync.Once

func ConnectDB(dbConfig config.DBConfig) (*gorm.DB, error) {
	once.Do(func() {
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
			err_instance = err
			return
		}
		db.Logger = db.Logger.LogMode(logger.Info)

		db_instance = db
	})

	return db_instance, err_instance
}
