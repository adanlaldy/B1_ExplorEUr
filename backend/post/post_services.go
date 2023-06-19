package post

import (
	"errors"
	"exploreur/backend/register"
	_ "exploreur/backend/register"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Content string
	UserID  int
}

var post = &Post{}

func AddPost(userID int, content string) {
	db, err := gorm.Open(postgres.Open(register.GetEnv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.Create(&Post{UserID: userID, Content: content})
}

func DeletePost(postID int) {
	db, err := gorm.Open(postgres.Open(register.GetEnv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.Delete(&Post{}, postID)
}

func Clear() {
	register.Db.Exec("DROP TABLE posts")
}

func UpdatePost(content string, postID int) {
	db, err := gorm.Open(postgres.Open(register.GetEnv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.Model(&Post{}).Where("id = ?", postID).Update("content", content)
}

func GetPost(content string) (int, error) {
	db, err := gorm.Open(postgres.Open(register.GetEnv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	var singlePost Post
	result := db.Select("id").Where("content = ?", content).First(&singlePost)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			panic("post not found")
		}
		return 0, result.Error
	}
	return int(singlePost.ID), nil
}

func ResetPostTable() {
	db, err := gorm.Open(postgres.Open(GetEnv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.Migrator().DropTable(&Post{})
	if err != nil {
		panic("problem to delete post table")
	}
	err = db.AutoMigrate(&Post{})
	if err != nil {
		panic("failed to auto migrate: ")
	}
}
