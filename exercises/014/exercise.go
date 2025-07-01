package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

/*
Exercise 014: 並行性とワーカープール

このエクササイズでは、Goの並行性機能を使ったワーカープールパターンを学びます：

1. ワーカープールの実装
   - 複数のgoroutineでタスクを並列処理
   - チャネルを使った通信
   - ワーカーの制御と同期

2. タスクスケジューリング
   - ジョブキューの管理
   - 優先度付きタスク処理
   - 負荷分散

3. 結果の集約
   - ワーカーからの結果収集
   - エラーハンドリング
   - 統計情報の計算

期待される動作:
- 効率的な並列処理
- リソースの適切な管理
- 結果の正確な集約
*/

func main() {
	fmt.Println("Exercise 014: Concurrency and Worker Pools")
	
	// ワーカープールマネージャーを初期化
	manager := NewWorkerPoolManager(5, 100) // 5ワーカー、キューサイズ100
	
	// ワーカープールを開始
	fmt.Println("=== Starting Worker Pool ===")
	manager.Start()
	defer manager.Stop()
	
	// 数値処理タスクのデモ
	fmt.Println("\n=== Number Processing Tasks ===")
	numbers := []int{10, 15, 20, 25, 30, 35, 40, 45, 50}
	
	// 平方計算タスクを送信
	for _, num := range numbers {
		task := Task{
			ID:       fmt.Sprintf("square-%d", num),
			Type:     "square",
			Data:     num,
			Priority: 1,
		}
		manager.SubmitTask(task)
	}
	
	// 素数チェックタスクを送信
	primeNumbers := []int{17, 29, 31, 37, 41, 43, 47, 53, 59, 61}
	for _, num := range primeNumbers {
		task := Task{
			ID:       fmt.Sprintf("prime-%d", num),
			Type:     "prime",
			Data:     num,
			Priority: 2, // 高優先度
		}
		manager.SubmitTask(task)
	}
	
	// ファイル処理タスクのデモ
	fmt.Println("\n=== File Processing Tasks ===")
	fileSizes := []int{1024, 2048, 4096, 8192, 16384}
	for i, size := range fileSizes {
		task := Task{
			ID:       fmt.Sprintf("file-%d", i),
			Type:     "file",
			Data:     size,
			Priority: 1,
		}
		manager.SubmitTask(task)
	}
	
	// ネットワークタスクのデモ
	fmt.Println("\n=== Network Tasks ===")
	urls := []string{
		"https://httpbin.org/delay/1",
		"https://httpbin.org/delay/2",
		"https://httpbin.org/delay/3",
	}
	for i, url := range urls {
		task := Task{
			ID:       fmt.Sprintf("network-%d", i),
			Type:     "network",
			Data:     url,
			Priority: 3, // 最高優先度
		}
		manager.SubmitTask(task)
	}
	
	// 処理完了を待つ
	fmt.Println("\n=== Waiting for completion ===")
	time.Sleep(10 * time.Second)
	
	// 統計情報を表示
	fmt.Println("\n=== Statistics ===")
	stats := manager.GetStatistics()
	fmt.Printf("Total tasks submitted: %d\n", stats.TotalSubmitted)
	fmt.Printf("Total tasks completed: %d\n", stats.TotalCompleted)
	fmt.Printf("Total tasks failed: %d\n", stats.TotalFailed)
	fmt.Printf("Average processing time: %.2fms\n", stats.AvgProcessingTime)
	fmt.Printf("Active workers: %d\n", stats.ActiveWorkers)
	fmt.Printf("Queue size: %d\n", stats.QueueSize)
	
	// バッチ処理のデモ
	fmt.Println("\n=== Batch Processing ===")
	batchProcessor := NewBatchProcessor(3, 50) // 3ワーカー、バッチサイズ50
	
	// 大量のタスクを生成
	var batchTasks []BatchItem
	for i := 0; i < 200; i++ {
		item := BatchItem{
			ID:   i,
			Data: rand.Intn(1000),
		}
		batchTasks = append(batchTasks, item)
	}
	
	// バッチ処理を実行
	results := batchProcessor.ProcessBatch(batchTasks)
	fmt.Printf("Processed %d items in batch\n", len(results))
	
	// 結果の統計
	var sum int
	for _, result := range results {
		if val, ok := result.Result.(int); ok {
			sum += val
		}
	}
	fmt.Printf("Sum of results: %d\n", sum)
	fmt.Printf("Average result: %.2f\n", float64(sum)/float64(len(results)))
}

