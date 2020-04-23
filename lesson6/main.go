package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

func main() {
	csvFile, err := os.Create("posts.csv")
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	allPosts := []Post{
		Post{Id: 1, Content: "Hello Go!", Author: "Ichiro"},
		Post{Id: 2, Content: "Hello Ruby!", Author: "Jiro"},
		Post{Id: 3, Content: "Hello JavaScript!", Author: "Saburo"},
		Post{Id: 4, Content: "Hello Java!", Author: "Shiro"},
	}

	// ライターの生成
	writer := csv.NewWriter(csvFile)
	for _, post := range allPosts {
		line := []string{strconv.Itoa(post.Id), post.Content, post.Author}
		// Writeメソッドでlineをファイルに書き込む。エラーが有ると変数errがnilでなくなるのでpanicが起こる
		err := writer.Write(line)
		if err != nil {
			panic(err)
		}
	}
	// バッファにあるすべてのデータをファイルに書き込むためのメソッド
	writer.Flush()

	file, err := os.Open("posts.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// リーダーの生成
	reader := csv.NewReader(file)
	// ファイル内のレコードにすべてのフィールドがなくても良い場合は−１を指定する
	reader.FieldsPerRecord = -1
	record, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	var posts []Post
	for _, item := range record {
		// ParseIntで文字列を数値に変換する。Idの値を取得する。第２引数が０の場合は文字列の先頭文字で基数を判定。基本は１０進数。第三引数は精度。
		id, _ := strconv.ParseInt(item[0], 0, 0)
		post := Post{Id: int(id), Content: item[1], Author: item[2]}
		posts = append(posts, post)
	}
	fmt.Println(posts[0].Id)
	fmt.Println(posts[0].Content)
	fmt.Println(posts[0].Author)
}
