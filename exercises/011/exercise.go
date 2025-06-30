package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

/*
Exercise 011: JSON処理とRESTful API

このエクササイズでは、JSON処理とRESTful APIの実装を学びます：

1. Task構造体とJSON操作
   - タスク管理システムの実装
   - JSONエンコード/デコード

2. RESTful APIエンドポイント
   - GET /tasks - 全タスク取得
   - GET /tasks/{id} - 特定タスク取得
   - POST /tasks - 新規タスク作成
   - PUT /tasks/{id} - タスク更新
   - DELETE /tasks/{id} - タスク削除

3. HTTPハンドラーの実装
   - リクエスト処理
   - レスポンス生成
   - エラーハンドリング

期待される動作:
- RESTful APIサーバーの起動
- JSON形式でのデータ操作
- 適切なHTTPステータスコードの返却
*/

func main() {
	fmt.Println("Exercise 011: JSON Processing and RESTful API")
	
	// タスクマネージャーを初期化
	tm := NewTaskManager()
	
	// サンプルタスクを追加
	task1 := Task{Title: "Learn Go", Description: "Study Go programming language", Status: "pending"}
	task2 := Task{Title: "Build API", Description: "Create RESTful API", Status: "in_progress"}
	
	id1 := tm.AddTask(task1)
	id2 := tm.AddTask(task2)
	
	fmt.Printf("Added task 1 with ID: %d\n", id1)
	fmt.Printf("Added task 2 with ID: %d\n", id2)
	
	// JSON操作のデモ
	fmt.Println("\n=== JSON Operations Demo ===")
	jsonData, err := json.Marshal(task1)
	if err != nil {
		log.Printf("JSON marshal error: %v", err)
	} else {
		fmt.Printf("Task 1 JSON: %s\n", jsonData)
	}
	
	var decodedTask Task
	err = json.Unmarshal(jsonData, &decodedTask)
	if err != nil {
		log.Printf("JSON unmarshal error: %v", err)
	} else {
		fmt.Printf("Decoded task: %+v\n", decodedTask)
	}
	
	// HTTPサーバーを開始
	fmt.Println("\n=== Starting HTTP Server ===")
	fmt.Println("Server starting on :8080")
	fmt.Println("Available endpoints:")
	fmt.Println("  GET    /tasks")
	fmt.Println("  GET    /tasks/{id}")
	fmt.Println("  POST   /tasks")
	fmt.Println("  PUT    /tasks/{id}")
	fmt.Println("  DELETE /tasks/{id}")
	
	server := NewTaskServer(tm)
	log.Fatal(http.ListenAndServe(":8080", server))
}

// Task構造体
type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"` // pending, in_progress, completed
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// TaskManager構造体
type TaskManager struct {
	tasks  map[int]Task
	nextID int
}

// NewTaskManager関数の実装
func NewTaskManager() *TaskManager {
	// 1. TaskManager構造体を初期化
	return &TaskManager{
	    // 2. tasksマップを初期化
		tasks:  make(map[int]Task),
	    // 3. nextIDを1に設定
		nextID: 1,
	}
}

// AddTask メソッドの実装
func (tm *TaskManager) AddTask(task Task) int {
	// 1. タスクにIDと作成日時を設定
	// 2. tasksマップに追加
	task.ID = tm.nextID
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()
	tm.tasks[tm.nextID] = task
	// 3. nextIDをインクリメント
	tm.nextID++
	// 4. 新しいIDを返す
	return task.ID
}

// GetAllTasks メソッドの実装
func (tm *TaskManager) GetAllTasks() []Task {
	// 1. tasksマップから全タスクを取得
	tasks := make([]Task, 0, len(tm.tasks))
	for _, task := range tm.tasks {
	    // 2. スライスに変換して返す
		tasks = append(tasks, task)
	}
	return tasks
}

// GetTaskByID メソッドの実装
func (tm *TaskManager) GetTaskByID(id int) (Task, bool) {
	// 1. tasksマップからIDでタスクを検索
	task, exists := tm.tasks[id]

	// 2. 存在する場合はタスクとtrueを返す
	// 3. 存在しない場合は空のタスクとfalseを返す
	return task, exists
}

// UpdateTask メソッドの実装
func (tm *TaskManager) UpdateTask(id int, task Task) bool {
	if _, exists := tm.tasks[id]; exists {
	    // 1. IDでタスクを検索
		task.ID = id
	    // 2. 存在する場合は更新日時を設定して更新
		task.UpdatedAt = time.Now()
		tm.tasks[id] = task
	    // 3. 成功時はtrueを返す
		return true
	}
	return false
}

// DeleteTask メソッドの実装
func (tm *TaskManager) DeleteTask(id int) bool {
	if _, exists := tm.tasks[id]; exists {
	// 1. tasksマップからIDでタスクを削除
		delete(tm.tasks, id)
	// 2. 削除成功時はtrueを返す
		return true
	}
	return false
}

// TaskServer構造体
type TaskServer struct {
	taskManager *TaskManager
}

