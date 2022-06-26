package main

import (
	"simple-attendance-manager/attendance/database/in_memory"
	"simple-attendance-manager/attendance/gateway/http_server"
	"simple-attendance-manager/attendance/usecase"

	"github.com/gin-gonic/gin"
)

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

	a_repo := in_memory.NewAttendanceAPI(db)
	u_repo := in_memory.NewUserAPI(db)

	attendance_interactor := usecase.NewAttendanceInteractor(&a_repo, &u_repo)
	user_interactor := usecase.NewUserInteractor(&u_repo)

	http_server.SubmitHandlers(engine, attendance_interactor, user_interactor)

	engine.Run("localhost:3000")

}
