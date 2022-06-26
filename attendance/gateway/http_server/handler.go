package http_server

import (
	"net/http"
	"simple-attendance-manager/attendance/entity"
	"simple-attendance-manager/attendance/usecase"
	"simple-attendance-manager/attendance/utility"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	AttendanceUsecase usecase.AttendanceUsecase
	UserUsecase       usecase.UserUsecase
}

func SubmitHandlers(
	engine *gin.Engine,
	auc usecase.AttendanceUsecase,
	uuc usecase.UserUsecase,
) {
	handler := Handler{
		AttendanceUsecase: auc,
		UserUsecase:       uuc,
	}

	engine.GET("/attendance/today_all_users", handler.HandleTodayAllUsers)
	engine.GET("/attendance/today", handler.HandleToday)
	engine.GET("/attendance", handler.HandleDate)
	engine.POST("/users", handler.HandleCreateUser)
	engine.GET("/users/:id", handler.HandleGetUserByID)
	engine.POST("/users/:id/arrive", handler.HandleUserArrive)
	engine.POST("/users/:id/leave", handler.HandleUserLeave)
	engine.DELETE("/users/:id", handler.HandleDeleteUser)
}

// GET "/attendance/today_all_users"
func (h Handler) HandleTodayAllUsers(c *gin.Context) {
	now := time.Now()
	today := utility.SimpleDateFromTime(now)
	result :=
		h.UserUsecase.GetAllWithAttendancesByDate(today)

	c.JSON(http.StatusOK, result)
}

// GET "/attendance/today"
func (h Handler) HandleToday(c *gin.Context) {
	// Controller
	now := time.Now()
	date := utility.SimpleDate{
		Year:  now.Year(),
		Month: int(now.Month()),
		Day:   now.Day(),
	}
	result := h.AttendanceUsecase.GetByDate(date)

	// Presenter
	c.JSON(http.StatusOK, result)
}

// GET "/attendance?year=2020&month=12&day=1"
func (h Handler) HandleDate(c *gin.Context) {
	// Date Validation
	year, err := strconv.Atoi(c.Query("year"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "year is invalid")
		return
	}
	month, err := strconv.Atoi(c.Query("month"))
	if err != nil || !(0 < month && month < 13) {
		c.JSON(http.StatusBadRequest, "month is invalid")
		return
	}
	day, err := strconv.Atoi(c.Query("day"))
	if err != nil || !(0 < day && day < 32) {
		c.JSON(http.StatusBadRequest, "day is invalid")
		return
	}

	result := h.AttendanceUsecase.GetByDate(utility.SimpleDate{
		Year:  year,
		Month: int(month),
		Day:   day,
	})

	c.JSON(http.StatusOK, result)
}

// POST /users
type UserCreateRequest struct {
	Name  string `json:"name"`
	Grade int8   `json:"grade"`
}

func (h Handler) HandleCreateUser(c *gin.Context) {
	var request UserCreateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input := usecase.UserCreateInputData{
		Name:  request.Name,
		Grade: entity.Grade(request.Grade),
	}
	user, err := h.UserUsecase.Create(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, user)
}

// GET /users/:id
func (h Handler) HandleGetUserByID(c *gin.Context) {
	id_str := c.Param("id")
	if id_str == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing parameter: id"})
		return
	}

	id, err := strconv.Atoi(id_str)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.UserUsecase.GetByID(entity.UserID(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// POST /users/:id/arrive
func (h Handler) HandleUserArrive(c *gin.Context) {
	id_str := c.Param("id")
	if id_str == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing parameter: id"})
		return
	}

	id, err := strconv.Atoi(id_str)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	input := usecase.UserArriveInputData{
		UserID: entity.UserID(id),
		At:     time.Now().Local(),
	}
	res, err := h.AttendanceUsecase.UserArrive(input)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// POST /users/:id/leave
func (h Handler) HandleUserLeave(c *gin.Context) {
	id_str := c.Param("id")
	if id_str == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing parameter: id"})
		return
	}

	id, err := strconv.Atoi(id_str)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	input := usecase.UserLeaveInputData{
		UserID: entity.UserID(id),
		At:     time.Now().Local(),
	}
	res, err := h.AttendanceUsecase.UserLeave(input)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// DELETE /users/:id
func (h Handler) HandleDeleteUser(c *gin.Context) {
	id_str := c.Param("id")
	if id_str == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing parameter: id"})
		return
	}

	id, err := strconv.Atoi(id_str)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err = h.UserUsecase.Delete(entity.UserID(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, "User deleted")
}