// Task構造体
type Task struct {
	ID       string      `json:"id"`
	Type     string      `json:"type"`
	Data     interface{} `json:"data"`
	Priority int         `json:"priority"`
	SubmitTime time.Time `json:"submit_time"`
}

// TaskResult構造体
type TaskResult struct {
	TaskID      string        `json:"task_id"`
	Success     bool          `json:"success"`
	Result      interface{}   `json:"result"`
	Error       string        `json:"error,omitempty"`
	ProcessTime time.Duration `json:"process_time"`
}

// WorkerStatistics構造体
type WorkerStatistics struct {
	TotalSubmitted     int     `json:"total_submitted"`
	TotalCompleted     int     `json:"total_completed"`
	TotalFailed        int     `json:"total_failed"`
	AvgProcessingTime  float64 `json:"avg_processing_time_ms"`
	ActiveWorkers      int     `json:"active_workers"`
	QueueSize          int     `json:"queue_size"`
}

// WorkerPoolManager構造体
type WorkerPoolManager struct {
	workerCount int
	taskQueue   chan Task
	resultQueue chan TaskResult
	workers     []*Worker
	statistics  *WorkerStatistics
	mutex       sync.RWMutex
	stopChan    chan bool
	wg          sync.WaitGroup
}

// Worker構造体
type Worker struct {
	ID          int
	taskQueue   chan Task
	resultQueue chan TaskResult
	stopChan    chan bool
	active      bool
	mutex       sync.RWMutex
}

// BatchItem構造体
type BatchItem struct {
	ID   int         `json:"id"`
	Data interface{} `json:"data"`
}

// BatchResult構造体
type BatchResult struct {
	ID     int         `json:"id"`
	Result interface{} `json:"result"`
	Error  string      `json:"error,omitempty"`
}

// BatchProcessor構造体
type BatchProcessor struct {
	workerCount int
	batchSize   int
}

// NewWorkerPoolManager関数の実装
func NewWorkerPoolManager(workerCount, queueSize int) *WorkerPoolManager {
	// 1. WorkerPoolManager構造体を初期化
	// 2. タスクキューと結果キューを作成
	taskQueue := make(chan Task, queueSize)
	resultQueue := make(chan TaskResult, queueSize)
	stopChan := make(chan bool, workerCount)
	
	// 3. 統計情報を初期化
	statistics := &WorkerStatistics{}
	
	wpm := &WorkerPoolManager{
		workerCount: workerCount,
		taskQueue:   taskQueue,
		resultQueue: resultQueue,
		statistics:  statistics,
		stopChan:    stopChan,
	}
	
	// 4. ワーカーを作成（まだ開始しない）
	for i := 0; i < workerCount; i++ {
		worker := NewWorker(i, taskQueue, resultQueue, stopChan)
		wpm.workers = append(wpm.workers, worker)
	}
	
	return wpm
}

// Start メソッドの実装
func (wpm *WorkerPoolManager) Start() {
	// 1. 各ワーカーを開始
	// 3. ワーカーグループに追加
	for _, worker := range wpm.workers {
		wpm.wg.Add(1)
		go func(w *Worker) {
			defer wpm.wg.Done()
			w.Start()
		}(worker)
	}
	
	// 2. 結果処理のgoroutineを開始
	go wpm.processResults()
}

// Stop メソッドの実装
func (wpm *WorkerPoolManager) Stop() {
	// 1. stopChanを使ってワーカーに停止信号を送信
	for i := 0; i < len(wpm.workers); i++ {
		select {
		case wpm.stopChan <- true:
		default:
		}
	}
	
	// 2. 全ワーカーの完了を待つ
	wpm.wg.Wait()
	
	// 3. チャネルを閉じる
	close(wpm.taskQueue)
	close(wpm.resultQueue)
}

