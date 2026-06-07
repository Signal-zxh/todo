package main

import (
	"html/template"
	"net/http"
)

type Todo struct {
	ID   int
	Task string
	Done bool
}

var todos = []Todo{}
var nextID = 1

func listHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: 渲染模板，显示所有任务
	tmpl, err := template.ParseFiles("templates/list.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, todos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: 解析表单，添加任务
}

func doneHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: 标记任务完成
}
func main() {
	http.HandleFunc("/", listHandler)
	http.HandleFunc("/add", addHandler)
	http.HandleFunc("/done", doneHandler)
	http.ListenAndServe(":8080", nil)
}
