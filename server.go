package main

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

	// engine.Run("localhost:3000")
}
