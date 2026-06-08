package main

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

type Todo struct {
	ID   int
	Task string
	Done bool
}

var todos = []Todo{}
var nextID = 1

func listHandler(w http.ResponseWriter, r *http.Request) {
	// 渲染模板，显示所有任务
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
	//  解析表单，添加任务
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	task := r.FormValue("task")
	if strings.TrimSpace(task) == "" {
		http.Error(w, "Task cannot be empty", http.StatusBadRequest)
		return
	}

	newTodo := Todo{
		ID:   nextID,
		Task: task,
		Done: false,
	}
	todos = append(todos, newTodo)
	nextID++

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func doneHandler(w http.ResponseWriter, r *http.Request) {
	// 标记任务完成
	// 从URL提取ID
	idStr := strings.TrimPrefix(r.URL.Path, "/done/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	for i := range todos {
		if todos[i].ID == id {
			todos[i].Done = true
			break
		}
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
func main() {
	http.HandleFunc("/", listHandler)
	http.HandleFunc("/add", addHandler)
	http.HandleFunc("/done/", doneHandler)
	http.ListenAndServe(":8080", nil)
}
