package main

import (
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
