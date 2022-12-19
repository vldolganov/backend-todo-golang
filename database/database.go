package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"serverForMyTodo/database/models"
)

var Connection *gorm.DB

func InitConnection() {

	dsn := "host=localhost user=dolganoffadmin password=dolganoffadmin dbname=todo_test_db port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.Todo{})

	Connection = db
}
