package main

import (
	"net/http"
)

func main() {
	// マルチプレクサの生成
	mux := http.NewServeMux()

	// 指定されたディレクトリからファイルを配信するハンドラ
	files := http.FileServer(http.Dir("/public"))
	// 作成したハンドラをマルチプレクサの関数Handleにわたす。
	// リクエストのURLパスからプレフィックスを削除する
	mux.Handle("/static", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", index)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
