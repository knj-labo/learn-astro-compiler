package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

/*
Exercise 010: データベース操作とCRUD

このエクササイズでは、Goでのデータベース操作とCRUD（Create, Read, Update, Delete）を学びます：

1. UserDB 構造体を実装する
   - SQLiteデータベースへの接続
   - ユーザーテーブルの作成
   - CRUD操作の実装

2. Create 操作を実装する
   - 新しいユーザーの作成
   - prepared statement の使用

3. Read 操作を実装する
   - 全ユーザーの取得
   - IDによる単一ユーザーの取得

4. Update 操作を実装する
   - 既存ユーザーの更新

5. Delete 操作を実装する
   - ユーザーの削除

期待される動作:
- データベースの初期化とテーブル作成
- ユーザーの作成、読み取り、更新、削除
- エラーハンドリングの適切な実装
*/

func main() {
	fmt.Println("Exercise 010: Database Operations and CRUD")
	
	// データベース接続
	userDB, err := NewUserDB("test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer userDB.Close()
	
	// テーブル作成
	err = userDB.CreateTable()
	if err != nil {
		log.Fatal(err)
	}
	
	// ユーザー作成
	fmt.Println("=== Creating Users ===")
	user1 := User{Name: "Alice", Email: "alice@example.com", Age: 30}
	user2 := User{Name: "Bob", Email: "bob@example.com", Age: 25}
	
	id1, err := userDB.CreateUser(user1)
	if err != nil {
		log.Printf("Error creating user1: %v", err)
	} else {
		fmt.Printf("Created user1 with ID: %d\n", id1)
	}
	
	id2, err := userDB.CreateUser(user2)
	if err != nil {
		log.Printf("Error creating user2: %v", err)
	} else {
		fmt.Printf("Created user2 with ID: %d\n", id2)
	}
	
	// 全ユーザー取得
	fmt.Println("\n=== All Users ===")
	users, err := userDB.GetAllUsers()
	if err != nil {
		log.Printf("Error getting all users: %v", err)
	} else {
		for _, user := range users {
			fmt.Printf("User: ID=%d, Name=%s, Email=%s, Age=%d\n", 
				user.ID, user.Name, user.Email, user.Age)
		}
	}
	
	// 単一ユーザー取得
	fmt.Println("\n=== Get User by ID ===")
	user, err := userDB.GetUserByID(id1)
	if err != nil {
		log.Printf("Error getting user by ID: %v", err)
	} else {
		fmt.Printf("Found user: %+v\n", *user)
	}
	
	// ユーザー更新
	fmt.Println("\n=== Update User ===")
	user.Age = 31
	user.Email = "alice.updated@example.com"
	err = userDB.UpdateUser(*user)
	if err != nil {
		log.Printf("Error updating user: %v", err)
	} else {
		fmt.Printf("Updated user: %+v\n", *user)
	}
	
	// ユーザー削除
	fmt.Println("\n=== Delete User ===")
	err = userDB.DeleteUser(id2)
	if err != nil {
		log.Printf("Error deleting user: %v", err)
	} else {
		fmt.Printf("Deleted user with ID: %d\n", id2)
	}
	
	// 削除後の全ユーザー取得
	fmt.Println("\n=== All Users After Delete ===")
	users, err = userDB.GetAllUsers()
	if err != nil {
		log.Printf("Error getting all users: %v", err)
	} else {
		for _, user := range users {
			fmt.Printf("User: ID=%d, Name=%s, Email=%s, Age=%d\n", 
				user.ID, user.Name, user.Email, user.Age)
		}
	}
}

// User 構造体
type User struct {
	ID        int       `db:"id"`
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	Age       int       `db:"age"`
	CreatedAt time.Time `db:"created_at"`
}

// UserDB 構造体
type UserDB struct {
	db *sql.DB
}

// NewUserDB 関数の実装
func NewUserDB(dbPath string) (*UserDB, error) {
	// TODO: 実装する
	// ヒント:
	// 1. sql.Open() でデータベースに接続
	// 2. db.Ping() で接続をテスト
	// 3. UserDB構造体を返す
	return nil, nil
}

// Close メソッドの実装
func (udb *UserDB) Close() error {
	// TODO: 実装する
	return nil
}

// CreateTable メソッドの実装
func (udb *UserDB) CreateTable() error {
	// TODO: 実装する
	// ヒント:
	// 1. CREATE TABLE IF NOT EXISTS のSQL文を作成
	// 2. id (INTEGER PRIMARY KEY), name (TEXT), email (TEXT), age (INTEGER), created_at (DATETIME)
	// 3. db.Exec() でテーブルを作成
	return nil
}

// CreateUser メソッドの実装
func (udb *UserDB) CreateUser(user User) (int64, error) {
	// TODO: 実装する
	// ヒント:
	// 1. INSERT INTOのSQL文とprepared statementを使用
	// 2. time.Now() で現在時刻を設定
	// 3. result.LastInsertId() で新しいIDを取得
	return 0, nil
}

// GetAllUsers メソッドの実装
func (udb *UserDB) GetAllUsers() ([]User, error) {
	// TODO: 実装する
	// ヒント:
	// 1. SELECT * FROM users のクエリを実行
	// 2. rows.Next() でループ
	// 3. rows.Scan() で各フィールドをスキャン
	return nil, nil
}

// GetUserByID メソッドの実装
func (udb *UserDB) GetUserByID(id int64) (*User, error) {
	// TODO: 実装する
	// ヒント:
	// 1. SELECT * FROM users WHERE id = ? のクエリ
	// 2. QueryRow() を使用
	// 3. Scan() で結果を取得
	// 4. sql.ErrNoRows をチェック
	return nil, nil
}

// UpdateUser メソッドの実装
func (udb *UserDB) UpdateUser(user User) error {
	// TODO: 実装する
	// ヒント:
	// 1. UPDATE users SET name=?, email=?, age=? WHERE id=? のクエリ
	// 2. prepared statement を使用
	// 3. result.RowsAffected() で更新件数をチェック
	return nil
}

// DeleteUser メソッドの実装
func (udb *UserDB) DeleteUser(id int64) error {
	// TODO: 実装する
	// ヒント:
	// 1. DELETE FROM users WHERE id=? のクエリ
	// 2. prepared statement を使用
	// 3. result.RowsAffected() で削除件数をチェック
	return nil
}