package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTaskManager(t *testing.T) {
	tm := NewTaskManager()
	if tm == nil {
		t.Fatal("NewTaskManager returned nil")
	}

	// Test AddTask
	task := Task{
		Title:       "Test Task",
		Description: "Test Description",
		Status:      "pending",
	}

	id := tm.AddTask(task)
	if id <= 0 {
		t.Error("AddTask should return positive ID")
	}

	// Test GetTaskByID
	retrievedTask, exists := tm.GetTaskByID(id)
	if !exists {
		t.Error("Task should exist")
	}

	if retrievedTask.Title != task.Title {
		t.Errorf("Expected title %s, got %s", task.Title, retrievedTask.Title)
	}

	// Test GetAllTasks
	tasks := tm.GetAllTasks()
	if len(tasks) != 1 {
		t.Errorf("Expected 1 task, got %d", len(tasks))
	}

	// Test UpdateTask
	updatedTask := Task{
		Title:       "Updated Task",
		Description: "Updated Description",
		Status:      "completed",
	}

	success := tm.UpdateTask(id, updatedTask)
	if !success {
		t.Error("UpdateTask should succeed")
	}

	// Verify update
	retrievedTask, exists = tm.GetTaskByID(id)
	if !exists {
		t.Error("Task should still exist after update")
	}

	if retrievedTask.Title != updatedTask.Title {
		t.Errorf("Expected updated title %s, got %s", updatedTask.Title, retrievedTask.Title)
	}

	// Test DeleteTask
	success = tm.DeleteTask(id)
	if !success {
		t.Error("DeleteTask should succeed")
	}

	// Verify deletion
	_, exists = tm.GetTaskByID(id)
	if exists {
		t.Error("Task should not exist after deletion")
	}
}

func TestTaskServer(t *testing.T) {
	tm := NewTaskManager()
	server := NewTaskServer(tm)

	// Test GET /tasks (empty)
	req := httptest.NewRequest("GET", "/tasks", nil)
	w := httptest.NewRecorder()
	server.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	var tasks []Task
	err := json.Unmarshal(w.Body.Bytes(), &tasks)
	if err != nil {
		t.Errorf("Failed to decode JSON response: %v", err)
	}

	if len(tasks) != 0 {
		t.Errorf("Expected 0 tasks, got %d", len(tasks))
	}
}

func TestCreateTask(t *testing.T) {
	tm := NewTaskManager()
	server := NewTaskServer(tm)

	task := Task{
		Title:       "Test Task",
		Description: "Test Description",
		Status:      "pending",
	}

	taskJSON, _ := json.Marshal(task)
	req := httptest.NewRequest("POST", "/tasks", bytes.NewBuffer(taskJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	server.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status 201, got %d", w.Code)
	}

	var createdTask Task
	err := json.Unmarshal(w.Body.Bytes(), &createdTask)
	if err != nil {
		t.Errorf("Failed to decode JSON response: %v", err)
	}

	if createdTask.ID <= 0 {
		t.Error("Created task should have positive ID")
	}

	if createdTask.Title != task.Title {
		t.Errorf("Expected title %s, got %s", task.Title, createdTask.Title)
	}
}

func TestGetTaskByID(t *testing.T) {
	tm := NewTaskManager()
	server := NewTaskServer(tm)

	// Create a task first
	task := Task{
		Title:       "Test Task",
		Description: "Test Description",
		Status:      "pending",
	}
	id := tm.AddTask(task)

	// Test GET /tasks/{id}
	req := httptest.NewRequest("GET", "/tasks/"+string(rune(id+'0')), nil)
	w := httptest.NewRecorder()
	server.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

func TestUpdateTask(t *testing.T) {
	tm := NewTaskManager()
	server := NewTaskServer(tm)

	// Create a task first
	task := Task{
		Title:       "Test Task",
		Description: "Test Description",
		Status:      "pending",
	}
	id := tm.AddTask(task)

	// Update the task
	updatedTask := Task{
		Title:       "Updated Task",
		Description: "Updated Description",
		Status:      "completed",
	}

	taskJSON, _ := json.Marshal(updatedTask)
	req := httptest.NewRequest("PUT", "/tasks/"+string(rune(id+'0')), bytes.NewBuffer(taskJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	server.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

func TestDeleteTask(t *testing.T) {
	tm := NewTaskManager()
	server := NewTaskServer(tm)

	// Create a task first
	task := Task{
		Title:       "Test Task",
		Description: "Test Description",
		Status:      "pending",
	}
	id := tm.AddTask(task)

	// Delete the task
	req := httptest.NewRequest("DELETE", "/tasks/"+string(rune(id+'0')), nil)
	w := httptest.NewRecorder()
	server.ServeHTTP(w, req)

	if w.Code != http.StatusNoContent {
		t.Errorf("Expected status 204, got %d", w.Code)
	}

	// Verify deletion
	_, exists := tm.GetTaskByID(id)
	if exists {
		t.Error("Task should not exist after deletion")
	}
}

func TestGetNonExistentTask(t *testing.T) {
	tm := NewTaskManager()
	server := NewTaskServer(tm)

	req := httptest.NewRequest("GET", "/tasks/999", nil)
	w := httptest.NewRecorder()
	server.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status 404, got %d", w.Code)
	}
}

func TestInvalidTaskID(t *testing.T) {
	tm := NewTaskManager()
	server := NewTaskServer(tm)

	req := httptest.NewRequest("GET", "/tasks/invalid", nil)
	w := httptest.NewRecorder()
	server.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400, got %d", w.Code)
	}
}

func TestExtractIDFromPath(t *testing.T) {
	tests := []struct {
		path        string
		expectedID  int
		expectError bool
	}{
		{"/tasks/1", 1, false},
		{"/tasks/123", 123, false},
		{"/tasks/invalid", 0, true},
		{"/tasks/", 0, true},
	}

	for _, test := range tests {
		id, err := extractIDFromPath(test.path)
		if test.expectError {
			if err == nil {
				t.Errorf("Expected error for path %s", test.path)
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error for path %s: %v", test.path, err)
			}
			if id != test.expectedID {
				t.Errorf("Expected ID %d for path %s, got %d", test.expectedID, test.path, id)
			}
		}
	}
}