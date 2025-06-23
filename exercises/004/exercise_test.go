package main

import (
	"fmt"
	"sort"
	"sync"
	"testing"
	"time"
)

func TestSimpleWorkerProcess(t *testing.T) {
	worker := &SimpleWorker{}
	
	tests := []struct {
		id       int
		expected string
	}{
		{0, "Processed: 0"},
		{1, "Processed: 1"},
		{42, "Processed: 42"},
		{-5, "Processed: -5"},
	}
	
	for _, tt := range tests {
		t.Run(fmt.Sprintf("id=%d", tt.id), func(t *testing.T) {
			result := worker.Process(tt.id)
			if result != tt.expected {
				t.Errorf("Process(%d) = %q, want %q", tt.id, result, tt.expected)
			}
		})
	}
}

func TestExercise004(t *testing.T) {
	worker := &SimpleWorker{}
	
	tests := []struct {
		numWorkers int
		expected   []string
	}{
		{
			numWorkers: 1,
			expected:   []string{"Processed: 0"},
		},
		{
			numWorkers: 3,
			expected:   []string{"Processed: 0", "Processed: 1", "Processed: 2"},
		},
		{
			numWorkers: 5,
			expected:   []string{"Processed: 0", "Processed: 1", "Processed: 2", "Processed: 3", "Processed: 4"},
		},
		{
			numWorkers: 0,
			expected:   []string{},
		},
	}
	
	for _, tt := range tests {
		t.Run(fmt.Sprintf("numWorkers=%d", tt.numWorkers), func(t *testing.T) {
			results := Exercise004(tt.numWorkers, worker)
			
			// Check length
			if len(results) != len(tt.expected) {
				t.Errorf("Exercise004(%d) returned %d results, want %d", 
					tt.numWorkers, len(results), len(tt.expected))
				return
			}
			
			// Sort both slices for comparison (order doesn't matter due to concurrency)
			sort.Strings(results)
			sort.Strings(tt.expected)
			
			// Compare sorted results
			for i := range results {
				if results[i] != tt.expected[i] {
					t.Errorf("Exercise004(%d) result[%d] = %q, want %q", 
						tt.numWorkers, i, results[i], tt.expected[i])
				}
			}
		})
	}
}

// Test for concurrent execution
func TestExercise004Concurrency(t *testing.T) {
	// Mock worker that tracks concurrent execution
	worker := &ConcurrencyTestWorker{
		activeWorkers: make(map[int]bool),
	}
	
	numWorkers := 10
	results := Exercise004(numWorkers, worker)
	
	if len(results) != numWorkers {
		t.Errorf("Expected %d results, got %d", numWorkers, len(results))
	}
	
	// Check that we had some concurrent execution
	if worker.maxConcurrent <= 1 && numWorkers > 1 {
		t.Error("Expected concurrent execution but workers ran sequentially")
	}
}

// Helper worker for testing concurrency
type ConcurrencyTestWorker struct {
	mu            sync.Mutex
	activeWorkers map[int]bool
	maxConcurrent int
}

func (w *ConcurrencyTestWorker) Process(id int) string {
	w.mu.Lock()
	w.activeWorkers[id] = true
	current := len(w.activeWorkers)
	if current > w.maxConcurrent {
		w.maxConcurrent = current
	}
	w.mu.Unlock()
	
	// Simulate some work
	time.Sleep(10 * time.Millisecond)
	
	w.mu.Lock()
	delete(w.activeWorkers, id)
	w.mu.Unlock()
	
	return fmt.Sprintf("Processed: %d", id)
}
