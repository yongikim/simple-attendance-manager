package repository

import (
	"simple-attendance-manager/attendance/entity"
	"simple-attendance-manager/attendance/utility"
)

// Data Access Interface (Repository)
type AttendanceRepository interface {
	CreateAttendance(entity.Attendance) error
	FindByDateWithUser(utility.SimpleDate) []AttendanceWithUser
	FindByDateRangeWithUser(from utility.SimpleDate, to utility.SimpleDate) []AttendanceWithUser
}

// Output Data
type AttendanceWithUser struct {
	User       entity.User
	Attendance entity.Attendance
}
