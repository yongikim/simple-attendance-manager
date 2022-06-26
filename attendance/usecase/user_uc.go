package usecase

import (
	"simple-attendance-manager/attendance/entity"
	"simple-attendance-manager/attendance/repository"
	"simple-attendance-manager/attendance/utility"
)

// Input Boundary
type UserUsecase interface {
	Create(UserCreateInputData) (*entity.User, error)
	GetByID(UserGetByIDInputData) (*entity.User, error)
	GetByName(UserGetByNameInputData) (*entity.User, error)
	GetByGrade(UserGetByGradeInputData) (*entity.User, error)
	UpdateName(UserUpdateNameInputData) error
	UpdateGrade(UserUpdateGradeInputData) error
	Delete(UserDeleteInputData) error
	GetAllUsersWithAttendanceByDate(utility.SimpleDate) []repository.UserWithAttendances
}

// Input Data
type UserCreateInputData struct {
	Name  string
	Grade entity.Grade
}
type UserGetByIDInputData = entity.UserID
type UserGetByNameInputData = string
type UserGetByGradeInputData = entity.Grade
type UserUpdateNameInputData struct {
	ID   entity.UserID
	Name string
}
type UserUpdateGradeInputData struct {
	ID    entity.UserID
	Grade entity.Grade
}
type UserDeleteInputData = entity.UserID

// Output Data
type UserOutput struct {
	ID    entity.UserID
	Name  string
	Grade entity.Grade
}

// Interactor
type UserInteractor struct {
	UserRepo repository.UserRepository
}

func NewUserInteractor(u_repo repository.UserRepository) UserInteractor {
	return UserInteractor{
		UserRepo: u_repo,
	}
}

func (interactor UserInteractor) Create(input UserCreateInputData) (*entity.User, error) {
	request := repository.UserCreateRequest{
		Name:  input.Name,
		Grade: input.Grade,
	}
	user, err := interactor.UserRepo.Create(request)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (interactor UserInteractor) GetByID(input UserGetByIDInputData) (*entity.User, error) {
	user, err := interactor.UserRepo.FindByID(input)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (interactor UserInteractor) GetByName(input UserGetByNameInputData) (*entity.User, error) {
	user, err := interactor.UserRepo.FindByName(input)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (interactor UserInteractor) GetByGrade(input UserGetByGradeInputData) (*entity.User, error) {
	user, err := interactor.UserRepo.FindByGrade(input)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (interactor UserInteractor) UpdateName(input UserUpdateNameInputData) error {
	if err := interactor.UserRepo.UpdateName(input.ID, input.Name); err != nil {
		return err
	}

	return nil
}

func (interactor UserInteractor) UpdateGrade(input UserUpdateGradeInputData) error {
	if err := interactor.UserRepo.UpdateGrade(input.ID, input.Grade); err != nil {
		return err
	}

	return nil
}

func (interactor UserInteractor) Delete(input UserDeleteInputData) error {
	if err := interactor.UserRepo.Delete(input); err != nil {
		return err
	}

	return nil
}

func (interactor UserInteractor) GetAllUsersWithAttendanceByDate(
	date utility.SimpleDate,
) []repository.UserWithAttendances {
	result := interactor.UserRepo.FindAllWithAttendancesByDate(date)
	return result
}
