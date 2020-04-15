package main

import (
	"gowebprog/gowebprog/ch02/chitchat/data"
	"net/http"
	"text/template"
)

func index(write http.ResponseWriter, request *http.Request) {
	files := []string{
		"templates/layout.html",
		"templates/navbar.html",
		"templates/index.html"}
	templates := template.Must(template.ParseFiles(files...))
	threads, err := data.Threads()
	if err == nil {
		templates.ExecuteTemplate(w, "layout", threads)
	}
}
