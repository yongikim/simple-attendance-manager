package in_memory

import (
	"simple-attendance-manager/attendance/entity"
	"simple-attendance-manager/attendance/repository"
	"simple-attendance-manager/attendance/utility"
)

type UserRecord struct {
	ID    uint
	Name  string
	Grade int8
}

type UserAPI struct {
	DataBase *DataBase
}

func NewUserAPI(db *DataBase) UserAPI {
	return UserAPI{DataBase: db}
}

func (db *UserAPI) Create(input repository.UserCreateRequest) (*entity.User, error) {
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

func (db *UserAPI) FindByID(id entity.UserID) (*entity.User, error) {
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

func (db *UserAPI) FindByName(name string) (*entity.User, error) {
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

func (db *UserAPI) FindByGrade(grade entity.Grade) (*entity.User, error) {
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

func (db *UserAPI) UpdateName(id entity.UserID, name string) error {
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

func (db *UserAPI) UpdateGrade(id entity.UserID, grade entity.Grade) error {
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

func (db *UserAPI) Delete(id entity.UserID) error {
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

func (db *UserAPI) FindAllWithAttendancesByDate(
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
