package usecase

import (
	"simple-attendance-manager/attendance/entity"
	"time"
)

// Input Data
type UserArriveInputData struct {
	UserID entity.UserID
	At     time.Time
}
type UserLeaveInputData struct {
	UserID entity.UserID
	At     time.Time
}
type SimpleDate struct {
	Year  int8
	Month int8
	Day   int8
}

// Input Boundary
type AttendanceUsecase interface {
	/*
		出席する

		入力: 時刻とユーザID
		出力: 出席記録

		1. 時刻とユーザIDを検証する
		2. 出席を記録する
	*/
	UserArrive(UserArriveInputData) (*UserArriveOutput, error)

	// 離席する
	UserLeave(UserLeaveInputData) (*UserLeaveOutput, error)

	/*
		与えられた日付の出社簿を取得する

		入力: 日付
		出力: 当該日の出社記録(出社した人・出社時間・退社時間)の一覧(出社簿)

		1. 日付を受け取り、検証する
		2. 当該日付の出社記録一覧を取得する
	*/

	GetByDate(SimpleDate) GetByDateOutput

	/*
		与えられた期間の出社簿を取得する

		入力: 開始日付、終了日付
		出力: 当該期間の出社記録(出社した人・出社時間・退社時間)の一覧(出社簿)

		1. 期間を受け取り、検証する
		2. 当該期間の出社記録一覧を取得する
	*/
	GetByDateRange(from SimpleDate, to SimpleDate) GetByDateRangeOutput
}

// Output Boundary
type Presenter interface {
	OnSuccess(interface{})
	OnError(error)
}

// Output Data
type AttendanceOutput struct {
	Type entity.AttendanceType
	At   time.Time
}
type UserArriveOutput struct {
	User       UserOutput
	Attendance AttendanceOutput
}
type UserLeaveOutput struct {
	User       UserOutput
	Attendance AttendanceOutput
}

type GetByDateOutput []entity.Attendance

type GetByDateRangeOutput []entity.Attendance

// Data Access Interface (Repository)
type DataAccess interface {
	CreateAttendance(entity.Attendance) error
	CreateUser(UserCreateInputData) (*entity.User, error)
	FindUserByID(entity.UserID) (*entity.User, error)
	FindUserByName(string) (*entity.User, error)
	FindUserByGrade(entity.Grade) (*entity.User, error)
	UpdateUserName(id entity.UserID, name string) error
	UpdateUserGrade(id entity.UserID, grade entity.Grade) error
	DeleteUser(id entity.UserID) error
	FindByDate(SimpleDate) []entity.Attendance
	FindByDateRange(from SimpleDate, to SimpleDate) []entity.Attendance
}

// Interactor
type AttendanceInteractor struct {
	DataAccess DataAccess
}

func (i AttendanceInteractor) UserArrive(params UserArriveInputData) (*UserArriveOutput, error) {
	attendance := entity.Attendance{
		UserID: params.UserID,
		Type:   entity.Arrive,
		At:     params.At,
	}
	if err := i.DataAccess.CreateAttendance(attendance); err != nil {
		return nil, err
	}

	user, err := i.DataAccess.FindUserByID(params.UserID)
	if err != nil {
		return nil, err
	}

	output := UserArriveOutput{
		User: UserOutput{
			ID: user.ID, Name: user.Name, Grade: user.Grade,
		},
		Attendance: AttendanceOutput{
			At:   params.At,
			Type: entity.Arrive,
		},
	}

	return &output, err
}

func (i AttendanceInteractor) UserLeave(params UserLeaveInputData) (*UserLeaveOutput, error) {
	attendance := entity.Attendance{
		UserID: params.UserID,
		Type:   entity.Leave,
		At:     params.At,
	}
	if err := i.DataAccess.CreateAttendance(attendance); err != nil {
		return nil, err
	}

	user, err := i.DataAccess.FindUserByID(params.UserID)
	if err != nil {
		return nil, err
	}

	output := UserLeaveOutput{
		User: UserOutput{
			ID: user.ID, Name: user.Name, Grade: user.Grade,
		},
		Attendance: AttendanceOutput{
			At:   params.At,
			Type: entity.Leave,
		},
	}

	return &output, err
}

func (i AttendanceInteractor) GetByDate(date SimpleDate) GetByDateOutput {
	attendances := i.DataAccess.FindByDate(date)
	return attendances
}

func (i AttendanceInteractor) GetByDateRange(from SimpleDate, to SimpleDate) GetByDateRangeOutput {
	attendances := i.DataAccess.FindByDateRange(from, to)
	return attendances
}
