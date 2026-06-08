package main

import (
	"encoding/json"
	"html/template"
	"net/http"
	"os"
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

	//在 addHandler 里，添加任务后调用 saveTodos(todos)
	err := saveTodos(todos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
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

	//在 doneHandler 里，修改任务后调用 saveTodos(todos)
	err = saveTodos(todos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func loadTodos() {
	// 1. 读取 todos.json
	data, err := os.ReadFile("data/todos.json")
	if err != nil {
		// 2. 如果文件不存在，返回空的 todos 和 nextID=1
		if !os.IsNotExist(err) {
			todos = []Todo{}
			nextID = 1
		}
		return
	}
	// 3. 如果存在，解析 JSON 到 todos 切片
	err = json.Unmarshal(data, &todos)
	if err != nil {
		todos = []Todo{}
		nextID = 1
		return
	}
	// 4. 计算出下一个可用的 nextID（遍历 todos，取最大 ID + 1）
	maxID := 0
	for _, t := range todos {
		if t.ID > maxID {
			maxID = t.ID
		}
	}
	nextID = maxID + 1
}

func saveTodos(todos []Todo) error {
	// 确保data目录存在
	if err := os.MkdirAll("data", 0755); err != nil {
		return err
	}
	// 1. 把 todos 转成 JSON 格式（用 json.MarshalIndent 更美观）
	data, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		return err
	}
	// 2. 写入 todos.json 文件
	return os.WriteFile("data/todos.json", data, 0644)
}
func main() {
	// 在 main 函数开头，用 loadTodos() 初始化 todos 和 nextID
	loadTodos()
	http.HandleFunc("/", listHandler)
	http.HandleFunc("/add", addHandler)
	http.HandleFunc("/done/", doneHandler)
	http.ListenAndServe(":8080", nil)
}
