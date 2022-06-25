package main

import (
	"simple-attendance-manager/attendance/database/in_memory"
	"simple-attendance-manager/attendance/gateway/http_server"
	"simple-attendance-manager/attendance/usecase"

	"github.com/gin-gonic/gin"
)

type Attendance struct {
	Name string `json:"name"`
	In   string `json:"in"`
	Out  string `json:"out"`
}

func main() {
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	// mysql := db.ConnectMysql()
	// db_service := db.DBService{DB: mysql}

	// user_repo := repository.UserRepository{DBService: db_service}

	// engine := gin.Default()
	// api := api.API{UserRepo: user_repo}
	// api.Route(engine)

	engine := gin.Default()
	db := in_memory.NewInMemoryDB()
	attendance_usecase := usecase.AttendanceInteractor{
		DataAccess: db,
	}
	http_server.NewAttendanceHandler(engine, attendance_usecase)

	engine.Run("localhost:3000")

}