// SubmitTask メソッドの実装
func (wpm *WorkerPoolManager) SubmitTask(task Task) {
	// 1. タスクに送信時刻を設定
	task.SubmitTime = time.Now()
	
	// 2. 優先度に基づいてタスクをキューに追加
	select {
	case wpm.taskQueue <- task:
		// 3. 統計情報を更新
		wpm.mutex.Lock()
		wpm.statistics.TotalSubmitted++
		wpm.statistics.QueueSize = len(wpm.taskQueue)
		wpm.mutex.Unlock()
	default:
		// キューが満杯の場合
		fmt.Printf("Task queue is full, dropping task %s\n", task.ID)
	}
}

// GetStatistics メソッドの実装
func (wpm *WorkerPoolManager) GetStatistics() WorkerStatistics {
	// 1. 読み取りロックを取得
	wpm.mutex.RLock()
	defer wpm.mutex.RUnlock()
	
	// 2. 現在の統計情報をコピーして返す
	stats := *wpm.statistics
	stats.ActiveWorkers = len(wpm.workers)
	stats.QueueSize = len(wpm.taskQueue)
	
	return stats
}

// processResults processes task results and updates statistics
func (wpm *WorkerPoolManager) processResults() {
	var totalProcessTime time.Duration
	var processedCount int
	
	for result := range wpm.resultQueue {
		wpm.mutex.Lock()
		if result.Success {
			wpm.statistics.TotalCompleted++
		} else {
			wpm.statistics.TotalFailed++
		}
		
		totalProcessTime += result.ProcessTime
		processedCount++
		
		if processedCount > 0 {
			wpm.statistics.AvgProcessingTime = float64(totalProcessTime.Nanoseconds()) / float64(processedCount) / 1e6 // ms
		}
		
		wpm.mutex.Unlock()
	}
}

// NewWorker関数の実装
func NewWorker(id int, taskQueue chan Task, resultQueue chan TaskResult, stopChan chan bool) *Worker {
	// 1. Worker構造体を初期化
	// 2. チャネルを設定
	return &Worker{
		ID:          id,
		taskQueue:   taskQueue,
		resultQueue: resultQueue,
		stopChan:    stopChan,
		active:      false,
	}
}

// Start メソッド（Worker）の実装
func (w *Worker) Start() {
	// 1. goroutineでワーカーのメインループを開始
	w.mutex.Lock()
	w.active = true
	w.mutex.Unlock()
	
	for {
		select {
		// 2. タスクキューからタスクを受信
		case task, ok := <-w.taskQueue:
			if !ok {
				return // チャネルが閉じられた
			}
			
			// 3. タスクを処理して結果を送信
			result := w.processTask(task)
			
			select {
			case w.resultQueue <- result:
			default:
				// 結果キューが満杯の場合はログ出力
				fmt.Printf("Result queue is full, dropping result for task %s\n", task.ID)
			}
			
		// 4. 停止信号をチェック
		case <-w.stopChan:
			w.mutex.Lock()
			w.active = false
			w.mutex.Unlock()
			return
		}
	}
}

// processTask メソッドの実装
func (w *Worker) processTask(task Task) TaskResult {
	// 2. 処理時間を測定
	startTime := time.Now()
	
	result := TaskResult{
		TaskID:  task.ID,
		Success: true,
	}
	
	// 1. タスクタイプに基づいて適切な処理を実行
	// 3. エラーハンドリングを実装
	switch task.Type {
	case "square":
		if num, ok := task.Data.(int); ok {
			result.Result = calculateSquare(num)
		} else {
			result.Success = false
			result.Error = "Invalid data type for square operation"
		}
		
	case "prime":
		if num, ok := task.Data.(int); ok {
			result.Result = isPrime(num)
		} else {
			result.Success = false
			result.Error = "Invalid data type for prime operation"
		}
		
	case "file":
		if size, ok := task.Data.(int); ok {
			result.Result = simulateFileProcessing(size)
		} else {
			result.Success = false
			result.Error = "Invalid data type for file operation"
		}
		
	case "network":
		if url, ok := task.Data.(string); ok {
			networkResult, err := simulateNetworkRequest(url)
			if err != nil {
				result.Success = false
				result.Error = err.Error()
			} else {
				result.Result = networkResult
			}
		} else {
			result.Success = false
			result.Error = "Invalid data type for network operation"
		}
		
	default:
		result.Success = false
		result.Error = fmt.Sprintf("Unknown task type: %s", task.Type)
	}
	
	// 4. TaskResult構造体を返す
	result.ProcessTime = time.Since(startTime)
	return result
}

