package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

/*
Exercise 005: HTTP API とJSON処理

このエクササイズでは、Go言語でのHTTPサーバーとJSON処理を学びます：

1. User 構造体を定義する
   - ID int `json:"id"`
   - Name string `json:"name"`
   - Email string `json:"email"`

2. UserManager 構造体を実装する
   - users []User フィールドを持つ
   - AddUser(user User) メソッド
   - GetUser(id int) (User, bool) メソッド
   - GetAllUsers() []User メソッド

3. HTTP ハンドラー関数を実装する
   - handleGetUsers: GET /users - 全ユーザーをJSONで返す
   - handleGetUser: GET /users/{id} - 指定IDのユーザーをJSON で返す
   - handleCreateUser: POST /users - JSONからユーザーを作成

4. StartServer 関数を実装する
   - ポート8080でHTTPサーバーを起動
   - 上記のハンドラーを登録

期待される動作:
- GET /users → 全ユーザーのJSONリストを返す
- GET /users/1 → ID=1のユーザーのJSONを返す
- POST /users → リクエストボディのJSONからユーザーを作成
*/

func main() {
	fmt.Println("Exercise 005: HTTP API Server")
	
	// サンプルデータを追加
	manager := &UserManager{}
	manager.AddUser(User{ID: 1, Name: "Alice", Email: "alice@example.com"})
	manager.AddUser(User{ID: 2, Name: "Bob", Email: "bob@example.com"})
	
	fmt.Println("サーバーを http://localhost:8080 で起動します...")
	fmt.Println("テスト用URL:")
	fmt.Println("  GET  http://localhost:8080/users")
	fmt.Println("  GET  http://localhost:8080/users/1")
	fmt.Println("  POST http://localhost:8080/users")
	
	// サーバーを起動（このコメントアウトを外すと実際にサーバーが起動します）
	// StartServer(manager)
}

// User 構造体の定義
type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

// UserManager 構造体の定義
type UserManager struct {
    users []User // ユーザーのスライス
}

// AddUser メソッドの実装
func (um *UserManager) AddUser(user User) {
	um.users = append(um.users, user)
}

// GetUser メソッドの実装
func (um *UserManager) GetUser(id int) (User, bool) {
	for _, user := range um.users {
		if user.ID == id {
			return user, true
		}
	}
	return User{}, false
}

// GetAllUsers メソッドの実装
func (um *UserManager) GetAllUsers() []User {
	// TODO: 実装する
	return um.users
}

// handleGetUsers ハンドラーの実装
func (um *UserManager) handleGetUsers(w http.ResponseWriter, r *http.Request) {
	// 1. Content-Type を application/json に設定
	w.Header().Set("Content-Type", "application/json")
	// 2. um.GetAllUsers() で全ユーザーを取得
	users := um.GetAllUsers()
	// 3. json.Marshal() でJSONに変換
	jsonData, err := json.Marshal(users)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// 4. w.Write() でレスポンスを送信
	w.Write(jsonData)
}

// handleGetUser ハンドラーの実装
func (um *UserManager) handleGetUser(w http.ResponseWriter, r *http.Request) {
	// 1. URLパスからIDを抽出 (/users/123 → "123")
	path := strings.TrimPrefix(r.URL.Path, "/users/")
	// 2. strconv.Atoi() で文字列を数値に変換
	id, err := strconv.Atoi(path)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	
	// 3. um.GetUser() でユーザーを取得
	user, found := um.GetUser(id)

	// 4. 見つからない場合は404エラーを返す
	if !found {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	
	// 5. 見つかった場合はJSONで返す
	w.Header().Set("Content-Type", "application/json")
	jsonData, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}

// handleCreateUser ハンドラーの実装
func (um *UserManager) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	// 1. リクエストボディを読み取る
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	
	// 2. json.Unmarshal() でUserに変換
	var user User
	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	
	// 3. um.AddUser() でユーザーを追加
	um.AddUser(user)
	
	// 4. 作成されたユーザーをJSONで返す
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	jsonData, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}

// StartServer 関数の実装
func StartServer(manager *UserManager) {
	// 1. http.HandleFunc() でルートを登録
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			if r.URL.Path == "/users" {
				manager.handleGetUsers(w, r)
			} else {
				manager.handleGetUser(w, r)
			}
		case "POST":
			manager.handleCreateUser(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	
	// 2. http.ListenAndServe() でサーバーを起動
	http.ListenAndServe(":8080", nil)
}