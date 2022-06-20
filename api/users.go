package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"simple-attendance-manager/db"
)

type User struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Grade int8   `json:"grade"`
}

func GetAllUsers(c *gin.Context) {
	db := db.GetDB()

	var users []User
	db.Where("grade >= ?", "0").Find(&users)

	c.JSON(http.StatusOK, users)
}
