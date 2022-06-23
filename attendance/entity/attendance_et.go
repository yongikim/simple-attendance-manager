package entity

import "time"

type Attendance struct {
	ID     AttendanceID
	UserID UserID
	Type   AttendanceType
	At     time.Time
}

type AttendanceID uint

type AttendanceType string

const (
	Arrive = AttendanceType("arrive")
	Leave  = AttendanceType("leave")
)
