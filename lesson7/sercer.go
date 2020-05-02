package main

import (
	"encoding/json"
	"log"
	"net/http"
	"path"
	"strconv"
)

type Post struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/post/", handleRequest)
	log.Println("server starting...")
	server.ListenAndServe()
}

// メソッドの内容により処理を振り分ける
func handleRequest(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = handleGet(w, r)
	case "POST":
		err = handlePost(w, r)
	case "PUT":
		err = handlePut(w, r)
	case "DELETE":
		err = handleDelete(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleGet(w http.ResponseWriter, r *http.Request) (err error) {
	// URLのパスを抽出し、関数path.Baseを使ってidを取得する
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	post, err := retrieve(id)
	if err != nil {
		return
	}
	// json.MarshalIndentで構造体postをJSONフォーマットのバイト列に変換する
	output, err := json.MarshalIndent(&post, "", "\t\t")
	if err != nil {
		return
	}
	// ヘッダの設定とResponseWriterへの書き込み
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

func handlePost(w http.ResponseWriter, r *http.Request) (err error) {
	len := r.ContentLength
	// バイト列を作成
	body := make([]byte, len)
	// バイト列にリクエストの本体を読み込み
	r.Body.Read(body)
	var post Post
	// バイト列を構造体Postに組み換え
	json.Unmarshal(body, &post)
	// データベースのレコードを作成
	err = post.create()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}

func handlePut(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	// DBからデータを取得し構造体Postに格納
	post, err := retrieve(id)
	if err != nil {
		return
	}
	len := r.ContentLength
	body := make([]byte, len)
	// リクエスト本体からJSONデータを読み出し
	r.Body.Read(body)
	// JSONデータを構造体Postに組み換え
	json.Unmarshal(body, &post)
	// DBを更新
	err = post.update()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}

func handleDelete(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	post, err := retrieve(id)
	if err != nil {
		return
	}
	err = post.delete()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}
