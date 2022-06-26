package usecase

import (
	"simple-attendance-manager/attendance/entity"
	"simple-attendance-manager/attendance/repository"
	"simple-attendance-manager/attendance/utility"
	"time"
)

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

	GetByDate(utility.SimpleDate) GetByDateOutput

	/*
		与えられた期間の出社簿を取得する

		入力: 開始日付、終了日付
		出力: 当該期間の出社記録(出社した人・出社時間・退社時間)の一覧(出社簿)

		1. 期間を受け取り、検証する
		2. 当該期間の出社記録一覧を取得する
	*/
	GetByDateRange(from utility.SimpleDate, to utility.SimpleDate) GetByDateRangeOutput
}

// Input Data
type UserArriveInputData struct {
	UserID entity.UserID
	At     time.Time
}
type UserLeaveInputData struct {
	UserID entity.UserID
	At     time.Time
}

// Output Data
type UserAttendanceOutput = repository.AttendanceWithUser
type UserArriveOutput = UserAttendanceOutput
type UserLeaveOutput = UserAttendanceOutput
type GetByDateOutput = []UserAttendanceOutput
type GetByDateRangeOutput = []UserAttendanceOutput

// Interactor
type AttendanceInteractor struct {
	AttendanceRepo repository.AttendanceRepository
	UserRepo       repository.UserRepository
}

func NewAttendanceInteractor(
	a_repo repository.AttendanceRepository,
	u_repo repository.UserRepository,
) AttendanceInteractor {
	return AttendanceInteractor{
		AttendanceRepo: a_repo,
		UserRepo:       u_repo,
	}
}

func (i AttendanceInteractor) UserArrive(params UserArriveInputData) (*UserArriveOutput, error) {
	attendance := entity.Attendance{
		UserID: params.UserID,
		Type:   entity.Arrive,
		At:     params.At,
	}
	user, err := i.UserRepo.FindByID(params.UserID)
	if err != nil {
		return nil, err
	}

	if err := i.AttendanceRepo.Create(attendance); err != nil {
		return nil, err
	}

	output := UserArriveOutput{
		User:       *user,
		Attendance: attendance,
	}

	return &output, err
}

func (i AttendanceInteractor) UserLeave(params UserLeaveInputData) (*UserLeaveOutput, error) {
	attendance := entity.Attendance{
		UserID: params.UserID,
		Type:   entity.Leave,
		At:     params.At,
	}
	user, err := i.UserRepo.FindByID(params.UserID)
	if err != nil {
		return nil, err
	}

	if err := i.AttendanceRepo.Create(attendance); err != nil {
		return nil, err
	}

	output := UserLeaveOutput{
		User:       *user,
		Attendance: attendance,
	}

	return &output, err
}

func (i AttendanceInteractor) GetByDate(date utility.SimpleDate) GetByDateOutput {
	attendances := i.AttendanceRepo.FindByDateWithUser(date)
	return attendances
}

func (i AttendanceInteractor) GetByDateRange(
	from utility.SimpleDate,
	to utility.SimpleDate,
) GetByDateRangeOutput {
	attendances := i.AttendanceRepo.FindByDateRangeWithUser(from, to)
	return attendances
}
