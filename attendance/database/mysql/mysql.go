package mysql

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlDB struct {
	// db *gorm.DB
}

// func (db *MysqlDB) Create(atdns *entity.Attendance) error {
// 	return nil
// }
// func (db *MysqlDB) Update(atdns *entity.Attendance) error {
// 	return nil

// }
// func (db *MysqlDB) GetAttendancesByRange(from time.Time, to time.Time) ([]entity.Attendance, error) {
// 	return nil, nil
// }

func Connect() *gorm.DB {
	user := os.Getenv("MYSQL_USER")
	pw := os.Getenv("MYSQL_PW")
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	db_name := os.Getenv("MYSQL_DB")

	dsn := user + ":" + pw + "@tcp(" + host + ":" + port + ")/" + db_name + "?charset=utf8mb4&parseTime=True&loc=Local"
	gdb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Database connection not established.")
	}

	// gdb.AutoMigrate(&entity.User{}, &entity.Attendance{}, &entity.Leaving{})

	return gdb
}

func ConnectTest() *gorm.DB {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	user := os.Getenv("MYSQL_USER_TEST")
	pw := os.Getenv("MYSQL_PW_TEST")
	host := os.Getenv("MYSQL_HOST_TEST")
	port := os.Getenv("MYSQL_PORT_TEST")
	db_name := os.Getenv("MYSQL_DB_TEST")

	dsn := user + ":" + pw + "@tcp(" + host + ":" + port + ")/" + db_name + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Database connection not established.")
	}

	// db.AutoMigrate(&entity.User{}, &entity.Attendance{}, &entity.Leaving{})

	return db
}

// func (s *MysqlDB) GetAllUsersWithLastAttendedAndLeft() []model.UserWithLastAttendedAndLeft {
// 	var users = []model.UserWithLastAttendedAndLeft{}

// 	s.db.Table("users").
// 		Select("users.id, users.name, users.grade, max(attendances.at) as last_attended, max(leavings.at) as last_left").
// 		Joins("left join attendances on attendances.user_id = users.id").
// 		Joins("left join leavings on leavings.user_id = users.id").
// 		Group("users.id").
// 		Scan(&users)

// 	return users
// }

// func (s *MysqlDB) GetCurrentUsers() []entity.User {
// 	var users []entity.User

// 	s.db.Where("grade >= ?", "0").Find(&users)

// 	return users
// }

// func (s *MysqlDB) CreateUser(user *entity.User) entity.User {
// 	return entity.User{}
// }
