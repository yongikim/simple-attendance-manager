package usecase

// import (
// 	"reflect"
// 	"simple-attendance-manager/db"
// 	"simple-attendance-manager/entity"
// 	"simple-attendance-manager/repository"
// 	"testing"
// 	"time"
// )

// func TestGetAllUsers(t *testing.T) {
// 	tx := db.ConnectMysqlTest().Begin()

// 	// Create test data
// 	test_users := []entity.User{
// 		{Name: "Test B4 User", Grade: 1},
// 		{Name: "Test M1 User", Grade: 2},
// 		{Name: "Test M2 User", Grade: 3},
// 	}
// 	tx.Create(&test_users)

// 	// Test application
// 	db_service := db.DBService{DB: tx}
// 	user_repo := repository.UserRepository{DBService: db_service}
// 	users := user_repo.GetCurrentUsers()

// 	var got_ids []uint
// 	for _, u := range users {
// 		got_ids = append(got_ids, u.ID)
// 	}
// 	var expected_ids []uint
// 	for _, u := range test_users {
// 		expected_ids = append(expected_ids, u.ID)
// 	}

// 	if !reflect.DeepEqual(got_ids, expected_ids) {
// 		t.Errorf("got %v, expected %v", got_ids, expected_ids)
// 	}

// 	tx.Rollback()
// }

// func TestGetAllUsersWithAttendanceAndLeaving(t *testing.T) {
// 	tx := db.ConnectMysqlTest().Begin()

// 	// Create test data
// 	test_users := []entity.User{
// 		{Name: "Test B4 User", Grade: 1},
// 		{Name: "Test M1 User", Grade: 2},
// 		{Name: "Test M2 User", Grade: 3},
// 	}
// 	tx.Create(&test_users)
// 	for _, u := range test_users {
// 		var attends []entity.Attendance
// 		var leavings []entity.Leaving
// 		for j := 0; j < 3; j++ {
// 			at := time.Date(2000, time.Month(1), j+1, 0, 0, 0, 0, time.Local)
// 			attends = append(attends, entity.Attendance{At: at})
// 			leavings = append(leavings, entity.Leaving{At: at})
// 		}
// 		tx.Model(&u).Association("Attendances").Append(attends)
// 		tx.Model(&u).Association("Leavings").Append(leavings)
// 	}

// 	// Test application
// 	db_service := db.DBService{DB: tx}
// 	user_repo := repository.UserRepository{DBService: db_service}
// 	users := user_repo.GetCurrentUsersWithLastAttendedAndLeft()

// 	expected_day := 3
// 	for _, user := range users {
// 		got_day := user.LastAttended.Day()
// 		if got_day != expected_day {
// 			t.Errorf("got %d, expected %d", got_day, expected_day)
// 		}
// 	}

// 	tx.Rollback()
// }
