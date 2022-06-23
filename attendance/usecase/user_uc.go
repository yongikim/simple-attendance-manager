package usecase

import "simple-attendance-manager/attendance/entity"

// Input Boundary
type UserUsecase interface {
	Create(UserCreateInputData)
	GetByID(UserGetByIDInputData)
	GetByName(UserGetByNameInputData)
	GetByGrade(UserGetByGradeInputData)
	UpdateName(UserUpdateNameInputData)
	UpdateGrade(UserUpdateGradeInputData)
	Delete(UserDeleteInputData)
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

func (interactor UserInteractor) Create(input UserCreateInputData) {
	user, err := interactor.DataAccess.CreateUser(input)
	if err != nil {
		interactor.Presenter.OnError(err)
	}

	interactor.Presenter.OnSuccess(user)
}

func (interactor UserInteractor) GetByID(input UserGetByIDInputData) {
	user, err := interactor.DataAccess.FindUserByID(input)
	if err != nil {
		interactor.Presenter.OnError(err)
	}

	interactor.Presenter.OnSuccess(user)
}

func (interactor UserInteractor) GetByName(input UserGetByNameInputData) {
	user, err := interactor.DataAccess.FindUserByName(input)
	if err != nil {
		interactor.Presenter.OnError(err)
	}

	interactor.Presenter.OnSuccess(user)
}

func (interactor UserInteractor) GetByGrade(input UserGetByGradeInputData) {
	user, err := interactor.DataAccess.FindUserByGrade(input)
	if err != nil {
		interactor.Presenter.OnError(err)
	}

	interactor.Presenter.OnSuccess(user)
}

func (interactor UserInteractor) UpdateName(input UserUpdateNameInputData) {
	if err := interactor.DataAccess.UpdateUserName(input.ID, input.Name); err != nil {
		interactor.Presenter.OnError(err)
	}

	interactor.Presenter.OnSuccess(nil)
}

func (interactor UserInteractor) UpdateGrade(input UserUpdateGradeInputData) {
	if err := interactor.DataAccess.UpdateUserGrade(input.ID, input.Grade); err != nil {
		interactor.Presenter.OnError(err)
	}

	interactor.Presenter.OnSuccess(nil)
}

func (interactor UserInteractor) Delete(input UserDeleteInputData) {
	if err := interactor.DataAccess.DeleteUser(input); err != nil {
		interactor.Presenter.OnError(err)
	}

	interactor.Presenter.OnSuccess(nil)
}
