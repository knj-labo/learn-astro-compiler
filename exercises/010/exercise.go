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
// SQLiteデータベースへの接続を行い、UserDBインスタンスを作成する
func NewUserDB(dbPath string) (*UserDB, error) {
	// SQLiteデータベースへの接続を開く
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}
	
	// データベース接続の正常性を確認
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}
	
	// UserDB構造体を初期化して返す
	return &UserDB{db: db}, nil
}

// Close メソッドの実装
// データベース接続を閉じる
func (udb *UserDB) Close() error {
	// データベース接続が存在する場合のみ閉じる
	if udb.db != nil {
		return udb.db.Close()
	}
	return nil
}

// CreateTable メソッドの実装
// usersテーブルを作成する（存在しない場合のみ）
func (udb *UserDB) CreateTable() error {
	// usersテーブルの作成SQL
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL,
		age INTEGER NOT NULL,
		created_at DATETIME NOT NULL
	)
	`
	
	// SQLを実行してテーブルを作成
	_, err := udb.db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to create table: %w", err)
	}
	
	return nil
}

// CreateUser メソッドの実装
// 新しいユーザーをデータベースに作成し、生成されたIDを返す
func (udb *UserDB) CreateUser(user User) (int64, error) {
	// ユーザー挿入用のSQL文
	query := `INSERT INTO users (name, email, age, created_at) VALUES (?, ?, ?, ?)`
	
	// prepared statementを作成（SQLインジェクション攻撃を防ぐ）
	stmt, err := udb.db.Prepare(query)
	if err != nil {
		return 0, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()
	
	// ユーザーデータを挿入（created_atは現在時刻を設定）
	result, err := stmt.Exec(user.Name, user.Email, user.Age, time.Now())
	if err != nil {
		return 0, fmt.Errorf("failed to execute insert: %w", err)
	}
	
	// 自動生成されたIDを取得
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to get last insert id: %w", err)
	}
	
	return id, nil
}

// GetAllUsers メソッドの実装
// データベースの全ユーザーを取得してスライスで返す
func (udb *UserDB) GetAllUsers() ([]User, error) {
	// 全ユーザー取得用のSQL文
	query := `SELECT id, name, email, age, created_at FROM users`
	
	// クエリを実行して結果セットを取得
	rows, err := udb.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query users: %w", err)
	}
	defer rows.Close()
	
	// ユーザースライスを初期化
	var users []User
	// 結果セットを一行ずつ処理
	for rows.Next() {
		var user User
		// 各列の値をUser構造体にスキャン
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Age, &user.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}
		// スライスにユーザーを追加
		users = append(users, user)
	}
	
	// イテレーション中のエラーをチェック
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during row iteration: %w", err)
	}
	
	return users, nil
}

// GetUserByID メソッドの実装
// 指定されたIDのユーザーを取得する
func (udb *UserDB) GetUserByID(id int64) (*User, error) {
	// ID指定でユーザーを取得するSQL文
	query := `SELECT id, name, email, age, created_at FROM users WHERE id = ?`
	
	var user User
	// 単一行の結果を取得してUser構造体にスキャン
	err := udb.db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email, &user.Age, &user.CreatedAt)
	if err != nil {
		// ユーザーが見つからない場合のエラーハンドリング
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user with id %d not found", id)
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	
	// ユーザーのポインタを返す
	return &user, nil
}

// UpdateUser メソッドの実装
// 指定されたユーザーの情報を更新する
func (udb *UserDB) UpdateUser(user User) error {
	// ユーザー情報更新用のSQL文
	query := `UPDATE users SET name = ?, email = ?, age = ? WHERE id = ?`
	
	// prepared statementを作成
	stmt, err := udb.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()
	
	// 更新を実行
	result, err := stmt.Exec(user.Name, user.Email, user.Age, user.ID)
	if err != nil {
		return fmt.Errorf("failed to execute update: %w", err)
	}
	
	// 更新された行数を取得
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	
	// 更新された行が0の場合はユーザーが存在しない
	if rowsAffected == 0 {
		return fmt.Errorf("user with id %d not found", user.ID)
	}
	
	return nil
}

// DeleteUser メソッドの実装
// 指定されたIDのユーザーを削除する
func (udb *UserDB) DeleteUser(id int64) error {
	// ユーザー削除用のSQL文
	query := `DELETE FROM users WHERE id = ?`
	
	// prepared statementを作成
	stmt, err := udb.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()
	
	// 削除を実行
	result, err := stmt.Exec(id)
	if err != nil {
		return fmt.Errorf("failed to execute delete: %w", err)
	}
	
	// 削除された行数を取得
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	
	// 削除された行が0の場合はユーザーが存在しない
	if rowsAffected == 0 {
		return fmt.Errorf("user with id %d not found", id)
	}
	
	return nil
}