package infra

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartConnection() (*gorm.DB, error) {
	strConn := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(strConn))
	if err != nil {
		return nil, err
	}
	return db, nil
}
