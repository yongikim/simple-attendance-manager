package repository

import (
	"simple-attendance-manager/attendance/entity"
	"simple-attendance-manager/attendance/utility"
)

type UserRepository interface {
	Create(UserCreateRequest) (*entity.User, error)
	FindByID(entity.UserID) (*entity.User, error)
	FindByName(string) (*entity.User, error)
	FindByGrade(entity.Grade) (*entity.User, error)
	UpdateGrade(id entity.UserID, grade entity.Grade) error
	UpdateName(id entity.UserID, name string) error
	Delete(id entity.UserID) error
	FindAllWithAttendancesByDate(utility.SimpleDate) []UserWithAttendances
}

type UserCreateRequest struct {
	Name  string
	Grade entity.Grade
}

type UserWithAttendances struct {
	User        entity.User
	Attendances []entity.Attendance
}
