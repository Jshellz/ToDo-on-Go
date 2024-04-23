package db

import (
	"ToDo/app/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const dsn = "postgresql://user:pass@localhost:5432/todo"

var (
	panicToDB      = "failed to connect database"
	panicToMigrate = "failed to migrate in database"
)

func ConnectToDB() *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(panicToDB)
	}

	if err := db.AutoMigrate(&model.Todo{}); err != nil {
		panic(panicToMigrate)
	}

	return db
}

func CreateModel(model *model.Todo) *gorm.DB {
	db := ConnectToDB()
	db.Create(&model)
	return db
}
