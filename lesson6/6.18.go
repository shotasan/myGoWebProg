package main

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type Post struct {
	Id        int
	Content   string
	Author    string `sql:"not null"`
	Comments  []Comment
	CreatedAt time.Time
}

type Comment struct {
	Id        int
	Content   string
	Author    string `sql:"not null"`
	PostId    int
	CreatedAt time.Time
}

var Db *gorm.DB

func init() {
	var err error
	Db, err = gorm.Open("postgres", "user=gwp dbname=gwp password=gwp sslmode=disable")
	if err != nil {
		panic(err)
	}
	Db.AutoMigrate(&Post{}, &Comment{})
}

func main() {
	post := Post{Content: "Hello Gooo!", Author: "Jaaaaaaaav"}
	fmt.Println(post)

	Db.Create(&post)
	fmt.Println(post)

	comment := Comment{Content: "Goood post", Author: "Ruuuuuuuby"}
	fmt.Println(comment)
	Db.Model(&post).Association("Comments").Append(comment)
	fmt.Println(post)

	var readPost Post
	Db.Where("id = $1", 10).First(&readPost)

	var comments []Comment
	Db.Model(&readPost).Related(&comments)

	fmt.Println(comments)
	fmt.Println(comments[0])
}
