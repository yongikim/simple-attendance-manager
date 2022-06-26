package usecase

import "simple-attendance-manager/attendance/entity"

// Input Boundary
type UserUsecase interface {
	Create(UserCreateInputData) (*entity.User, error)
	GetByID(UserGetByIDInputData) (*entity.User, error)
	GetByName(UserGetByNameInputData) (*entity.User, error)
	GetByGrade(UserGetByGradeInputData) (*entity.User, error)
	UpdateName(UserUpdateNameInputData) error
	UpdateGrade(UserUpdateGradeInputData) error
	Delete(UserDeleteInputData) error
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
	DataAccess DataAccess
	Presenter  Presenter
}

func (interactor UserInteractor) Create(input UserCreateInputData) (*entity.User, error) {
	user, err := interactor.DataAccess.CreateUser(input)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (interactor UserInteractor) GetByID(input UserGetByIDInputData) (*entity.User, error) {
	user, err := interactor.DataAccess.FindUserByID(input)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (interactor UserInteractor) GetByName(input UserGetByNameInputData) (*entity.User, error) {
	user, err := interactor.DataAccess.FindUserByName(input)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (interactor UserInteractor) GetByGrade(input UserGetByGradeInputData) (*entity.User, error) {
	user, err := interactor.DataAccess.FindUserByGrade(input)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (interactor UserInteractor) UpdateName(input UserUpdateNameInputData) error {
	if err := interactor.DataAccess.UpdateUserName(input.ID, input.Name); err != nil {
		return err
	}

	return nil
}

func (interactor UserInteractor) UpdateGrade(input UserUpdateGradeInputData) error {
	if err := interactor.DataAccess.UpdateUserGrade(input.ID, input.Grade); err != nil {
		return err
	}

	return nil
}

func (interactor UserInteractor) Delete(input UserDeleteInputData) error {
	if err := interactor.DataAccess.DeleteUser(input); err != nil {
		return err
	}

	return nil
}
