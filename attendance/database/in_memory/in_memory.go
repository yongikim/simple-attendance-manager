package in_memory

import (
	"simple-attendance-manager/attendance/entity"
	"simple-attendance-manager/attendance/usecase"
	"time"
)

type UserRecord struct {
	ID    uint
	Name  string
	Grade int8
}

type AttendanceRecord struct {
	ID     uint
	Type   string
	At     time.Time
	UserID uint
}

type DataBase struct {
	Users       []UserRecord
	Attendances []AttendanceRecord
}

type InMemmoryDB struct {
	DataBase DataBase
}

func NewInMemoryDB() *DataBase {
	return &DataBase{}
}

func (db *InMemmoryDB) CreateAttendance(input entity.Attendance) error {
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

func (db *InMemmoryDB) CreateUser(input usecase.UserCreateInputData) (*entity.User, error) {
	id := len(db.DataBase.Users)
	user := UserRecord{
		ID:    uint(id),
		Name:  input.Name,
		Grade: int8(input.Grade),
	}
	db.DataBase.Users = append(db.DataBase.Users, user)

	result := entity.User{
		ID:    entity.UserID(user.ID),
		Name:  user.Name,
		Grade: entity.Grade(user.Grade),
	}
	return &result, nil
}

type NotFoundError struct{}

func (e NotFoundError) Error() string {
	return "Not Found"
}

func (db *InMemmoryDB) FindUserByID(id entity.UserID) (*entity.User, error) {
	var user *entity.User
	for _, user_record := range db.DataBase.Users {
		if entity.UserID(user_record.ID) == id {
			user = &entity.User{
				ID:    entity.UserID(user_record.ID),
				Name:  user_record.Name,
				Grade: entity.Grade(user_record.Grade),
			}
		}
	}
	if user == nil {
		return nil, NotFoundError{}
	}

	return user, nil
}

func (db *InMemmoryDB) FindUserByName(name string) (*entity.User, error) {
	var user *entity.User
	for _, user_record := range db.DataBase.Users {
		if user_record.Name == name {
			user = &entity.User{
				ID:    entity.UserID(user_record.ID),
				Name:  user_record.Name,
				Grade: entity.Grade(user_record.Grade),
			}
		}
	}
	if user == nil {
		return nil, NotFoundError{}
	}

	return user, nil
}

func (db *InMemmoryDB) FindUserByGrade(grade entity.Grade) (*entity.User, error) {
	var user *entity.User
	for _, user_record := range db.DataBase.Users {
		if entity.Grade(user_record.Grade) == grade {
			user = &entity.User{
				ID:    entity.UserID(user_record.ID),
				Name:  user_record.Name,
				Grade: entity.Grade(user_record.Grade),
			}
		}
	}
	if user == nil {
		return nil, NotFoundError{}
	}

	return user, nil
}

func (db *InMemmoryDB) UpdateUserName(id entity.UserID, name string) error {
	var found bool
	for i := 0; i < len(db.DataBase.Users); i++ {
		if db.DataBase.Users[i].ID == uint(id) {
			db.DataBase.Users[i].Name = name
			found = true
		}
	}

	if !found {
		return NotFoundError{}
	}

	return nil
}

func (db *InMemmoryDB) UpdateUserGrade(id entity.UserID, grade entity.Grade) error {
	var found bool
	for i := 0; i < len(db.DataBase.Users); i++ {
		if db.DataBase.Users[i].ID == uint(id) {
			db.DataBase.Users[i].Grade = int8(grade)
			found = true
		}
	}

	if !found {
		return NotFoundError{}
	}

	return nil
}

func (db *InMemmoryDB) DeleteUser(id entity.UserID) error {
	var found bool
	for i := 0; i < len(db.DataBase.Users); i++ {
		if db.DataBase.Users[i].ID == uint(id) {
			found = true
			db.DataBase.Users = append(db.DataBase.Users[:i], db.DataBase.Users[i+1:]...)
		}
	}

	if !found {
		return NotFoundError{}
	}

	return nil
}

func (db *InMemmoryDB) FindByDate(date usecase.SimpleDate) []entity.Attendance {
	var attendances []entity.Attendance
	for _, record := range db.DataBase.Attendances {
		if record.At.Year() == int(date.Year) &&
			record.At.Month() == time.Month(date.Month) &&
			record.At.Day() == int(date.Day) {
			attendance := entity.Attendance{
				ID:     entity.AttendanceID(record.ID),
				UserID: entity.UserID(record.UserID),
				Type:   entity.AttendanceType(record.Type),
				At:     record.At,
			}
			attendances = append(attendances, attendance)
		}
	}
	return attendances
}

func (db *InMemmoryDB) FindByDateRange(from usecase.SimpleDate, to usecase.SimpleDate) []entity.Attendance {
	var attendances []entity.Attendance
	for _, record := range db.DataBase.Attendances {
		from_time := time.Date(int(from.Year), time.Month(from.Month), int(from.Day), 0, 0, 0, 0, time.Local)
		to_time := time.Date(int(to.Year), time.Month(to.Month), int(to.Day), 0, 0, 0, 0, time.Local)
		if record.At.After(from_time) && record.At.Before(to_time) {
			attendance := entity.Attendance{
				ID:     entity.AttendanceID(record.ID),
				UserID: entity.UserID(record.UserID),
				Type:   entity.AttendanceType(record.Type),
				At:     record.At,
			}
			attendances = append(attendances, attendance)
		}
	}
	return attendances
}
