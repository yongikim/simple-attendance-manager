package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

type User struct {
	gorm.Model
	Name        string
	Grade       int8
	Attendances []Attendance
	Leavings    []Leaving
}

type Attendance struct {
	gorm.Model
	UserID uint
}

type Leaving struct {
	gorm.Model
	UserID uint
}

func Connect() {
	dsn := "dev:password@tcp(127.0.0.1:3306)/simple_attendance_manager_dev?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Database connection not established.")
	}
	db.AutoMigrate(&User{}, &Attendance{}, &Leaving{})
}

func GetDB() *gorm.DB {
	return db
}
