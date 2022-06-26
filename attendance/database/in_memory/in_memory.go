package in_memory

type NotFoundError struct {
	message string
}

func (e NotFoundError) Error() string {
	return e.message
}

type DataBase struct {
	Users       []UserRecord
	Attendances []AttendanceRecord
}

func NewInMemoryDB() *DataBase {
	return &DataBase{}
}
