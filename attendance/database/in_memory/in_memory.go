package in_memory

import (
	"simple-attendance-manager/attendance/entity"
	"simple-attendance-manager/attendance/repository"
	"simple-attendance-manager/attendance/usecase"
	"simple-attendance-manager/attendance/utility"
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

type NotFoundError struct {
	message string
}

func (e NotFoundError) Error() string {
	return e.message
}

type dataBase struct {
	Users       []UserRecord
	Attendances []AttendanceRecord
}

type InMemoryDB struct {
	DataBase dataBase
}

func NewInMemoryDB() *InMemoryDB {
	return &InMemoryDB{}
}

func (db *InMemoryDB) CreateAttendance(input entity.Attendance) error {
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

func (db *InMemoryDB) CreateUser(input repository.UserCreateRequest) (*entity.User, error) {
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

func (db *InMemoryDB) FindUserByID(id entity.UserID) (*entity.User, error) {
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
		return nil, NotFoundError{message: "User not found"}
	}

	return user, nil
}

func (db *InMemoryDB) FindUserByName(name string) (*entity.User, error) {
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
		return nil, NotFoundError{message: "User not found"}
	}

	return user, nil
}

func (db *InMemoryDB) FindUserByGrade(grade entity.Grade) (*entity.User, error) {
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

func (db *InMemoryDB) UpdateUserName(id entity.UserID, name string) error {
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

func (db *InMemoryDB) UpdateUserGrade(id entity.UserID, grade entity.Grade) error {
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

func (db *InMemoryDB) DeleteUser(id entity.UserID) error {
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

func (db *InMemoryDB) FindByDateWithUser(date utility.SimpleDate) []usecase.UserAttendanceOutput {
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

func (db *InMemoryDB) FindByDateRangeWithUser(
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

func (db *InMemoryDB) FindAllUsersWithAttendanceByDate(
	date utility.SimpleDate,
) []repository.UserWithAttendances {
	today := date.Time()
	result := &[]repository.UserWithAttendances{}

	for _, user_rec := range db.DataBase.Users {
		atds := &[]entity.Attendance{}
		for _, atd := range db.DataBase.Attendances {
			if atd.UserID == user_rec.ID &&
				atd.At.After(today) {
				*atds = append(*atds, entity.Attendance{
					ID:     entity.AttendanceID(atd.ID),
					UserID: entity.UserID(atd.UserID),
					Type:   entity.AttendanceType(atd.Type),
					At:     atd.At,
				})
			}
		}
		user := entity.User{
			ID:    entity.UserID(user_rec.ID),
			Name:  user_rec.Name,
			Grade: entity.Grade(user_rec.Grade),
		}
		*result = append(*result, repository.UserWithAttendances{User: user, Attendances: *atds})
	}

	return *result
}
