package in_memory

import (
	"simple-attendance-manager/attendance/entity"
	"simple-attendance-manager/attendance/repository"
	"simple-attendance-manager/attendance/usecase"
	"simple-attendance-manager/attendance/utility"
	"time"
)

type AttendanceRecord struct {
	ID     uint
	Type   string
	At     time.Time
	UserID uint
}

type AttendanceAPI struct {
	DataBase *DataBase
}

func NewAttendanceAPI(db *DataBase) AttendanceAPI {
	return AttendanceAPI{DataBase: db}
}

func (db *AttendanceAPI) CreateAttendance(input entity.Attendance) error {
	id := len(db.DataBase.Attendances)
	attendance := AttendanceRecord{
		ID:     uint(id),
		Type:   string(input.Type),
		At:     input.At,
		UserID: uint(input.UserID),
	}
	db.DataBase.Attendances = append(db.DataBase.Attendances, attendance)
	return nil
}

func (db *AttendanceAPI) FindByDateWithUser(
	date utility.SimpleDate,
) []usecase.UserAttendanceOutput {
	user_attendances := &[]repository.AttendanceWithUser{}
	for _, record := range db.DataBase.Attendances {
		if record.At.After(date.Time()) {
			attendance := entity.Attendance{
				ID:     entity.AttendanceID(record.ID),
				UserID: entity.UserID(record.UserID),
				Type:   entity.AttendanceType(record.Type),
				At:     record.At,
			}
			var user *entity.User
			for _, user_record := range db.DataBase.Users {
				if user_record.ID == record.UserID {
					user = &entity.User{
						ID:    entity.UserID(user_record.ID),
						Name:  user_record.Name,
						Grade: entity.Grade(user_record.Grade),
					}
				}
			}
			user_attendance := usecase.UserAttendanceOutput{
				Attendance: attendance,
				User:       *user,
			}
			*user_attendances = append(*user_attendances, user_attendance)
		}
	}
	return *user_attendances
}

func (db *AttendanceAPI) FindByDateRangeWithUser(
	from utility.SimpleDate,
	to utility.SimpleDate,
) []usecase.UserAttendanceOutput {
	user_attendances := &[]repository.AttendanceWithUser{}
	for _, record := range db.DataBase.Attendances {
		from_time := from.Time()
		to_time := to.Time()
		if record.At.After(from_time) && record.At.Before(to_time) {
			attendance := entity.Attendance{
				ID:     entity.AttendanceID(record.ID),
				UserID: entity.UserID(record.UserID),
				Type:   entity.AttendanceType(record.Type),
				At:     record.At,
			}
			var user *entity.User
			for _, user_record := range db.DataBase.Users {
				if entity.UserID(user_record.ID) == entity.UserID(record.UserID) {
					user = &entity.User{
						ID:    entity.UserID(user_record.ID),
						Name:  user_record.Name,
						Grade: entity.Grade(user_record.Grade),
					}
				}
			}
			user_attendance := usecase.UserAttendanceOutput{
				Attendance: attendance,
				User:       *user,
			}
			*user_attendances = append(*user_attendances, user_attendance)
		}
	}
	return *user_attendances
}