// NewBatchProcessor関数の実装
func NewBatchProcessor(workerCount, batchSize int) *BatchProcessor {
	// 1. BatchProcessor構造体を初期化
	return &BatchProcessor{
		workerCount: workerCount,
		batchSize:   batchSize,
	}
}

// ProcessBatch メソッドの実装
func (bp *BatchProcessor) ProcessBatch(items []BatchItem) []BatchResult {
	results := make([]BatchResult, 0, len(items))
	resultChan := make(chan BatchResult, len(items))
	
	// 4. sync.WaitGroupを使用して同期
	var wg sync.WaitGroup
	
	// 1. アイテムをバッチに分割
	batches := make([][]BatchItem, 0)
	for i := 0; i < len(items); i += bp.batchSize {
		end := i + bp.batchSize
		if end > len(items) {
			end = len(items)
		}
		batches = append(batches, items[i:end])
	}
	
	// 2. 各バッチを並列処理
	batchChan := make(chan []BatchItem, len(batches))
	
	// ワーカーを開始
	for i := 0; i < bp.workerCount; i++ {
		wg.Add(1)
		go processBatchItems(batchChan, resultChan, &wg)
	}
	
	// バッチをワーカーに送信
	go func() {
		for _, batch := range batches {
			batchChan <- batch
		}
		close(batchChan)
	}()
	
	// 3. 結果を収集して返す
	go func() {
		wg.Wait()
		close(resultChan)
	}()
	
	for result := range resultChan {
		results = append(results, result)
	}
	
	return results
}

// processBatchItems ヘルパー関数の実装
func processBatchItems(batchChan <-chan []BatchItem, results chan<- BatchResult, wg *sync.WaitGroup) {
	// 1. defer wg.Done()を設定
	defer wg.Done()
	
	for batch := range batchChan {
		// 2. 各アイテムを処理
		for _, item := range batch {
			result := BatchResult{
				ID: item.ID,
			}
			
			// 簡単な処理例：データを2倍にする
			if num, ok := item.Data.(int); ok {
				result.Result = num * 2
			} else {
				result.Error = "Invalid data type"
			}
			
			// 3. 結果をチャネルに送信
			results <- result
		}
	}
}

// calculateSquare ヘルパー関数の実装
func calculateSquare(n int) int {
	// 2. 処理時間をシミュレート
	time.Sleep(time.Duration(rand.Intn(10)+1) * time.Millisecond)
	
	// 1. 平方を計算
	return n * n
}

// isPrime ヘルパー関数の実装
func isPrime(n int) bool {
	// 処理時間をシミュレート
	time.Sleep(time.Duration(rand.Intn(20)+10) * time.Millisecond)
	
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	
	// 1. 素数判定のアルゴリズムを実装
	// 2. 2からsqrt(n)まで除算チェック
	sqrt := int(math.Sqrt(float64(n)))
	for i := 3; i <= sqrt; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// simulateFileProcessing ヘルパー関数の実装
func simulateFileProcessing(size int) string {
	// 2. ランダムな処理時間を追加
	// サイズに基づいて処理時間を計算
	processingTime := time.Duration(size/100+rand.Intn(50)) * time.Millisecond
	time.Sleep(processingTime)
	
	// 1. ファイル処理をシミュレート
	// 3. 処理結果を文字列で返す
	return fmt.Sprintf("Processed file of size %d bytes in %v", size, processingTime)
}

// simulateNetworkRequest ヘルパー関数の実装
func simulateNetworkRequest(url string) (string, error) {
	// 2. time.Sleep()で遅延を追加
	delay := time.Duration(rand.Intn(1000)+500) * time.Millisecond
	time.Sleep(delay)
	
	// 3. ランダムにエラーを発生させる
	if rand.Float32() < 0.1 { // 10%の確率でエラー
		return "", fmt.Errorf("network timeout for URL: %s", url)
	}
	
	// 1. ネットワークリクエストをシミュレート
	return fmt.Sprintf("Response from %s (took %v)", url, delay), nil
}