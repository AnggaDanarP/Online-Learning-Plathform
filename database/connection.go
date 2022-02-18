package database

import (
	"github.com/AnggaDanarP/Online-Learning-Plathform/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// Open connection to MySQL using root:yourpassword
	dsn := "root:@anggadanar10@tcp(127.0.0.1:3306)/onile_course?parseTime=true"
	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("could not connect to database")
	}

	DB = connection

	connection.AutoMigrate(&models.User{}, &models.Course{})


}