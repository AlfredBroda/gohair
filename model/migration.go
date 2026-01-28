package model

import "gorm.io/gorm"

func CreateDB(dialector gorm.Dialector) error {
	db, err := InitDB(dialector)
	if err != nil {
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	_, err = sqlDB.Exec("CREATE DATABASE rhair;")
	if err != nil {
		return err
	}

	return nil
}

func Migrate(dialector gorm.Dialector) error {
	db, err := InitDB(dialector)
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	return db.AutoMigrate(&Article{})
}
