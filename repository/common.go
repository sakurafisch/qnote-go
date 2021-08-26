package repository

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/sakurafisch/qnote-go/entity"
)

var MainDB *gorm.DB

func ConnDB() (db *gorm.DB, err error) {
	dsn := "qnote:pa$$w0rd@tcp(127.0.0.1:3306)/qnote?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return
}

func InitDB() {
	var err error
	MainDB, err = ConnDB()
	if err != nil {
		panic("failed to connect database")
	}

	MainDB.AutoMigrate(&entity.User{})
}
