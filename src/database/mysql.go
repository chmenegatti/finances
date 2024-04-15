package database

import (
	"fmt"
	"sync"

	"finances/src/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB   *gorm.DB
	once sync.Once
)

func Connect() {
	once.Do(
		func() {
			dsn := fmt.Sprintf(
				"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Cfg.DbUser, config.Cfg.DbPass,
				config.Cfg.DbHost, config.Cfg.DbPort, config.Cfg.DbName,
			)
			db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
			if err != nil {
				panic(err)
			}
			DB = db
		},
	)
}
