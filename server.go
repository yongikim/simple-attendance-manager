package main

import (
	"github.com/gin-gonic/gin"

	"simple-attendance-manager/api"
	"simple-attendance-manager/db"
)

type Attendance struct {
	Name string `json:"name"`
	In   string `json:"in"`
	Out  string `json:"out"`
}

func main() {
	db.Connect()

	r := gin.Default()
	r.GET("/users", api.GetAllUsers)

	r.Run("localhost:3000")
}
