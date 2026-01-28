package model

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBConfig struct {
	DBUser string
	DBPass string
	DBAddr string
	DBPort int
	DBName string
}

func ConfigureMySQL(conf DBConfig) gorm.Dialector {
	return mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", conf.DBUser, conf.DBPass, conf.DBAddr, conf.DBPort, conf.DBName))
}

func InitDB(dialector gorm.Dialector) (*gorm.DB, error) {
	return gorm.Open(dialector, &gorm.Config{})
}
