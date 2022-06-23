package entity

type UserID uint
type Grade int8

type User struct {
	ID    UserID
	Name  string
	Grade Grade
}
