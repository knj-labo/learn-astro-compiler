package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUser(t *testing.T) {
	user := User{
		ID:    1,
		Name:  "Test User",
		Email: "test@example.com",
	}
	
	// JSON変換のテスト
	data, err := json.Marshal(user)
	if err != nil {
		t.Fatalf("Failed to marshal user: %v", err)
	}
	
	var unmarshaled User
	err = json.Unmarshal(data, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal user: %v", err)
	}
	
	if unmarshaled.ID != user.ID {
		t.Errorf("ID mismatch: got %d, want %d", unmarshaled.ID, user.ID)
	}
	if unmarshaled.Name != user.Name {
		t.Errorf("Name mismatch: got %s, want %s", unmarshaled.Name, user.Name)
	}
	if unmarshaled.Email != user.Email {
		t.Errorf("Email mismatch: got %s, want %s", unmarshaled.Email, user.Email)
	}
}

func TestUserManager(t *testing.T) {
	manager := &UserManager{}
	
	// 初期状態のテスト
	if len(manager.GetAllUsers()) != 0 {
		t.Error("Expected empty user list initially")
	}
	
	// ユーザー追加のテスト
	user1 := User{ID: 1, Name: "Alice", Email: "alice@example.com"}
	user2 := User{ID: 2, Name: "Bob", Email: "bob@example.com"}
	
	manager.AddUser(user1)
	manager.AddUser(user2)
	
	// 全ユーザー取得のテスト
	users := manager.GetAllUsers()
	if len(users) != 2 {
		t.Errorf("Expected 2 users, got %d", len(users))
	}
	
	// 特定ユーザー取得のテスト
	foundUser, found := manager.GetUser(1)
	if !found {
		t.Error("Expected to find user with ID 1")
	}
	if foundUser.Name != "Alice" {
		t.Errorf("Expected Alice, got %s", foundUser.Name)
	}
	
	// 存在しないユーザーのテスト
	_, found = manager.GetUser(999)
	if found {
		t.Error("Expected not to find user with ID 999")
	}
}

func TestHandleGetUsers(t *testing.T) {
	manager := &UserManager{}
	manager.AddUser(User{ID: 1, Name: "Alice", Email: "alice@example.com"})
	manager.AddUser(User{ID: 2, Name: "Bob", Email: "bob@example.com"})
	
	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatal(err)
	}
	
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(manager.handleGetUsers)
	handler.ServeHTTP(rr, req)
	
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, status)
	}
	
	contentType := rr.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("Expected Content-Type application/json, got %s", contentType)
	}
	
	var users []User
	err = json.Unmarshal(rr.Body.Bytes(), &users)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}
	
	if len(users) != 2 {
		t.Errorf("Expected 2 users, got %d", len(users))
	}
}

func TestHandleGetUser(t *testing.T) {
	manager := &UserManager{}
	manager.AddUser(User{ID: 1, Name: "Alice", Email: "alice@example.com"})
	
	// 存在するユーザーのテスト
	req, err := http.NewRequest("GET", "/users/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(manager.handleGetUser)
	handler.ServeHTTP(rr, req)
	
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, status)
	}
	
	var user User
	err = json.Unmarshal(rr.Body.Bytes(), &user)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}
	
	if user.Name != "Alice" {
		t.Errorf("Expected Alice, got %s", user.Name)
	}
	
	// 存在しないユーザーのテスト
	req, err = http.NewRequest("GET", "/users/999", nil)
	if err != nil {
		t.Fatal(err)
	}
	
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("Expected status %d, got %d", http.StatusNotFound, status)
	}
}

func TestHandleCreateUser(t *testing.T) {
	manager := &UserManager{}
	
	user := User{ID: 3, Name: "Charlie", Email: "charlie@example.com"}
	userJSON, err := json.Marshal(user)
	if err != nil {
		t.Fatal(err)
	}
	
	req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(userJSON))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(manager.handleCreateUser)
	handler.ServeHTTP(rr, req)
	
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Expected status %d, got %d", http.StatusCreated, status)
	}
	
	var createdUser User
	err = json.Unmarshal(rr.Body.Bytes(), &createdUser)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}
	
	if createdUser.Name != "Charlie" {
		t.Errorf("Expected Charlie, got %s", createdUser.Name)
	}
	
	// ユーザーが実際に追加されたかチェック
	foundUser, found := manager.GetUser(3)
	if !found {
		t.Error("Expected user to be added to manager")
	}
	if foundUser.Name != "Charlie" {
		t.Errorf("Expected Charlie, got %s", foundUser.Name)
	}
}