package category

import (
	"errors"
	"exploreur/backend/database"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name   string
	PostID int
}

var category = &Category{}

func AddCategory(name string) {
	db, err := gorm.Open(postgres.Open(database.GetEnv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.Create(&Category{Name: name})
}

func DeleteCategory(id int) {
	db, err := gorm.Open(postgres.Open(database.GetEnv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.Delete(&Category{}, id)
}

func UpdateCategoryName(name string, id int) {
	db, err := gorm.Open(postgres.Open(database.GetEnv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.Model(&Category{}).Where("id = ?", id).Update("name", name)
}

func GetUser(name string) (int, error) {
	db, err := gorm.Open(postgres.Open(database.GetEnv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	var singleCategory Category
	result := db.Select("id").Where("name = ?", name).First(&singleCategory)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			panic("user not found")
		}
		return 0, result.Error
	}
	return int(singleCategory.ID), nil
}