// NewTaskServer関数の実装
func NewTaskServer(tm *TaskManager) *TaskServer {
	// TODO: 実装する
	return &TaskServer{
		taskManager: tm,
	}
}

// ServeHTTP メソッドの実装
func (ts *TaskServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// 1. URLパスを解析してルーティング
	path := r.URL.Path

	// 2. HTTPメソッドに応じて適切なハンドラーを呼び出し
	// 3. /tasks と /tasks/{id} のパターンを処理
	if path == "/tasks" {
		switch r.Method {
		case "GET":
			ts.handleGetTasks(w, r)
		case "POST":
			ts.handleCreateTask(w, r)
		default:
			writeErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
		}

	} else if strings.HasPrefix(path, "/tasks/") {
		id, err := extractIDFromPath(path)
		if err != nil {
			writeErrorResponse(w, http.StatusBadRequest, "Invalid task ID")
			return
		}
		
		switch r.Method {
		case "GET":
			ts.handleGetTask(w, r, id)
		case "PUT":
			ts.handleUpdateTask(w, r, id)
		case "DELETE":
			ts.handleDeleteTask(w, r, id)
		default:
			writeErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
		}
	} else {
		writeErrorResponse(w, http.StatusNotFound, "Not found")
	}
}

// handleGetTasks メソッドの実装
func (ts *TaskServer) handleGetTasks(w http.ResponseWriter, r *http.Request) {
	// 1. 全タスクを取得
	tasks := ts.taskManager.GetAllTasks()

	// 2. JSON形式でレスポンス
	// 3. Content-Typeヘッダーを設定
	writeJSONResponse(w, http.StatusOK, tasks)
}

// handleGetTask メソッドの実装
func (ts *TaskServer) handleGetTask(w http.ResponseWriter, r *http.Request, id int) {
	// 1. IDでタスクを取得
	task, exists := ts.taskManager.GetTaskByID(id)

	// 2. 存在しない場合は404エラー
	if !exists {
		writeErrorResponse(w, http.StatusNotFound, "Task not found")
		return
	}

	// 3. JSON形式でレスポンス
	writeJSONResponse(w, http.StatusOK, task)
}

// handleCreateTask メソッドの実装
func (ts *TaskServer) handleCreateTask(w http.ResponseWriter, r *http.Request) {
	// 1. リクエストボディからJSONを読み取り
	var task Task

	// 2. Task構造体にデコード
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		writeErrorResponse(w, http.StatusBadRequest, "Invalid JSON")
		return
	}
	id := ts.taskManager.AddTask(task)

	// 3. 新規タスクを作成
	createdTask, _ := ts.taskManager.GetTaskByID(id)

	// 4. 201ステータスでレスポンス
	writeJSONResponse(w, http.StatusCreated, createdTask)
}

// handleUpdateTask メソッドの実装
func (ts *TaskServer) handleUpdateTask(w http.ResponseWriter, r *http.Request, id int) {
	var task Task
	// 1. リクエストボディからJSONを読み取り
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		writeErrorResponse(w, http.StatusBadRequest, "Invalid JSON")
		return
	}
	
	// 2. タスクを更新
	success := ts.taskManager.UpdateTask(id, task)
	if !success {
		writeErrorResponse(w, http.StatusNotFound, "Task not found")
		return
	}
	updatedTask, _ := ts.taskManager.GetTaskByID(id)

	// 3. 更新後のタスクをレスポンス
	writeJSONResponse(w, http.StatusOK, updatedTask)
}

// handleDeleteTask メソッドの実装
func (ts *TaskServer) handleDeleteTask(w http.ResponseWriter, r *http.Request, id int) {
	// 1. タスクを削除
	success := ts.taskManager.DeleteTask(id)
	if !success {
		writeErrorResponse(w, http.StatusNotFound, "Task not found")
		return
	}
	// 2. 204ステータスでレスポンス
	w.WriteHeader(http.StatusNoContent)
}

// extractIDFromPath ヘルパー関数の実装
func extractIDFromPath(path string) (int, error) {
	// 1. パスから"/tasks/"を除去
	idStr := strings.TrimPrefix(path, "/tasks/")
	// 2. 残りの文字列を整数に変換
	if idStr == "" {
		return 0, fmt.Errorf("empty ID")
	}
	// 3. 変換エラーをハンドリング
	return strconv.Atoi(idStr)
}

// writeJSONResponse ヘルパー関数の実装
func writeJSONResponse(w http.ResponseWriter, status int, data interface{}) {
	// 1. Content-Typeをapplication/jsonに設定
	w.Header().Set("Content-Type", "application/json")
	// 2. ステータスコードを設定
	w.WriteHeader(status)
	// 3. データをJSONエンコードして書き込み
	json.NewEncoder(w).Encode(data)
}

// writeErrorResponse ヘルパー関数の実装
func writeErrorResponse(w http.ResponseWriter, status int, message string) {
	// 1. エラーメッセージの構造体を作成
	errorResponse := map[string]string{"error": message}
	// 2. writeJSONResponseを使用してエラーレスポンス
	writeJSONResponse(w, status, errorResponse)
}