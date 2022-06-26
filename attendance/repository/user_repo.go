package repository

import (
	"simple-attendance-manager/attendance/entity"
	"simple-attendance-manager/attendance/utility"
)

type UserRepository interface {
	CreateUser(UserCreateRequest) (*entity.User, error)
	FindUserByID(entity.UserID) (*entity.User, error)
	FindUserByName(string) (*entity.User, error)
	FindUserByGrade(entity.Grade) (*entity.User, error)
	UpdateUserGrade(id entity.UserID, grade entity.Grade) error
	UpdateUserName(id entity.UserID, name string) error
	DeleteUser(id entity.UserID) error
	FindAllUsersWithAttendanceByDate(utility.SimpleDate) []UserWithAttendances
}

type UserCreateRequest struct {
	Name  string
	Grade entity.Grade
}

type UserWithAttendances struct {
	User        entity.User
	Attendances []entity.Attendance
}
