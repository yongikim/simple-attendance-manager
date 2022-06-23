package http

import (
	"net/http"
	"simple-attendance-manager/attendance/usecase"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type AttendanceHandler struct {
	AttendanceUsecase usecase.AttendanceUsecase
}

func NewAttendanceHandler(engine *gin.Engine, auc usecase.AttendanceUsecase) {
	handler := AttendanceHandler{
		AttendanceUsecase: auc,
	}

	engine.GET("/attendance/today", handler.HandleToday)
	engine.GET("/attendance", handler.HandleDate)
}

// GET "/attendance"
func (h AttendanceHandler) HandleToday(c *gin.Context) {
	// Controller
	now := time.Now()
	date := usecase.SimpleDate{
		Year:  int8(now.Year()),
		Month: int8(now.Month()),
		Day:   int8(now.Day()),
	}
	result := h.AttendanceUsecase.GetByDate(date)

	// Presenter
	c.JSON(http.StatusOK, result)
}

// GET "/attendance?year=2020&month=12&day=1"
func (h AttendanceHandler) HandleDate(c *gin.Context) {
	// Date Validation
	year, err := strconv.Atoi(c.Query("year"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "year is invalid")
	}
	month, err := strconv.Atoi(c.Query("month"))
	if err != nil || !(0 < month && month < 13) {
		c.JSON(http.StatusBadRequest, "month is invalid")
	}
	day, err := strconv.Atoi(c.Query("day"))
	if err != nil || !(0 < day && day < 32) {
		c.JSON(http.StatusBadRequest, "day is invalid")
	}

	result := h.AttendanceUsecase.GetByDate(usecase.SimpleDate{
		Year:  int8(year),
		Month: int8(month),
		Day:   int8(day),
	})

	c.JSON(http.StatusOK, result)
}
