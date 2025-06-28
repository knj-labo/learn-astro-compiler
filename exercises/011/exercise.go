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
	// TODO: 実装する
	// ヒント:
	// 1. TaskManager構造体を初期化
	// 2. tasksマップを初期化
	// 3. nextIDを1に設定
	return nil
}

// AddTask メソッドの実装
func (tm *TaskManager) AddTask(task Task) int {
	// TODO: 実装する
	// ヒント:
	// 1. タスクにIDと作成日時を設定
	// 2. tasksマップに追加
	// 3. nextIDをインクリメント
	// 4. 新しいIDを返す
	return 0
}

// GetAllTasks メソッドの実装
func (tm *TaskManager) GetAllTasks() []Task {
	// TODO: 実装する
	// ヒント:
	// 1. tasksマップから全タスクを取得
	// 2. スライスに変換して返す
	return nil
}

// GetTaskByID メソッドの実装
func (tm *TaskManager) GetTaskByID(id int) (Task, bool) {
	// TODO: 実装する
	// ヒント:
	// 1. tasksマップからIDでタスクを検索
	// 2. 存在する場合はタスクとtrueを返す
	// 3. 存在しない場合は空のタスクとfalseを返す
	return Task{}, false
}

// UpdateTask メソッドの実装
func (tm *TaskManager) UpdateTask(id int, task Task) bool {
	// TODO: 実装する
	// ヒント:
	// 1. IDでタスクを検索
	// 2. 存在する場合は更新日時を設定して更新
	// 3. 成功時はtrueを返す
	return false
}

// DeleteTask メソッドの実装
func (tm *TaskManager) DeleteTask(id int) bool {
	// TODO: 実装する
	// ヒント:
	// 1. tasksマップからIDでタスクを削除
	// 2. 削除成功時はtrueを返す
	return false
}

// TaskServer構造体
type TaskServer struct {
	taskManager *TaskManager
}

// NewTaskServer関数の実装
func NewTaskServer(tm *TaskManager) *TaskServer {
	// TODO: 実装する
	return nil
}

// ServeHTTP メソッドの実装
func (ts *TaskServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO: 実装する
	// ヒント:
	// 1. URLパスを解析してルーティング
	// 2. HTTPメソッドに応じて適切なハンドラーを呼び出し
	// 3. /tasks と /tasks/{id} のパターンを処理
}

// handleGetTasks メソッドの実装
func (ts *TaskServer) handleGetTasks(w http.ResponseWriter, r *http.Request) {
	// TODO: 実装する
	// ヒント:
	// 1. 全タスクを取得
	// 2. JSON形式でレスポンス
	// 3. Content-Typeヘッダーを設定
}

// handleGetTask メソッドの実装
func (ts *TaskServer) handleGetTask(w http.ResponseWriter, r *http.Request, id int) {
	// TODO: 実装する
	// ヒント:
	// 1. IDでタスクを取得
	// 2. 存在しない場合は404エラー
	// 3. JSON形式でレスポンス
}

// handleCreateTask メソッドの実装
func (ts *TaskServer) handleCreateTask(w http.ResponseWriter, r *http.Request) {
	// TODO: 実装する
	// ヒント:
	// 1. リクエストボディからJSONを読み取り
	// 2. Task構造体にデコード
	// 3. 新規タスクを作成
	// 4. 201ステータスでレスポンス
}

// handleUpdateTask メソッドの実装
func (ts *TaskServer) handleUpdateTask(w http.ResponseWriter, r *http.Request, id int) {
	// TODO: 実装する
	// ヒント:
	// 1. リクエストボディからJSONを読み取り
	// 2. タスクを更新
	// 3. 更新後のタスクをレスポンス
}

// handleDeleteTask メソッドの実装
func (ts *TaskServer) handleDeleteTask(w http.ResponseWriter, r *http.Request, id int) {
	// TODO: 実装する
	// ヒント:
	// 1. タスクを削除
	// 2. 204ステータスでレスポンス
}

// extractIDFromPath ヘルパー関数の実装
func extractIDFromPath(path string) (int, error) {
	// TODO: 実装する
	// ヒント:
	// 1. パスから"/tasks/"を除去
	// 2. 残りの文字列を整数に変換
	// 3. 変換エラーをハンドリング
	return 0, nil
}

// writeJSONResponse ヘルパー関数の実装
func writeJSONResponse(w http.ResponseWriter, status int, data interface{}) {
	// TODO: 実装する
	// ヒント:
	// 1. Content-Typeをapplication/jsonに設定
	// 2. ステータスコードを設定
	// 3. データをJSONエンコードして書き込み
}

// writeErrorResponse ヘルパー関数の実装
func writeErrorResponse(w http.ResponseWriter, status int, message string) {
	// TODO: 実装する
	// ヒント:
	// 1. エラーメッセージの構造体を作成
	// 2. writeJSONResponseを使用してエラーレスポンス
}