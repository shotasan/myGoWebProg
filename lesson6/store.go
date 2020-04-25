package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

var Db *sql.DB

// DBの設定。Openだけでは接続するわけではない。init関数はmainより前に実行される
func init() {
	var err error
	Db, err = sql.Open("postgres", "user=gwp dbname=gwp password=gwp sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func Posts(limit int) (posts []Post, err error) {
	rows, err := Db.Query("select id, content, author from posts limit $1", limit)
	if err != nil {
		return
	}
	// クエリで呼び出したレコードを一行づつ読み込む。読み込むものがなくなるまで実行される
	for rows.Next() {
		post := Post{}
		err = rows.Scan(&post.Id, &post.Content, &post.Author)
		if err != nil {
			return
		}
		// 戻り値になるposts(Postのスライス型)に要素を追加する
		posts = append(posts, post)
	}
	rows.Close()
	return
}

// 投稿一件の取得
func GetPost(id int) (post Post, err error) {
	post = Post{}
	err = Db.QueryRow("select id, content, author from posts where id = $1", id).Scan(&post.Id, &post.Content, &post.Author)
	return
}

func (post *Post) Create() (err error) {
	statement := "insert into posts (content, author) values ($1, $2) returning id"
	// Db.PrepareでSQL文を実行するためのステートメントを準備する。Closeする必要がある。
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(post.Content, post.Author).Scan(&post.Id)
	return
}

func (post *Post) Update() (err error) {
	_, err = Db.Exec("update posts set content = $2, author = $3 where id = $1", post.Id, post.Content, post.Author)
	return
}

func (post *Post) Delete() (err error) {
	_, err = Db.Exec("delete from posts where id = $1", post.Id)
	return
}

func main() {
	post := Post{Content: "Hello Go", Author: "JonDo"}

	fmt.Println(post)
	post.Create()
	fmt.Println(post)

	readPost, _ := GetPost(post.Id)
	fmt.Println(readPost)

	readPost.Content = "Bonjour Java"
	readPost.Author = "Pierre"
	readPost.Update()
	fmt.Println(readPost)

	posts, _ := Posts(10)
	fmt.Println(posts)

	readPost.Delete()
	posts, _ = Posts(10)
	fmt.Println(posts)
}
