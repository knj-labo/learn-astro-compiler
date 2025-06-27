package main

import (
	"os"
	"testing"
	"time"
)

func TestNewUserDB(t *testing.T) {
	dbPath := "test_new.db"
	defer os.Remove(dbPath)
	
	userDB, err := NewUserDB(dbPath)
	if err != nil {
		t.Fatalf("NewUserDB failed: %v", err)
	}
	defer userDB.Close()
	
	if userDB.db == nil {
		t.Error("Database connection is nil")
	}
}

func TestCreateTable(t *testing.T) {
	dbPath := "test_create_table.db"
	defer os.Remove(dbPath)
	
	userDB, err := NewUserDB(dbPath)
	if err != nil {
		t.Fatalf("NewUserDB failed: %v", err)
	}
	defer userDB.Close()
	
	err = userDB.CreateTable()
	if err != nil {
		t.Fatalf("CreateTable failed: %v", err)
	}
	
	// テーブルが作成されたかチェック
	var count int
	err = userDB.db.QueryRow("SELECT COUNT(*) FROM sqlite_master WHERE type='table' AND name='users'").Scan(&count)
	if err != nil {
		t.Fatalf("Failed to check table existence: %v", err)
	}
	
	if count != 1 {
		t.Error("Users table was not created")
	}
}

func TestCRUDOperations(t *testing.T) {
	dbPath := "test_crud.db"
	defer os.Remove(dbPath)
	
	userDB, err := NewUserDB(dbPath)
	if err != nil {
		t.Fatalf("NewUserDB failed: %v", err)
	}
	defer userDB.Close()
	
	err = userDB.CreateTable()
	if err != nil {
		t.Fatalf("CreateTable failed: %v", err)
	}
	
	// Create
	user := User{
		Name:  "Test User",
		Email: "test@example.com",
		Age:   25,
	}
	
	userID, err := userDB.CreateUser(user)
	if err != nil {
		t.Fatalf("CreateUser failed: %v", err)
	}
	
	if userID <= 0 {
		t.Error("Invalid user ID returned")
	}
	
	// Read by ID
	foundUser, err := userDB.GetUserByID(userID)
	if err != nil {
		t.Fatalf("GetUserByID failed: %v", err)
	}
	
	if foundUser.Name != user.Name {
		t.Errorf("Expected name %s, got %s", user.Name, foundUser.Name)
	}
	if foundUser.Email != user.Email {
		t.Errorf("Expected email %s, got %s", user.Email, foundUser.Email)
	}
	if foundUser.Age != user.Age {
		t.Errorf("Expected age %d, got %d", user.Age, foundUser.Age)
	}
	
	// Read all
	users, err := userDB.GetAllUsers()
	if err != nil {
		t.Fatalf("GetAllUsers failed: %v", err)
	}
	
	if len(users) != 1 {
		t.Errorf("Expected 1 user, got %d", len(users))
	}
	
	// Update
	foundUser.Name = "Updated User"
	foundUser.Age = 30
	err = userDB.UpdateUser(*foundUser)
	if err != nil {
		t.Fatalf("UpdateUser failed: %v", err)
	}
	
	// Verify update
	updatedUser, err := userDB.GetUserByID(userID)
	if err != nil {
		t.Fatalf("GetUserByID after update failed: %v", err)
	}
	
	if updatedUser.Name != "Updated User" {
		t.Errorf("Expected updated name 'Updated User', got %s", updatedUser.Name)
	}
	if updatedUser.Age != 30 {
		t.Errorf("Expected updated age 30, got %d", updatedUser.Age)
	}
	
	// Delete
	err = userDB.DeleteUser(userID)
	if err != nil {
		t.Fatalf("DeleteUser failed: %v", err)
	}
	
	// Verify deletion
	_, err = userDB.GetUserByID(userID)
	if err == nil {
		t.Error("Expected error when getting deleted user")
	}
	
	// Verify empty list
	users, err = userDB.GetAllUsers()
	if err != nil {
		t.Fatalf("GetAllUsers after delete failed: %v", err)
	}
	
	if len(users) != 0 {
		t.Errorf("Expected 0 users after delete, got %d", len(users))
	}
}

func TestGetUserByIDNotFound(t *testing.T) {
	dbPath := "test_not_found.db"
	defer os.Remove(dbPath)
	
	userDB, err := NewUserDB(dbPath)
	if err != nil {
		t.Fatalf("NewUserDB failed: %v", err)
	}
	defer userDB.Close()
	
	err = userDB.CreateTable()
	if err != nil {
		t.Fatalf("CreateTable failed: %v", err)
	}
	
	_, err = userDB.GetUserByID(999)
	if err == nil {
		t.Error("Expected error for non-existent user")
	}
}

func TestUpdateNonExistentUser(t *testing.T) {
	dbPath := "test_update_nonexistent.db"
	defer os.Remove(dbPath)
	
	userDB, err := NewUserDB(dbPath)
	if err != nil {
		t.Fatalf("NewUserDB failed: %v", err)
	}
	defer userDB.Close()
	
	err = userDB.CreateTable()
	if err != nil {
		t.Fatalf("CreateTable failed: %v", err)
	}
	
	user := User{
		ID:    999,
		Name:  "Non-existent",
		Email: "test@example.com",
		Age:   25,
	}
	
	err = userDB.UpdateUser(user)
	if err == nil {
		t.Error("Expected error when updating non-existent user")
	}
}

func TestDeleteNonExistentUser(t *testing.T) {
	dbPath := "test_delete_nonexistent.db"
	defer os.Remove(dbPath)
	
	userDB, err := NewUserDB(dbPath)
	if err != nil {
		t.Fatalf("NewUserDB failed: %v", err)
	}
	defer userDB.Close()
	
	err = userDB.CreateTable()
	if err != nil {
		t.Fatalf("CreateTable failed: %v", err)
	}
	
	err = userDB.DeleteUser(999)
	if err == nil {
		t.Error("Expected error when deleting non-existent user")
	}
}