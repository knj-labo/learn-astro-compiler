package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
Exercise 008: 高度な並行処理とコンテキスト

このエクササイズでは、Goの高度な並行処理パターンとコンテキストを学びます：

1. WorkerPool を実装する
   - 固定数のワーカーでタスクを並行処理
   - チャネルを使ったタスクの配布
   - 結果の収集

2. TimeoutOperation 関数を実装する
   - 指定された時間内に処理を完了する
   - context.WithTimeout を使用
   - タイムアウト時はエラーを返す

3. RateLimiter を実装する
   - 一定間隔でのみ処理を許可
   - time.Ticker を使用
   - 指定された回数まで処理を実行

期待される動作:
- WorkerPool(3, tasks) で3つのワーカーでタスクを処理
- TimeoutOperation(1*time.Second, operation) で1秒以内に処理
- RateLimiter(100*time.Millisecond, 5) で100ms間隔で5回実行
*/

func main() {
	fmt.Println("Exercise 008: Advanced Concurrency and Context")
	
	// ワーカープールのテスト
	fmt.Println("=== Worker Pool Test ===")
	tasks := []Task{
		{ID: 1, Data: "Task 1"},
		{ID: 2, Data: "Task 2"},
		{ID: 3, Data: "Task 3"},
		{ID: 4, Data: "Task 4"},
		{ID: 5, Data: "Task 5"},
	}
	
	results := WorkerPool(3, tasks)
	for _, result := range results {
		fmt.Printf("Result: %s\n", result.Output)
	}
	
	// タイムアウト操作のテスト
	fmt.Println("\n=== Timeout Operation Test ===")
	
	// 成功する操作
	err := TimeoutOperation(2*time.Second, func() error {
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Operation completed successfully")
		return nil
	})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	
	// タイムアウトする操作
	err = TimeoutOperation(500*time.Millisecond, func() error {
		time.Sleep(1 * time.Second)
		fmt.Println("This should not print")
		return nil
	})
	if err != nil {
		fmt.Printf("Expected timeout error: %v\n", err)
	}
	
	// レート制限のテスト
	fmt.Println("\n=== Rate Limiter Test ===")
	RateLimiter(200*time.Millisecond, 3, func(i int) {
		fmt.Printf("Execution %d at %v\n", i, time.Now().Format("15:04:05.000"))
	})
}

// Task 構造体
type Task struct {
	ID   int
	Data string
}

// Result 構造体
type Result struct {
	TaskID int
	Output string
}

// WorkerPool 関数の実装
func WorkerPool(numWorkers int, tasks []Task) []Result {
	// 1. タスク用のチャネルを作成
	taskChan := make(chan Task, len(tasks))
	// 2. 結果用のチャネルを作成
	resultChan := make(chan Result, len(tasks))
	var wg sync.WaitGroup

	// 3. 指定された数のワーカーゴルーチンを起動
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i+1, taskChan, resultChan, &wg)
	}

	// 4. 各ワーカーはタスクを処理してResultを返す
	// 5. すべてのタスクを送信してチャネルを閉じる
	for _, task := range tasks {
		taskChan <- task
	}
	close(taskChan)

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// 6. 結果を収集して返す
	var results []Result
	for result := range resultChan {
		results = append(results, result)
	}

	return results
}

// worker 関数（ワーカープール用のヘルパー）
func worker(id int, tasks <-chan Task, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		// タスクを処理（実際の処理をシミュレート）
		time.Sleep(100 * time.Millisecond)
		result := Result{
			TaskID: task.ID,
			Output: fmt.Sprintf("Worker %d processed %s", id, task.Data),
		}
		results <- result
	}
}

// TimeoutOperation 関数の実装
func TimeoutOperation(timeout time.Duration, operation func() error) error {
	// 1. context.WithTimeout でコンテキストを作成
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// 2. done チャネルを作成
	done := make(chan error, 1)

	// 3. 別のゴルーチンで operation を実行
	go func() {
		done <- operation()
	}()

	// 4. select でタイムアウトと完了を待つ
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		// 5. タイムアウトの場合は適切なエラーを返す
		return ctx.Err()
	}
}

// RateLimiter 関数の実装
func RateLimiter(interval time.Duration, count int, operation func(int)) {
	if count <= 0 {
		return
	}

	// 1. time.NewTicker でティッカーを作成
	ticker := time.NewTicker(interval)
	// 2. defer ticker.Stop()
	defer ticker.Stop()

	// 3. 指定された回数だけループ
	for i := 0; i < count; i++ {
		// 4. <-ticker.C で間隔を待つ
		if i > 0 {
			<-ticker.C
		}
		// 5. operation を実行
		operation(i + 1)
	}
}