package infra

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartConnection() (*gorm.DB, error) {
	strConn := "host=localhost user=root password=root dbname=meet_people port=5432 sslmode=disable"
	fmt.Println("try get connection")
	db, err := gorm.Open(postgres.Open(strConn))
	if err != nil {
		return nil, err
	}
	return db, nil
}
