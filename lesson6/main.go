package main

import (
	"fmt"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

var PostById map[int]*Post
var PostsByAuthor map[string][]*Post

func store(post Post) {
	PostById[post.Id] = &post
	// appendは必ず変数への代入が必要
	PostsByAuthor[post.Author] = append(PostsByAuthor[post.Author], &post)
}

func main() {
	PostById = make(map[int]*Post)
	PostsByAuthor = make(map[string][]*Post)

	post1 := Post{Id: 1, Content: "Hello World!", Author: "Ichiro"}
	post2 := Post{Id: 2, Content: "Hello Java!", Author: "Jiro"}
	post3 := Post{Id: 3, Content: "Hello Go!", Author: "Saburo"}
	post4 := Post{Id: 4, Content: "Hello Ruby!", Author: "Ichiro"}

	store(post1)
	store(post2)
	store(post3)
	store(post4)

	fmt.Printf("%v\n", PostById)
	fmt.Printf("%v\n", PostsByAuthor)

	fmt.Println(PostById[1])
	fmt.Println(PostById[2])

	for _, post := range PostsByAuthor["Ichiro"] {
		fmt.Println(post)
	}
	for _, post := range PostsByAuthor["Jiro"] {
		fmt.Println(post)
	}
}
