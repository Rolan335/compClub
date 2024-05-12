package db

import (
	"errors"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func DbInit() error {
	var err error
	Db, err = ConnectToDB("postgres", "Pa$$w0rd", "compclub")
	if err != nil {
		return errors.New("can't connect to db")
	}
	Db.AutoMigrate(&Computer{}, &User{}, &Admin{}, &Shift{}, &Rent{})
	return nil
}

func ConnectToDB(user string, password string, dbname string) (*gorm.DB, error) {
	dsn := "host=localhost" + " user=" + user + " dbname=" + dbname + " password=" + password + " sslmode=disable"
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
