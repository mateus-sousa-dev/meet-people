package infra

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func StartDBConnection() (*gorm.DB, error) {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	sslMode := os.Getenv("DB_SSL_MODE")
	strConn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port + " sslmode=" + sslMode
	db, err := gorm.Open(postgres.Open(strConn))
	if err != nil {
		return nil, err
	}
	return db, nil
}
