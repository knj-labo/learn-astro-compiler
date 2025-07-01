package main

import (
	"fmt"
	"testing"
	"time"
)

func TestWorkerPoolManager(t *testing.T) {
	manager := NewWorkerPoolManager(2, 10)
	if manager == nil {
		t.Fatal("NewWorkerPoolManager returned nil")
	}

	manager.Start()
	defer manager.Stop()

	// Submit a test task
	task := Task{
		ID:   "test-1",
		Type: "square",
		Data: 5,
	}

	manager.SubmitTask(task)

	// Give some time for processing
	time.Sleep(100 * time.Millisecond)

	stats := manager.GetStatistics()
	if stats.TotalSubmitted == 0 {
		t.Error("Expected at least 1 submitted task")
	}
}

func TestWorker(t *testing.T) {
	taskQueue := make(chan Task, 1)
	resultQueue := make(chan TaskResult, 1)
	stopChan := make(chan bool, 1)

	worker := NewWorker(1, taskQueue, resultQueue, stopChan)
	if worker == nil {
		t.Fatal("NewWorker returned nil")
	}

	// Start worker
	go worker.Start()

	// Submit task
	task := Task{
		ID:   "test",
		Type: "square",
		Data: 4,
	}
	taskQueue <- task

	// Wait for result
	select {
	case result := <-resultQueue:
		if !result.Success {
			t.Errorf("Task failed: %s", result.Error)
		}
	case <-time.After(1 * time.Second):
		t.Error("Timeout waiting for result")
	}

	// Stop worker
	stopChan <- true
}

func TestBatchProcessor(t *testing.T) {
	processor := NewBatchProcessor(2, 5)
	if processor == nil {
		t.Fatal("NewBatchProcessor returned nil")
	}

	// Create test items
	items := []BatchItem{
		{ID: 1, Data: 10},
		{ID: 2, Data: 20},
		{ID: 3, Data: 30},
	}

	results := processor.ProcessBatch(items)

	if len(results) != len(items) {
		t.Errorf("Expected %d results, got %d", len(items), len(results))
	}
}

func TestCalculateSquare(t *testing.T) {
	result := calculateSquare(5)
	expected := 25

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestIsPrime(t *testing.T) {
	primeTests := []struct {
		number   int
		expected bool
	}{
		{2, true},
		{3, true},
		{4, false},
		{17, true},
		{25, false},
		{29, true},
	}

	for _, test := range primeTests {
		result := isPrime(test.number)
		if result != test.expected {
			t.Errorf("isPrime(%d) = %t, expected %t", test.number, result, test.expected)
		}
	}
}

func TestSimulateFileProcessing(t *testing.T) {
	result := simulateFileProcessing(1024)

	if result == "" {
		t.Error("Expected non-empty result from file processing")
	}
}

func TestSimulateNetworkRequest(t *testing.T) {
	result, err := simulateNetworkRequest("http://example.com")

	// Should return either a result or an error
	if result == "" && err == nil {
		t.Error("Expected either result or error from network request")
	}
}

func TestTaskProcessing(t *testing.T) {
	worker := &Worker{ID: 1}

	// Test square task
	squareTask := Task{
		ID:   "square-test",
		Type: "square",
		Data: 6,
	}

	result := worker.processTask(squareTask)
	if !result.Success {
		t.Errorf("Square task failed: %s", result.Error)
	}

	// Test prime task
	primeTask := Task{
		ID:   "prime-test",
		Type: "prime",
		Data: 17,
	}

	result = worker.processTask(primeTask)
	if !result.Success {
		t.Errorf("Prime task failed: %s", result.Error)
	}
}

func TestConcurrentProcessing(t *testing.T) {
	manager := NewWorkerPoolManager(3, 20)
	manager.Start()
	defer manager.Stop()

	// Submit multiple tasks
	for i := 0; i < 10; i++ {
		task := Task{
			ID:   fmt.Sprintf("concurrent-%d", i),
			Type: "square",
			Data: i,
		}
		manager.SubmitTask(task)
	}

	// Wait for processing
	time.Sleep(500 * time.Millisecond)

	stats := manager.GetStatistics()
	if stats.TotalSubmitted != 10 {
		t.Errorf("Expected 10 submitted tasks, got %d", stats.TotalSubmitted)
	}
}

func TestWorkerStatistics(t *testing.T) {
	manager := NewWorkerPoolManager(2, 5)
	stats := manager.GetStatistics()

	// Initial statistics should be zero
	if stats.TotalSubmitted != 0 {
		t.Error("Initial submitted count should be 0")
	}
	if stats.TotalCompleted != 0 {
		t.Error("Initial completed count should be 0")
	}
	if stats.TotalFailed != 0 {
		t.Error("Initial failed count should be 0")
	}
}