package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

type DatabaseMySQLOptions struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func NewDatabase() *Database {
	return &Database{}
}

func (d *Database) connect(dialector gorm.Dialector) error {
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return err
	}
	d.DB = db
	return nil
}

func (d *Database) ConnectMySQL(options DatabaseMySQLOptions) error {
	dsn := options.User + ":" + options.Password + "@tcp(" + options.Host + ":" + options.Port + ")/" + options.Database + "?charset=utf8mb4&parseTime=True&loc=Local"
	return d.connect(mysql.Open(dsn))
}
