package main

import (
	"errors"
	"testing"
	"time"
)

func TestWorkerPool(t *testing.T) {
	tasks := []Task{
		{ID: 1, Data: "Task 1"},
		{ID: 2, Data: "Task 2"},
		{ID: 3, Data: "Task 3"},
	}
	
	results := WorkerPool(2, tasks)
	
	if len(results) != len(tasks) {
		t.Errorf("Expected %d results, got %d", len(tasks), len(results))
	}
	
	// 結果にすべてのタスクIDが含まれているかチェック
	taskIDs := make(map[int]bool)
	for _, result := range results {
		taskIDs[result.TaskID] = true
	}
	
	for _, task := range tasks {
		if !taskIDs[task.ID] {
			t.Errorf("Task ID %d not found in results", task.ID)
		}
	}
}

func TestWorkerPoolEmpty(t *testing.T) {
	tasks := []Task{}
	results := WorkerPool(2, tasks)
	
	if len(results) != 0 {
		t.Errorf("Expected 0 results for empty tasks, got %d", len(results))
	}
}

func TestTimeoutOperationSuccess(t *testing.T) {
	start := time.Now()
	err := TimeoutOperation(1*time.Second, func() error {
		time.Sleep(100 * time.Millisecond)
		return nil
	})
	duration := time.Since(start)
	
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	
	if duration > 500*time.Millisecond {
		t.Errorf("Operation took too long: %v", duration)
	}
}

func TestTimeoutOperationTimeout(t *testing.T) {
	start := time.Now()
	err := TimeoutOperation(200*time.Millisecond, func() error {
		time.Sleep(1 * time.Second)
		return nil
	})
	duration := time.Since(start)
	
	if err == nil {
		t.Error("Expected timeout error, got nil")
	}
	
	// タイムアウト時間より少し長い程度であることを確認
	if duration > 400*time.Millisecond {
		t.Errorf("Timeout took too long: %v", duration)
	}
}

func TestTimeoutOperationError(t *testing.T) {
	expectedErr := errors.New("operation failed")
	err := TimeoutOperation(1*time.Second, func() error {
		return expectedErr
	})
	
	if err != expectedErr {
		t.Errorf("Expected %v, got %v", expectedErr, err)
	}
}

func TestRateLimiter(t *testing.T) {
	count := 0
	interval := 100 * time.Millisecond
	executions := 3
	
	start := time.Now()
	RateLimiter(interval, executions, func(i int) {
		count++
	})
	duration := time.Since(start)
	
	if count != executions {
		t.Errorf("Expected %d executions, got %d", executions, count)
	}
	
	// 実行時間が期待される範囲内かチェック
	expectedMin := time.Duration(executions-1) * interval
	expectedMax := expectedMin + 200*time.Millisecond // 多少のマージン
	
	if duration < expectedMin {
		t.Errorf("RateLimiter completed too quickly: %v (expected at least %v)", duration, expectedMin)
	}
	
	if duration > expectedMax {
		t.Errorf("RateLimiter took too long: %v (expected at most %v)", duration, expectedMax)
	}
}

func TestRateLimiterZeroCount(t *testing.T) {
	count := 0
	RateLimiter(100*time.Millisecond, 0, func(i int) {
		count++
	})
	
	if count != 0 {
		t.Errorf("Expected 0 executions, got %d", count)
	}
}